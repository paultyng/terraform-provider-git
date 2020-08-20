---
page_title: "git_repository Data Source - terraform-provider-git"
subcategory: ""
description: |-
  
---

# Data Source `git_repository`



## Example Usage

```terraform
data "git_repository" "tf" {
  path = "${path.module}"
}
```

## Schema

### Optional

- **id** (String, Optional) The ID of this resource.
- **path** (String, Optional) Path to the .git directory

### Read-only

- **branch** (String, Read-only)
- **commit_hash** (String, Read-only)


