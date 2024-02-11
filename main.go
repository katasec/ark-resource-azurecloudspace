package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/katasec/ark/resources/azure/cloudspaces"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	fmt.Println("hello")

	fileName := "./configdata.json"
	acsJson, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Could not find config data file: " + fileName)
		os.Exit(1)
	}

	var acs cloudspaces.AzureCloudspace

	err = json.Unmarshal(acsJson, &acs)
	if err != nil {
		fmt.Println("Could not parse json data to an Azure Cloudspace")
		os.Exit(1)
	}

	fmt.Println(acs.Hub.Name)
}

func pulumiStuff() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "myResourceGroup", &resources.ResourceGroupArgs{
			ResourceGroupName: pulumi.String("my-resource-group"),
			Location:          pulumi.String("WestUS"),
		})
		if err != nil {
			return err
		}

		// Export the name of the resource group
		ctx.Export("resourceGroupName", resourceGroup.Name)

		return nil
	})
}
