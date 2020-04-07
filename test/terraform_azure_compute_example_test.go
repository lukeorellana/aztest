package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"aztest/modules/azure"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var approvedRegions = []string{
	// Americas
	"centralus",
	"eastus",
	"eastus2",
	"northcentralus",
	"southcentralus",
	"westcentralus",
	"westus",
	"westus2",
}

// An example of how to test the Terraform module in examples/terraform-azure-example using Terratest.
func TestTerraformAzureComputeExample(t *testing.T) {
	t.Parallel()

	// Pick a random Azure region to test in.
	azureRegion := azure.GetRandomRegion(t, approvedRegions, nil, os.Getenv("ARM_SUBSCRIPTION_ID"))

	// Network Settings for Vnet and Subnet
	systemName := fmt.Sprintf("test-%s", strings.ToLower(random.UniqueId()))
	vnetAddress := "10.0.0.0/16"
	subnetPrefix := "10.0.0.0/24"

	terraformOptions := &terraform.Options{

		// The path to where our Terraform code is located
		TerraformDir: "../examples/compute",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"system":             systemName,
			"location":           azureRegion,
			"vnet_address_space": vnetAddress,
			"subnet_prefix":      subnetPrefix,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	vnetRG := terraform.Output(t, terraformOptions, "vnet_rg")
	subnetID := terraform.Output(t, terraformOptions, "subnet_id")
	nsgName := terraform.Output(t, terraformOptions, "nsg_name")
	vnetName := terraform.Output(t, terraformOptions, "vnet_name")

	// Look up the size of the given Virtual Machine
	actualVMSize := azure.GetSizeOfVirtualMachine(t, vmName, resourceGroupName, "")
	expectedVMSize := compute.VirtualMachineSizeTypes("Standard_B1s")

	// Test that the Virtual Machine size matches the Terraform specification
	assert.Equal(t, expectedVMSize, actualVMSize)

	// Look up the Disk size of the given Virtual Machine
	actualDiskType := azure.GetSizeOfVirtualMachineDisk(t, vmName, resourceGroupName, "")
	expectedDiskType := compute.StorageAccountTypes("Standard_LRS")

	// Test that the Virtual Machine size matches the Terraform specification
	assert.Equal(t, expectedDiskType, actualDiskType)
}
