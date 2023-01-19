variable "location" {
  description = "location of storage account to be created"
  type        = string
  default     = "East US"
}

variable "storage_account_kind" {
  description = "kind of storage account"
  type        = string
  default     = "StorageV2"
}

variable "storage_account_tier" {
  description = "tier of storage account"
  type        = string
  default     = "Standard"
}

variable "storage_replication_type" {
  description = "replication type"
  type        = string
  default     = "GRS"
}

variable "postfix" {
  description = "postfix string"
  type        = string
  default     = "test"
}

variable "enable_storage_version" {
  description = "enable versioning in storage account"
  type        = bool
  # Change the default value to false for testing if versioning is not enabled
  default     = true
}
