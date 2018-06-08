# Go client for the Robtex API

This library provides a little wrapper over the Robtex API (https://www.robtex.com/api/)


# Installation

#### Direct Installation

```go get -u github.com/marco-lancini/robtex-go```


#### Running in a Docker Container

Git clone the repo, then run in a container with the following commands:

* Clone the repo: `git clone https://github.com/marco-lancini/robtex-go`
* Build the docker container: `docker build -t robtex-go .`
* Run the container: `docker run --rm -it robtex-go`



# Usage

This library can be imported quickly, as shown below:

```go
package main

import (
	"fmt"
	"github.com/marco-lancini/robtex-go"
)


func main() {
	
	client := robtex.NewClient("https://freeapi.robtex.com", "")

	ipInfo := client.IpQuery("8.8.8.8")
	fmt.Println(ipInfo)

	asn := client.AsQuery(1234)
	fmt.Println(asn)

	passiveDns := client.PassiveDNS("www.google.com")
	fmt.Println(passiveDns)

}
```