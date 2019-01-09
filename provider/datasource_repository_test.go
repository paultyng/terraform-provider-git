package provider

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testProviders = map[string]terraform.ResourceProvider{
	"git": Provider(),
}

func execGit(t *testing.T, arg ...string) string {
	t.Helper()

	output, err := exec.Command("git", arg...).Output()
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(output))
}

func TestDataSourceRepository(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir = filepath.Join(dir, "..")

	expectedBranch := execGit(t, "rev-parse", "--abbrev-ref", "HEAD")
	expectedCommit := execGit(t, "rev-parse", "HEAD")

	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceRepositoryConfig(dir),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "branch", strings.TrimSpace(string(expectedBranch))),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_hash", strings.TrimSpace(string(expectedCommit))),
				),
			},
		},
	})
}

func testDataSourceRepositoryConfig(path string) string {
	return fmt.Sprintf(`
data git_repository test {
	path = "%s"
}
`, path)
}
