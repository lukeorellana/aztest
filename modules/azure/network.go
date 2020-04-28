package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/stretchr/testify/require"
)

// GetSecurityGroupsClient is a helper function that will setup an Azure SecurityGroups client
func GetSecurityGroupsClient(subscriptionID string) (*network.SecurityGroupsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a Network client
	nsgClient := network.NewSecurityGroupsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	nsgClient.Authorizer = *authorizer

	return &nsgClient, nil
}

// GetSecurityGroupsClient is a helper function that will setup an Azure SecurityGroups client
func GetSubnetsClient(subscriptionID string) (*network.SubnetsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a Network client
	snetClient := network.NewSubnetsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	snetClient.Authorizer = *authorizer

	return &snetClient, nil
}

// GetVirtualNetworkClient is a helper function to setup an Azure Virtual Network client
func GetVirtualNetworkClient(subscriptionID string) (*network.VirtualNetworksClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a VNet client
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	vnetClient.Authorizer = *authorizer

	return &vnetClient, nil
}

// GetSubnetsforVnet gets the list of subnets from a given Azure Virtual Network Name
func GetSubnetsforVnet(t *testing.T, resGroupName string, vNetName string, subscriptionID string) []string {
	subnets, err := GetSubnetsforVnetE(t, resGroupName, vNetName, subscriptionID)
	require.NoError(t, err)

	return subnets
}

// GetSubnetsforVnetE gets the list of subnets from a given Azure Virtual Network Name
func GetSubnetsforVnetE(t *testing.T, resGroupName string, vNetName string, subscriptionID string) ([]string, error) {

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Create a Vritual Networks client
	vnetClient, err := GetVirtualNetworkClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the details of the target virtual Network
	vnet, err := vnetClient.Get(context.Background(), resGroupName, vNetName, "")
	if err != nil {
		return nil, err
	}

	subnets := make([]string, len(*vnet.VirtualNetworkPropertiesFormat.Subnets))
	for i, subnet := range *vnet.VirtualNetworkPropertiesFormat.Subnets {
		subnets[i] = *subnet.ID
	}

	return subnets, nil
}

// GetAssociationsforNSG gets the Subnet and NIC ID associations of a given Network Security Group
func GetAssociationsforNSG(t *testing.T, resGroupName string, nsgName string, subscriptionID string) []string {
	nsgAssociations, err := GetAssociationsforNSGE(t, resGroupName, nsgName, subscriptionID)
	require.NoError(t, err)

	return nsgAssociations
}

// GetAssociationsforNSG gets the Subnet and NIC ID associations of a given Network Security Group
func GetAssociationsforNSGE(t *testing.T, resGroupName string, nsgName string, subscriptionID string) ([]string, error) {

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Create a Security Group client
	nsgClient, err := GetSecurityGroupsClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the details of the target virtual Network
	nsg, err := nsgClient.Get(context.Background(), resGroupName, nsgName, "")
	if err != nil {
		return nil, err
	}

	nsgAssociations := []string{}

	// Collect any NICs associated to the Network Security Group
	if nsg.SecurityGroupPropertiesFormat.NetworkInterfaces != nil {
		associatedNics := make([]string, len(*nsg.SecurityGroupPropertiesFormat.NetworkInterfaces))
		for i, nic := range *nsg.SecurityGroupPropertiesFormat.NetworkInterfaces {
			associatedNics[i] = *nic.ID
		}
		nsgAssociations = append(nsgAssociations, associatedNics...)

	}

	// Collect any subnets associated to the Network Security Group
	if nsg.SecurityGroupPropertiesFormat.Subnets != nil {
		associatedSubnets := make([]string, len(*nsg.SecurityGroupPropertiesFormat.Subnets))
		for i, subnet := range *nsg.SecurityGroupPropertiesFormat.Subnets {
			associatedSubnets[i] = *subnet.ID
		}
		nsgAssociations = append(nsgAssociations, associatedSubnets...)
	}

	return nsgAssociations, nil
}

// GetVnetbyName gets propteries of the Azure Virtual Network by its given name
func GetVnetbyName(t *testing.T, resGroupName string, vNetName string, subscriptionID string) network.VirtualNetwork {
	vnet, err := GetVnetbyNameE(t, resGroupName, vNetName, subscriptionID)
	require.NoError(t, err)

	return vnet
}

// GetVnetbyName gets propteries of the Azure Virtual Network by its given name
func GetVnetbyNameE(t *testing.T, resGroupName string, vNetName string, subscriptionID string) (network.VirtualNetwork, error) {
	vnet := network.VirtualNetwork{}

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return vnet, err
	}

	// Create a VNet client
	vnetClient, err := GetVirtualNetworkClient(subscriptionID)
	if err != nil {
		return vnet, err
	}

	// Get the details of the Virtual Network
	vnet, err = vnetClient.Get(context.Background(), resGroupName, vNetName, "")
	if err != nil {
		return vnet, err
	}
	return vnet, nil
}

// GetSubnetbyName gets propteries of the Azure Subnet by its given name
func GetSubnetbyName(t *testing.T, resGroupName string, vNetName string, sNetName string, subscriptionID string) network.Subnet {
	snet, err := GetSubnetbyNameE(t, resGroupName, vNetName, sNetName, subscriptionID)
	require.NoError(t, err)

	return snet
}

// GetSubnetbyName gets propteries of the Azure Subnet by its given name
func GetSubnetbyNameE(t *testing.T, resGroupName string, vNetName string, sNetName string, subscriptionID string) (network.Subnet, error) {
	snet := network.Subnet{}

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return snet, err
	}

	// Create a VNet client
	snetClient, err := GetSubnetsClient(subscriptionID)
	if err != nil {
		return snet, err
	}

	// Get the details of the Subnet
	snet, err = snetClient.Get(context.Background(), resGroupName, vNetName, sNetName, "")
	if err != nil {
		return snet, err
	}
	return snet, nil
}
