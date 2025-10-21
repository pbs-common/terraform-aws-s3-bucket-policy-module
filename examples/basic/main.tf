module "s3" {
  source = "github.com/pbs-common/terraform-aws-s3-module?ref=4.0.14"

  create_bucket_policy = false

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}

module "s3_bucket_policy" {
  source = "../.."

  name = module.s3.name

  product = var.product
}
