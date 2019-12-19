package main

import (
	"fmt"
	"github.com/marco-lancini/robtex-go/robtex"
)

func main() {
	// Load your API Key here
	apiKey := "<YOUR-API-KEY>"
	userAgent := ""
	client := robtex.NewClient("https://freeapi.robtex.com", userAgent, apiKey)

	ipInfo := client.IpQuery("8.8.8.8")
	fmt.Println(ipInfo)

	asn := client.AsQuery(1234)
	fmt.Println(asn)

	passiveDns := client.PassiveDNS("www.google.com")
	fmt.Println(passiveDns)
}
