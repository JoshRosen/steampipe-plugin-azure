
variable "resource_name" {
  type        = string
  default     = "turbot-test-20200125-create-update"
  description = "Name of the resource used throughout the test."
}

variable "azure_environment" {
  type        = string
  default     = "public"
  description = "Azure environment used for the test."
}

variable "azure_subscription" {
  type        = string
  default     = "3510ae4d-530b-497d-8f30-53b9616fc6c1"
  description = "Azure subscription used for the test."
}

provider "azurerm" {
  # Cannot be passed as a variable
  version         = "=1.36.0"
  environment     = var.azure_environment
  subscription_id = var.azure_subscription
}

data "azurerm_client_config" "current" {}

data "null_data_source" "resource" {
  inputs = {
    scope = "azure:///subscriptions/${data.azurerm_client_config.current.subscription_id}"
  }
}

resource "azurerm_resource_group" "named_test_resource" {
  name     = var.resource_name
  location = "East US"
}

resource "azurerm_managed_disk" "named_test_resource" {
  name                 = var.resource_name
  location             = azurerm_resource_group.named_test_resource.location
  resource_group_name  = azurerm_resource_group.named_test_resource.name
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "10"
}

resource "azurerm_snapshot" "named_test_resource" {
  name                = "${var.resource_name}snapshot"
  location            = azurerm_resource_group.named_test_resource.location
  resource_group_name = azurerm_resource_group.named_test_resource.name
  create_option       = "Copy"
  source_uri          = azurerm_managed_disk.named_test_resource.id

  tags = {
    name = var.resource_name
  }
}

output "resource_aka" {
  value = "azure://${azurerm_snapshot.named_test_resource.id}"
}

output "resource_aka_lower" {
  value = "azure://${lower(azurerm_snapshot.named_test_resource.id)}"
}

output "resource_name" {
  value = var.resource_name
}

output "resource_name_upper" {
  value = upper(var.resource_name)
}

output "resource_id" {
  value = azurerm_snapshot.named_test_resource.id
}

output "subscription_id" {
  value = var.azure_subscription
}
