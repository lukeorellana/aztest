output "vnet_rg" {
  description = "Location of vnet"
  value       = azurerm_virtual_network.vnet.resource_group_name
}

output "vnet_name" {
  description = "Location of vnet"
  value       = azurerm_virtual_network.vnet.name
}

output "subnet_id" {
  description = "Subnet ID"
  value       = azurerm_subnet.subnet.id
}

output "nsg_name" {
  description = "Subnet ID"
  value       = azurerm_network_security_group.nsg.name
}
