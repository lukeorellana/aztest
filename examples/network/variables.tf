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

variable "vnet_address_space" {
  type        = string
  description = "Address space for Virtual Network"
  default     = "10.0.0.0/16"
}

variable "subnet_prefix" {
  type        = string
  description = "Prefix of subnet address"
  default     = "10.0.0.0/24"

}