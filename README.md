# aztest
Aztest is a Go Library for Testing Azure Resources mainly used for Terratest. This code has been forked from Terratest and then expanded upon to include more functions for testing Azure Resources. 

Below are examples on how to test each Azure resource:


### Virtual Machine

##### Correct VM Size
```
// Look up the size of the given Virtual Machine
actualVMSize := azure.GetSizeOfVirtualMachine(t, vmName, resourceGroupName, "")
expectedVMSize := compute.VirtualMachineSizeTypes("Standard_B1s")

// Test that the Virtual Machine size matches the Terraform specification
assert.Equal(t, expectedVMSize, actualVMSize)
```

##### Correct VM Disk Type on All Disks
```
// Lookup Disk Types attached to a Virtual Machine
listVMDiskTypes := azure.GetTypeOfVirtualMachineDisks(t, "vmterraform", "rg-terraexample", "")

// Check if Virtual Machine does not have any Premium_LRS Disks attached
assert.NotContains(t, listVMDiskTypes, "Premium_LRS")
```

##### Boot Diagnostics Enabled
```
// Lookup Virtual Machine properties by specifying the Virtual Machine name and Resource Group
vmProperties := azure.GetVMbyName(t, "vmterraform", "rg-terraexample", "")

// Test is Boot Diagnostics is enabled on the Virtual Machine
assert.True(t, *vmProperties.VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.Enabled)
```



### Networking

##### Subnet is assigned NSG

```
// Look up Subnet and NIC ID associations of NSG
nsgAssociations := azure.GetAssociationsforNSG(t, nsgName, vnetRG, "")

//Check if subnet is associated wtih NSG
assert.Contains(t, nsgAssociations, subnetID)
```

##### VNet Exists
```
// Look up Virtual Network by Name
GetVnetbyName(t, resGroupName, vNetName, "")

```
##### Subnet exists in VNet
```
// Look up all subnet IDs from the Virtual Network Name
subnets := azure.GetSubnetsforVnet(t, vnetName, vnetRG, "")

//Check if the subnet exists in the Virtual Network
assert.Contains(t, subnets, subnetID)
```
