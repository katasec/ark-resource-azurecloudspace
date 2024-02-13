package main

import (
	"ark-resource-azurecloudspace/azuredc"

	"github.com/katasec/playground/utils"
	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	launchVmFlag      = false
	launchBastionFlag = false
	launchK8sFlag     = true
	configdata        = readConfigData()
)

// NewDC creates a new data centre based on a reference azuredc
func NewDC(ctx *pulumi.Context) error {

	// Create hub resource group and VNET
	hubrg, err := resources.NewResourceGroup(ctx, configdata.Hub.ResourceGroup, &resources.ResourceGroupArgs{
		Location: pulumi.String(configdata.Location),
	})
	utils.ExitOnError(err)
	//hubVnet, firewall := CreateHub(ctx, hubrg, &azuredc.ReferenceHubVNET)

	CreateHub(ctx, hubrg, &azuredc.ReferenceHubVNET)
	// Create some spokes
	//rg1, vnet1 := AddSpoke(ctx, "nprod", hubrg, hubVnet, firewall, 0)
	// rg2, vnet2 := AddSpoke(ctx, "prod", hubrg, hubVnet, firewall, 1)
	//AddSpoke(ctx, "prod", hubrg, hubVnet, firewall, 1)
	// AddSpoke(ctx, "nprod2", hubrg, hubVnet, firewall, 2)

	return err
}

// Creates an Azure Virtual Network and subnets using the provided VNETInfo
func CreateHub(ctx *pulumi.Context, rg *resources.ResourceGroup, vnetInfo *azuredc.VNETInfo) (*network.VirtualNetwork, *network.AzureFirewall) {

	// Generate list of subnets to create
	subnets := network.SubnetTypeArray{}
	for _, subnet := range vnetInfo.SubnetsInfo {
		subnets = append(subnets, network.SubnetTypeArgs{
			AddressPrefix: pulumi.String(subnet.AddressPrefix),
			Name:          pulumi.String(subnet.Name),
		})
	}

	// Create VNET + subnets
	vnet, err := network.NewVirtualNetwork(ctx, vnetInfo.Name, &network.VirtualNetworkArgs{
		AddressSpace: &network.AddressSpaceArgs{
			AddressPrefixes: pulumi.StringArray{
				pulumi.String(vnetInfo.AddressPrefix),
			},
		},
		ResourceGroupName: rg.Name,
		Subnets:           &subnets,
	})

	// Create Firewall
	//firewall := createFirewall(ctx, rg, vnet)
	utils.ExitOnError(err)

	firewall := &network.AzureFirewall{}
	return vnet, firewall
}
