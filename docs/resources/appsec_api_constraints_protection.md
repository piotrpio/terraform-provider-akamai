---
layout: "akamai"
page_title: "Akamai: API Constraints Protection"
subcategory: "Application Security"
description: |-
 API Constraints Protection
---

# akamai_appsec_api_constraints_protection

**Scopes**: Security policy

Enables or disables API constraints protection. These constraints specify the action to be taken when designated API endpoints are invoked.

**Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putprotections)

## Example Usage

Basic usage:

```
terraform {
  required_providers {
    akamai = {
      source = "akamai/akamai"
    }
  }
}

provider "akamai" {
  edgerc = "~/.edgerc"
}

// USE CASE: User wants to enable or disable API constraints protection.

data "akamai_appsec_configuration" "configuration" {
  name = "Documentation"
}

resource "akamai_appsec_api_constraints_protection" "protection" {
  config_id          = data.akamai_appsec_configuration.configuration.config_id
  security_policy_id = "gms1_134637"
  enabled            = true
}
```

## Argument Reference

This resource supports the following arguments:

`config_id` (Required). Unique identifier of the security configuration associated with the API constraint protection settings being modified.

`security_policy_id` (Required). Unique identifier of the security policy associated with the API constraint protection settings being modified.

`enabled` (Required). Set to **true** to enable API constraints protection; set to **false** to disable API constraints protection.

## Output Options

The following options can be used to determine the information returned, and how that returned information is formatted:

- `output_text`. Tabular report showing the current protection settings.

