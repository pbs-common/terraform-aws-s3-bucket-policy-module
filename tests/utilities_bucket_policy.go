package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testBucketPolicy(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
		Upgrade:      true,
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketName := terraform.Output(t, terraformOptions, "bucket_name")

	expectedPolicy := fmt.Sprintf(`{"Statement":[{"Action":"s3:*","Condition":{"Bool":{"aws:SecureTransport":"false"}},"Effect":"Deny","Principal":"*","Resource":["arn:aws:s3:::%s/*","arn:aws:s3:::%s"]}],"Version":"2012-10-17"}`, bucketName, bucketName)

	bucketPolicy := terraform.Output(t, terraformOptions, "bucket_policy")

	fmt.Println("Expected Policy:", expectedPolicy)
	fmt.Println("Bucket Policy:", bucketPolicy)
	assert.Equal(t, expectedPolicy, bucketPolicy)
}
