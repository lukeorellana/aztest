variable "system" {
  type        = string
  description = "Name of the system or environment"
  default     = "terratest"
}


variable "location" {
  type        = string
  description = "Azure location of terraform server environment"
  default     = "westus2"

}

variable "servername" {
  type        = string
  description = "Server name of the virtual machine"
  default     = "terraformvm"
}


variable "admin_username" {
  type        = string
  description = "Administrator username for server"
  default     = "terraadmin"
}

variable "admin_password" {
  type        = string
  description = "Administrator password for server"
  default     = "BadPassword1234"
}

variable "vnet_address_space" {
  type        = list
  description = "Address space for Virtual Network"
  default     = ["10.0.0.0/16"]
}

variable "subnet_prefix" {
  type        = string
  description = "Prefix of subnet address"
  default     = "10.0.0.0/24"

}

variable "managed_disk_type" {
  type        = string
  description = "Disk type Premium in Primary location Standard in DR location"
  default     = "Standard_LRS"

}

variable "vm_size" {
  type        = string
  description = "Size of VM"
  default     = "Standard_B1s"
}

  