# ------------------------------------------------------------
# Create Hub Resource Group
# ------------------------------------------------------------

resource "azurerm_resource_group" "hub" {
  name     = local.hub_rg_name
  location = local.location
}

# ------------------------------------------------------------
# Create Hub VNET
# ------------------------------------------------------------
resource "azurerm_virtual_network" "vnet_hub" {
  name                = local.hub_vnet_name
  location            = azurerm_resource_group.hub.location
  resource_group_name = azurerm_resource_group.hub.name
  address_space       = [local.hub_subnet_adress_space]

  subnet {
    name           = local.hub_subnet_fw_name
    address_prefix = local.hub_subnet_fw_prefix
  }

  subnet {
    name           = local.hub_subnet_bastion_name
    address_prefix = local.hub_subnet_bastion_prefix
  }

  subnet {
    name           = local.hub_subnet_fwmgmt_name
    address_prefix = local.hub_subnet_fwmgmt_prefix
    //security_group = azurerm_network_security_group.example.id
  }

  subnet {
    name           = local.hub_subnet_gateway_name
    address_prefix = local.hub_subnet_gateway_prefix
    //security_group = azurerm_network_security_group.example.id
  }

  tags = {
    environment = local.hub_env
  }
}