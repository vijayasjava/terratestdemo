package terratest

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEc2PowerUser(t *testing.T) {
	t.Parallel()
	awsRegion := "us-east-1"
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../tests",
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})
	defer terraform.Destroy(t, terraformOptions)
	//terraform.InitAndApply(t, terraformOptions)
	terraform.InitAndPlanAndShow(t, terraformOptions)
	ec2PowerUserPolicy := terraform.Output(t, terraformOptions, "ec2_poweruser")
	assert.Contains(t, "ec2:TerminateInstances", ec2PowerUserPolicy)
	assert.Contains(t, strings.ToLower("ModifyLaunchTemplate"), strings.ToLower(ec2PowerUserPolicy))
	assert.Contains(t, strings.ToLower("StartInstances"), strings.ToLower(ec2PowerUserPolicy))
	assert.Contains(t, strings.ToLower("StopInstances"), strings.ToLower(ec2PowerUserPolicy))

	assert.Contains(t, strings.ToLower("stop"), strings.ToLower(ec2PowerUserPolicy))
	assert.Contains(t, strings.ToLower("delete"), strings.ToLower(ec2PowerUserPolicy))
	assert.Contains(t, strings.ToLower("terminate"), strings.ToLower(ec2PowerUserPolicy))

}
