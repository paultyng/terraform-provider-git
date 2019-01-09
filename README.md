# Terraform Provider for Local Git Repository

## Usage

```hcl
data git_repository tf {
	path = "${path.module}"
}

resource "aws_vpc" "main" {
  tags = {
    git_branch = "${data.git_repository.tf.branch}"
  }
}
```
