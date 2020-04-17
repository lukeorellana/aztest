# aztest
Aztest is a Go Library for Testing Azure Resources mainly used for Terratest. This code has been forked from Terratest and then expanded upon to include more functions for testing Azure Resources. 

Below are examples on how to test each Azure resource:


### Virtual Machine

##### Correct VM Size
```
// Look up the size of the given Virtual Machine
actualVMSize := azure.GetSizeOfVirtualMachine(t, resourceGroupName, vmName,  "")
expectedVMSize := compute.VirtualMachineSizeTypes("Standard_B1s")

// Test that the Virtual Machine size matches the Terraform specification
assert.Equal(t, expectedVMSize, actualVMSize, "Check Size of VM")
```

##### Correct VM Disk Type On All Disks
```
// Lookup Disk Types attached to a Virtual Machine
listVMDiskTypes := azure.GetTypeOfVirtualMachineDisks(t,  "vmName",  "")

// Check if Virtual Machine does not have any Premium_LRS Disks attached
assert.NotContains(t, listVMDiskTypes, "Premium_LRS", "Check for correct Disk Type")
```

##### Correct Number Of Disks Attached to VM
```
// Lookup Disk Types attached to a Virtual Machine
listVMDiskTypes := azure.GetTypeOfVirtualMachineDisks(t,  "vmName",  "")

// Count the Number of Disks Attached to the Virtual Machine and check if there are the correct number
assert.Equal(t, 4, len(vmExtProperties), "Check for correct number of disks attached to VM")
```

##### Boot Diagnostics Enabled
```
// Lookup Virtual Machine properties by specifying the Virtual Machine name and Resource Group
vmProperties := azure.GetVMbyName(t, "resourceGroupName", "vmName, "")

// Test is Boot Diagnostics is enabled on the Virtual Machine
assert.True(t, *vmProperties.VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.Enabled, "Check if Boot Diagnostics is enabled")
```

##### VM Provisioning State Succeeded

```
// Lookup Virtual Machine properties by specifying the Virtual Machine name and Resource Group
vmProperties := azure.GetVMbyName(t, "resourceGroupName", "vmName, "")

// Test if VM Provisioned with Succeeded status
assert.Equal(t, "Succeeded", *vmProperties.VirtualMachineProperties.ProvisioningState, "Check if VM Provisioned successfully")
```

##### Virtual Machine Extension Provisioned Successfully
```
// Lookup Virtual Machine Extension properties by specifying the Virtual Machine name, Resource Group, and VM Extension Name
vmProperties := azure.GetVirtualMachineExt(t, "resourceGroupName", "vmName, "CustomScriptExtension", "")

// Test if VM Extension Provisioned with Succeeded status
assert.Equal(t, "Succeeded", *vmExtProperties.ProvisioningState, "Check for CustomScript Extension succeeded")
```

##### Check For Windows Bring Your Own License
```
// Lookup Virtual Machine properties by specifying the Virtual Machine name and Resource Group
vmProperties := azure.GetVMbyName(t, "resourceGroupName", "vmName, "")

// Test if VM has BYOL configured. Only works on Windows Servers
assert.Equal(t, "Windows_Server", *vmExtProperties.VirtualMachineProperties.LicenseType, "Check for BYOL")
```

##### Check If VM NIC Is Assigned To NSG
```
// Lookup Virtual Machine properties by specifying the Virtual Machine name and Resource Group
vmProperties := azure.GetVMbyName(t, "resourceGroupName", "vmName, "")

// Look up Subnet and NIC ID associations of NSG
nsgAssociations := azure.GetAssociationsforNSG(t, nsgName, vnetRG, "")

// For each NIC on Virtual Machine, check that it is assigned to the desired NSG
for _, NIC := range *vmProperties.NetworkProfile.NetworkInterfaces {
	assert.Contains(t, nsgAssociations, *NIC.ID, "Check if VM NIC is assigned to NSG")
}
```


### Networking

##### Subnet Is Assigned To NSG

```
// Look up Subnet and NIC ID associations of NSG
nsgAssociations := azure.GetAssociationsforNSG(t, nsgName, vnetRG, "")

//Check if subnet is associated wtih NSG
assert.Contains(t, nsgAssociations, subnetID, "Check if subnet is assigned to NSG")
```


##### VNet Exists
```
// Look up Virtual Network by Name
GetVnetbyName(t, resGroupName, vNetName, "")

```
##### Subnet Exists In Virtual Network
```
// Look up all subnet IDs from the Virtual Network Name
subnets := azure.GetSubnetsforVnet(t, resGroupName, vnetName, "")

//Check if the subnet exists in the Virtual Network
assert.Contains(t, subnets, subnetID, "Check if subnet exists in virutal network")
```
