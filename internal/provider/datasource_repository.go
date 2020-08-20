package provider

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	git "gopkg.in/src-d/go-git.v4"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the .git directory",
			},

			"commit_hash": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)

	log.Printf("[INFO] opening repository in %s", path)

	repo, err := git.PlainOpen(path)
	if err != nil {
		log.Printf("[ERROR] err opening repo: %s", err)
		return err
	}

	head, err := repo.Head()
	if err != nil {
		log.Printf("[ERROR] err reading HEAD: %s", err)
		return err
	}

	d.Set("commit_hash", head.Hash().String())
	d.Set("branch", "")

	refName := head.Name()

	d.SetId(refName.String())

	switch {
	case refName.IsBranch():
		d.Set("branch", refName.Short())
	}

	return nil
}
