package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/katasec/ark/resources/azure/cloudspaces"
)

func readConfigData() cloudspaces.AzureCloudspace {

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

	return acs
}
