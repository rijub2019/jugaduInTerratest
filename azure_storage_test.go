package test

import (
	"strings"
	"testing"
	"os/exec"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAzureStorage(t *testing.T) {
	t.Parallel()

	// subscriptionID is overridden by the environment variable "ARM_SUBSCRIPTION_ID"
	//subscriptionID := "61437745-1eff-43c2-9e82-3770481b9e24"
	uniquePostfix := random.UniqueId()

	terraformOptions := &terraform.Options{
		
		TerraformDir: "./terraform/",

		// #1 Input Variables
		Vars: map[string]interface{}{
			"postfix": strings.ToLower(uniquePostfix),
		},
	}

	// #5 Terraform Destroy
	defer terraform.Destroy(t, terraformOptions)

	// #2 Terraform Init & Apply
	terraform.InitAndApply(t, terraformOptions)

	// #3 Fetch Resources From Terraform Outputs
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")

    // #4 Validate Resources According To Test Conditions
	// Validate versioning is enabled in Azure storage account
	qrystrng :=  "isVersioningEnabled"
	expectedVE := "true"
	cmd := exec.Command("az", "storage", "account", "blob-service-properties", "show", "--account-name", storageAccountName, "--query", qrystrng)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run command %s: %s", cmd, err)
	}
	actualVE := strings.TrimSpace(string(output))
	
	assert.Equal(t, expectedVE, actualVE, "Versioning should be enabled in blob-service-properties of StorageAccount: %s", storageAccountName)
}