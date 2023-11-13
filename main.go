package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/netbox-community/go-netbox/netbox"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
)

var (
	publicNetboxURI = "https://demo.netbox.dev/"
)

func createDcimManufacturer(nbClient *client.NetBoxAPI, vendorDevice models.Manufacturer) error {
	_, err := nbClient.Dcim.DcimManufacturersCreate(&dcim.DcimManufacturersCreateParams{
		Context: context.TODO(),
		Data:    &vendorDevice,
	}, nil)
	if err != nil {
		return fmt.Errorf("DCIM Manufactur creation failed. device: %s. error: %w", vendorDevice.Display, err)
	}
	return nil
}

func isManufacturerExist(nbClient *client.NetBoxAPI, slug string) (found bool, id int64, err error) {
	_, err = nbClient.Dcim.DcimManufacturersList(&dcim.DcimManufacturersListParams{
		Context: context.TODO(),
		SlugIe:  &slug,
	}, nil)
	if err != nil {
		return found, id, fmt.Errorf("DCIM Manufactur finding failed. slug: %v. error: %w", &slug, err)
	}
	// TODO:
	return false, id, nil
}

func main() {
	apiToken := flag.String("token", "", "netbox api token")
	netboxURL := flag.String("url", publicNetboxURI, "netbox URL")
	flag.Parse()

	if *apiToken == "" {
		log.Print("Token is empty. Stopping....\n")
		return
	}

	nb := netbox.NewNetboxWithAPIKey(*netboxURL, *apiToken)
	fmt.Println(nb)
}
