data "local_file" "network_json" {
  filename = "${path.module}/configdata.json" // Replace with your file path
}

variable "hub_address_prefix" {
  default = "10.16.0"
}

locals {
  location = "UAE North"
  config_data = jsondecode(data.local_file.network_json.content)

  /***
  *  Majority use cases for ACS deployments are greenfield. Given a 
  *  brownfield deployment, an Azure VPN Gateway configured for SNAT
  *  can avoid any overlapping CIDR issues.
  */

  // We use 10.x.x.x prefixes to avoid the docker CIDR range
  // which is 172.17.x.x

  hub_rg_name = "rg-hub"
  hub_vnet_name = "vnet-hub"
  hub_env = "hub"
  hub_address_prefix = "10.16.0"
  hub_subnet_adress_space = "${local.hub_address_prefix}.0/24"

  hub_subnet_fw_name = "AzureFirewallSubnet"    
  hub_subnet_fw_prefix = "${local.hub_address_prefix}.0/26"

  hub_subnet_fwmgmt_name = "AzureFirewallManagementSubnet"
  hub_subnet_fwmgmt_prefix = "${local.hub_address_prefix}.128/26"

  hub_subnet_bastion_name = "AzureBastionSubnet"
  hub_subnet_bastion_prefix = "${local.hub_address_prefix}.64/26"

  hub_subnet_gateway_name = "GatewaySubnet"
  hub_subnet_gateway_prefix = "${local.hub_address_prefix}.192/26"
}