# resource "azurerm_resource_group" "aks" {
#   name     = "rg-aks"
#   location = local.location
# }


# //snet-tier2-aks

# data "azurerm_subnet" "tier2aks" {
#   name                 = "snet-tier2-aks"
#   virtual_network_name = azurerm_virtual_network.vnet_hub.name
#   resource_group_name  = azurerm_resource_group.hub.name
# }

# resource "azurerm_kubernetes_cluster" "example" {
#   name                = "aks01"
#   location            = azurerm_resource_group.dev.location
#   resource_group_name = azurerm_resource_group.example.name
#   dns_prefix          = "toptal"

#   default_node_pool {
#     name       = "default"
#     node_count = 1
#     vm_size    = "Standard_D2_v2"
#     vnet_subnet_id = data.azurerm_subnet.tier2aks.id
#   }

#   identity {
#     type = "SystemAssigned"
#   }

#   tags = {
#     Environment = "dev"
#   }
# }
