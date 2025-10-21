# PBS TF S3 Bucket Policy Module

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs-common/terraform-aws-s3-bucket-policy-module?ref=x.y.z
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
  source = "github.com/pbs-common/terraform-aws-s3-bucket-policy-module?ref=x.y.z"

  # Tagging Parameters
  product      = var.product

  # Optional Parameters
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs
