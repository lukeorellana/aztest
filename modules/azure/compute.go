package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/stretchr/testify/require"
)

// GetVirtualMachineClient is a helper function that will setup an Azure Virtual Machine client on your behalf
func GetVirtualMachineClient(subscriptionID string) (*compute.VirtualMachinesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a VM client
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	vmClient.Authorizer = *authorizer

	return &vmClient, nil
}

// GetVirtualMachineExtensionClient is a helper function that will setup an Azure Virtual Machine Extension client on your behalf
func GetVirtualMachineExtensionsClient(subscriptionID string) (*compute.VirtualMachineExtensionsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a VM Extension client
	vmExtClient := compute.NewVirtualMachineExtensionsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	vmExtClient.Authorizer = *authorizer

	return &vmExtClient, nil
}

// GetSizeOfVirtualMachine gets the size type of the given Azure Virtual Machine
func GetSizeOfVirtualMachine(t *testing.T, resGroupName string, vmName string, subscriptionID string) compute.VirtualMachineSizeTypes {
	size, err := GetSizeOfVirtualMachineE(t, resGroupName, vmName, subscriptionID)
	require.NoError(t, err)

	return size
}

// GetSizeOfVirtualMachineE gets the size type of the given Azure Virtual Machine
func GetSizeOfVirtualMachineE(t *testing.T, resGroupName string, vmName string, subscriptionID string) (compute.VirtualMachineSizeTypes, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return "", err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return "", err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return "", err
	}

	return vm.VirtualMachineProperties.HardwareProfile.VMSize, nil
}

// GetTagsForVirtualMachine gets the tags of the given Virtual Machine as a map
func GetTagsForVirtualMachine(t *testing.T, resGroupName string, vmName string, subscriptionID string) map[string]string {
	tags, err := GetTagsForVirtualMachineE(t, resGroupName, vmName, subscriptionID)
	require.NoError(t, err)

	return tags
}

// GetTagsForVirtualMachineE gets the tags of the given Virtual Machine as a map
func GetTagsForVirtualMachineE(t *testing.T, resGroupName string, vmName string, subscriptionID string) (map[string]string, error) {
	// Setup a blank map to populate and return
	tags := make(map[string]string)

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return tags, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return tags, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return tags, err
	}

	// Range through existing tags and populate above map accordingly
	for k, v := range vm.Tags {
		tags[k] = *v
	}

	return tags, nil
}

// GetVMbyName gets the properties of a Virtual Machine in Azure by Name
func GetVMbyName(t *testing.T, resGroupName string, vmName string, subscriptionID string) compute.VirtualMachine {
	vm, err := GetVMbyNameE(t, resGroupName, vmName, subscriptionID)
	require.NoError(t, err)

	return vm
}

// GetVMbyName gets the properties of a Virtual Machine in Azure by Name
func GetVMbyNameE(t *testing.T, resGroupName string, vmName string, subscriptionID string) (compute.VirtualMachine, error) {
	vmProperties := compute.VirtualMachine{}

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return vmProperties, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return vmProperties, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, "")
	if err != nil {
		return vm, err
	}

	return vm, nil
}

// GetTypeOfVirtualMachineDisks gets the types of the OS and Data disks attached to the Virtual Machine
func GetTypeOfVirtualMachineDisks(t *testing.T, resGroupName string, vmName string, subscriptionID string) []string {
	size, err := GetTypeOfVirtualMachineDisksE(t, resGroupName, vmName, subscriptionID)
	require.NoError(t, err)

	return size
}

// GetTypeOfVirtualMachineDisks gets the types of the OS and Data disks attached to the Virtual Machine
func GetTypeOfVirtualMachineDisksE(t *testing.T, resGroupName string, vmName string, subscriptionID string) ([]string, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return nil, err
	}

	storageAccountTypes := []string{}

	// Add OS Disk to the string slice
	storageAccountTypes = append(storageAccountTypes, string(vm.StorageProfile.OsDisk.ManagedDisk.StorageAccountType))

	// Add all attached disks to the string slice
	for _, disk := range *vm.StorageProfile.DataDisks {

		storageAccountTypes = append(storageAccountTypes, string(disk.ManagedDisk.StorageAccountType))

	}

	return storageAccountTypes, nil
}

// GetVirtualMachineExt gets the Virtual Machine Extensions Information
func GetVirtualMachineExt(t *testing.T, resGroupName string, vmName string, vmExtName string, subscriptionID string) compute.VirtualMachineExtension {
	size, err := GetVirtualMachineExtE(t, resGroupName, vmName, vmExtName, subscriptionID)
	require.NoError(t, err)

	return size
}

// GetVirtualMachineExt gets the Virtual Machine Extensions Information
func GetVirtualMachineExtE(t *testing.T, resGroupName string, vmName string, vmExtName string, subscriptionID string) (compute.VirtualMachineExtension, error) {
	vmExtProperties := compute.VirtualMachineExtension{}

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return vmExtProperties, err
	}

	// Create a VM client
	vmExtClient, err := GetVirtualMachineExtensionsClient(subscriptionID)
	if err != nil {
		return vmExtProperties, err
	}

	// Get the details of the target virtual machine
	vm, err := vmExtClient.Get(context.Background(), resGroupName, vmName, vmExtName, "")
	if err != nil {
		return vmExtProperties, err
	}
	return vm, nil
}
