

# # ------------------------------------------------------------
# # Create Dev Resource Group
# # ------------------------------------------------------------

# resource "azurerm_resource_group" "dev" {
#   name     = "rg-dev"
#   location = local.location
# }


# # ------------------------------------------------------------
# # Create Dev VNET
# # ------------------------------------------------------------
# resource "azurerm_virtual_network" "vnet_dev" {
#   name                = "vnet-dev"
#   location            = azurerm_resource_group.dev.location
#   resource_group_name = azurerm_resource_group.dev.name
#   address_space       = ["10.17.0.0/16"]

#   subnet {
#     name           = "snet-tier1-agw"
#     address_prefix = "10.17.1.0/24"
#   }

#   subnet {
#     name           = "snet-tier1-webin"
#     address_prefix = "10.17.2.0/24"
#   }

#   subnet {
#     name           = "snet-tier1-rsvd1"
#     address_prefix = "10.17.3.0/25"
#   }

#   subnet {
#     name           = "snet-tier1-rsvd2"
#     address_prefix = "10.17.3.128/25"
#   }

#   subnet {
#     name           = "snet-tier2-wbapp"
#     address_prefix = "10.17.4.0/23"
#   }

#   subnet {
#     name           = "snet-tier2-rsvd2"
#     address_prefix = "10.17.6.0/24"
#   }

#   subnet {
#     name           = "snet-tier2-pckr"
#     address_prefix = "10.17.7.0/24"
#   }

#   subnet {
#     name           = "snet-tier2-vm"
#     address_prefix = "10.17.8.0/21"
#   }

#   subnet {
#     name           = "snet-tier2-aks"
#     address_prefix = "10.17.16.0/20"
#   }

#   subnet {
#     name           = "snet-tier3-mi"
#     address_prefix = "10.17.32.0/26"
#   }

#   subnet {
#     name           = "snet-tier3-dbaz"
#     address_prefix = "10.17.32.64/26"
#   }

#   subnet {
#     name           = "snet-tier3-dblb"
#     address_prefix = "10.17.32.128/25"
#   }

#   subnet {
#     name           = "snet-tier3-dbvm"
#     address_prefix = "10.17.33.0/25"
#   }

#   subnet {
#     name           = "snet-tier3-strg"
#     address_prefix = "10.17.33.128/25"
#   }

#   subnet {
#     name           = "snet-tier3-redis"
#     address_prefix = "10.17.34.0/25"
#   }

#   tags = {
#     environment = "hub"
#   }
# }