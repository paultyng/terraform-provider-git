---
layout: ""
page_title: "Provider: Git"
description: |-
  The Git provider provides resources to interact with a local Git repository.
---

# Git Provider

The Git provider provides resources to interact with a local Git repository. Primarily this is useful
for using your current commit hash or tag in resource labeling, etc.

## Example Usage

```terraform
data "git_repository" "tf" {
  path = "${path.module}"
}
```

## Schema
