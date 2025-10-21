# PBS TF S3 Bucket Policy Module

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs-common/terraform-aws-s3-bucket-policy-module?ref=1.0.24
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

Provisions a bucket policy for an S3 bucket.

By default, adds a rule that prevents HTTP requests without TLS to the bucket. To disable this rule, set `force_tls` to `false` (this is not advised).

The `replication_source` variable configures permissions for S3 replication to take place. If this is set, the module will add a rule that allows the replication role to read from the bucket.

The `allow_anonymous_vpce_access` variable configures permissions for VPC endpoints to access the bucket. If this is set, the module will add a rule that allows the VPC endpoint to read from the bucket. The `vpce` variable must be set to the VPC endpoint ID that is to be allowed access.

Use `source_policy_documents` and `override_policy_documents` to add extra statements for the policy that you need if the module doesn't provide the statements out of the box. See [this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#example-using-a-source-document) for examples.

Integrate this module like so:

```hcl
module "s3_bucket_policy" {
  source = "github.com/pbs-common/terraform-aws-s3-bucket-policy-module?ref=1.0.24"

  # Tagging Parameters
  product      = var.product

  # Optional Parameters
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`1.0.24`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.9.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.24.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_s3_bucket_policy.bucket_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_policy) | resource |
| [aws_iam_policy_document.bucket_policy_doc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_product"></a> [product](#input\_product) | Tag used to group resources according to product | `string` | n/a | yes |
| <a name="input_allow_anonymous_vpce_access"></a> [allow\_anonymous\_vpce\_access](#input\_allow\_anonymous\_vpce\_access) | Create bucket policy that allows anonymous VPCE access. | `bool` | `false` | no |
| <a name="input_bucket_policy"></a> [bucket\_policy](#input\_bucket\_policy) | Policy to apply to the bucket. If null, one will be guessed based on other variables. | `string` | `null` | no |
| <a name="input_cloudfront_oac_access_statements"></a> [cloudfront\_oac\_access\_statements](#input\_cloudfront\_oac\_access\_statements) | List of objects that define the CloudFront origin access identity access statement. Each object must have a `cloudfront_arn` and `path` key. | <pre>list(object({<br>    cloudfront_arn = string<br>    path           = optional(string, "*")<br>  }))</pre> | `[]` | no |
| <a name="input_force_tls"></a> [force\_tls](#input\_force\_tls) | Deny HTTP requests that are made to the bucket without TLS. | `bool` | `true` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the s3 bucket this policy is to apply to. If null, will default to product. | `string` | `null` | no |
| <a name="input_override_policy_documents"></a> [override\_policy\_documents](#input\_override\_policy\_documents) | List of IAM policy documents that are merged together into the exported document. In merging, statements with non-blank sids will override statements with the same sid from earlier documents in the list. Statements with non-blank sids will also override statements with the same sid from documents provided in the source\_json and source\_policy\_documents arguments. Non-overriding statements will be added to the exported document. | `list(string)` | `null` | no |
| <a name="input_replication_source"></a> [replication\_source](#input\_replication\_source) | The account number and role for the source bucket in a replication configuration. | <pre>object({<br>    account_id = string<br>    role       = string<br>  })</pre> | `null` | no |
| <a name="input_source_policy_documents"></a> [source\_policy\_documents](#input\_source\_policy\_documents) | List of IAM policy documents that are merged together into the exported document. Statements defined in source\_policy\_documents or source\_json must have unique sids. Statements with the same sid from documents assigned to the override\_json and override\_policy\_documents arguments will override source statements. | `list(string)` | `null` | no |
| <a name="input_vpce"></a> [vpce](#input\_vpce) | Name of the VPC endpoint that should have access to this bucket. Only used when `allow_anonymous_vpce_access` is true. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_bucket_policy"></a> [bucket\_policy](#output\_bucket\_policy) | JSON value of the bucket policy. |
