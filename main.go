package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/netbox-community/go-netbox/netbox"
)

var (
	publicNetboxURI = "https://demo.netbox.dev/"
)

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
