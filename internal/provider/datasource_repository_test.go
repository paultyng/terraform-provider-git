package provider

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
	dir = filepath.Join(dir, "../..")
	dir = filepath.ToSlash(dir)

	expectedBranch := execGit(t, "rev-parse", "--abbrev-ref", "HEAD")
	expectedCommit := execGit(t, "rev-parse", "HEAD")

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"git": func() (*schema.Provider, error) {
				return New(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testDataSourceRepositoryConfig(dir),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "branch", strings.TrimSpace(string(expectedBranch))),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_hash", strings.TrimSpace(string(expectedCommit))),
				),
			},
			{
				Config: testDataSourceRepositoryConfigDotGit(true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "branch", strings.TrimSpace(string(expectedBranch))),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_hash", strings.TrimSpace(string(expectedCommit))),
				),
			},
			{
				Config: testDataSourceRepositoryConfigPathDotGit(dir, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "branch", strings.TrimSpace(string(expectedBranch))),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_hash", strings.TrimSpace(string(expectedCommit))),
				),
			},
			{
				Config: testDataSourceRepositoryConfigPathDotGit(dir, false),
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
data "git_repository" "test" {
	path = "%s"
}
`, path)
}

func testDataSourceRepositoryConfigDotGit(detectDotGit bool) string {
	return fmt.Sprintf(`
data "git_repository" "test" {
	detect_dot_git = "%s"
}
`, detectDotGit)
}

func testDataSourceRepositoryConfigPathDotGit(path string, detectDotGit bool) string {
	return fmt.Sprintf(`
data "git_repository" "test" {
        path = "%s"
	detect_dot_git = "%s"
}
`, path, detectDotGit)
}
