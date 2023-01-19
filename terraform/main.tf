provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "resource_group" {
  name     = "terratest-storage-rg-${var.postfix}"
  location = var.location
}

resource "azurerm_storage_account" "storage_account" {
  name                     = "storage${var.postfix}"
  resource_group_name      = azurerm_resource_group.resource_group.name
  location                 = azurerm_resource_group.resource_group.location
  account_kind             = var.storage_account_kind
  account_tier             = var.storage_account_tier
  account_replication_type = var.storage_replication_type
  blob_properties          {
    versioning_enabled     =  var.enable_storage_version
  }
}
