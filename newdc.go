package main

import (
	"fmt"

	"github.com/katasec/ark/resources/azure/cloudspaces"
	"github.com/katasec/playground/utils"
	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	configdata = readConfigData()
)

// NewDC creates a new data centre based on a reference azuredc
func NewDC(ctx *pulumi.Context) error {

	// Create hub resource group and VNET
	hubrg, err := resources.NewResourceGroup(ctx, configdata.Hub.ResourceGroup, &resources.ResourceGroupArgs{
		Location: pulumi.String(configdata.Location),
	})
	utils.ExitOnError(err)
	hubVnet, firewall := CreateHub(ctx, hubrg, configdata.Hub)

	//CreateHub(ctx, hubrg, configdata.Hub)

	// Create some spokes
	for _, spoke := range configdata.Spokes {
		AddSpoke(ctx, spoke.Name, hubrg, hubVnet, firewall, spoke)
	}
	//rg1, vnet1 := AddSpoke(ctx, "nprod", hubrg, hubVnet, firewall, 0)
	// rg2, vnet2 := AddSpoke(ctx, "prod", hubrg, hubVnet, firewall, 1)
	//AddSpoke(ctx, "prod", hubrg, hubVnet, firewall, 1)
	// AddSpoke(ctx, "nprod2", hubrg, hubVnet, firewall, 2)

	return err
}

// Creates an Azure Virtual Network and subnets using the provided VNETInfo
func CreateHub(ctx *pulumi.Context, rg *resources.ResourceGroup, vnetInfo cloudspaces.VNETInfo) (*network.VirtualNetwork, *network.AzureFirewall) {

	// Generate list of subnets to create
	subnets := network.SubnetTypeArray{}
	for _, subnet := range vnetInfo.SubnetsInfo {
		subnets = append(subnets, network.SubnetTypeArgs{
			AddressPrefix: pulumi.String(subnet.AddressPrefixes),
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
	firewall := createFirewall(ctx, rg, vnet)
	utils.ExitOnError(err)

	return vnet, firewall
}

func AddSpoke(ctx *pulumi.Context, spokeName string, hubrg *resources.ResourceGroup, hubVnet *network.VirtualNetwork, firewall *network.AzureFirewall, spoke cloudspaces.VNETInfo) (*resources.ResourceGroup, *network.VirtualNetwork) {

	// Create a resource group
	rg, err := resources.NewResourceGroup(ctx, spoke.ResourceGroup, &resources.ResourceGroupArgs{
		Location: hubrg.Location,
	})
	utils.ExitOnError(err)

	// Create a route to firewall
	routeName := fmt.Sprintf("rt-%s", spokeName)
	spokeRoute := createFWRoute(ctx, rg, routeName, firewall)

	// Create VNET using generated ResourceGroup, Route & CIDRs
	vnet := CreateVNET(ctx, rg, &spoke, spokeRoute)

	// Peer hub with spoke
	pulumiUrn1 := fmt.Sprintf("hub-to-%s", spokeName)
	peerNetworks(ctx, pulumiUrn1, hubrg, hubVnet, vnet)

	// Peer spoke with hub
	pulumiUrn2 := fmt.Sprintf("%s-to-hub", spokeName)
	peerNetworks(ctx, pulumiUrn2, rg, vnet, hubVnet)

	return rg, vnet

}
