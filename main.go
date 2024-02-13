package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// readConfigData
func main() {
	pulumi.Run(NewDC)
	//fmt.Println(&azuredc.ReferenceHubVNET)
}
