package test

import (
	"testing"

	"github.com/allanore/aztest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-azure-example using Terratest.
func TestTerraformAzureNetworkingExample(t *testing.T) {
	t.Parallel()

	// Sandbox Details to Check under
	terraformOptions := &terraform.Options{

		// The path to where our Terraform code is located
		TerraformDir: "../examples/network",
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	appSubnet := azure.GetSubnetbyName(t, ResourceGroup, VritualNetwork, Subnet, "")

	// Look up the subnet ID from the Virtual Network Name
	subnets := azure.GetSubnetsforVnet(t, VritualNetwork, ResourceGroup, "")

	// Look up Subnet and NIC ID associations of NSG
	associations := azure.GetAssociationsforNSG(t, NetworkSecurityGroup, ResourceGroup, "")

	//Check if the subnet exists in the Virtual Network
	assert.Contains(t, subnets, *appSubnet.ID)

	//Check if subnet is associated wtih NSG
	assert.Contains(t, associations, *appSubnet.ID)

	//Check if Subnet has a /28 CIDR Notation
	subnetAddressPrefix := *appSubnet.SubnetPropertiesFormat.AddressPrefix
	expectedCIDR := "/28"
	actualCIDR := subnetAddressPrefix[len(subnetAddressPrefix)-3:]
	assert.Equal(t, expectedCIDR, actualCIDR)

}
