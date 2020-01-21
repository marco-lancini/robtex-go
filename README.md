# Go client for the Robtex API

This library provides a little wrapper over the Robtex API (https://www.robtex.com/api/).

## Installation

### Direct Installation

```go get -u github.com/marco-lancini/robtex-go```

### Running in a Docker Container

Git clone the repo, then run in a container with the following commands:

* Clone the repo: `git clone https://github.com/marco-lancini/robtex-go`
* Build the docker container: `docker build -t robtex-go .`
* Run the container: `docker run --rm -it robtex-go`

## Usage

This library can be imported quickly, as shown below:

```go
package main
import (
  "log"
  "github.com/marco-lancini/robtex-go/robtex"
)

func main() {
  // Use freeapi
  client := robtex.NewClient("https://freeapi.robtex.com", "", "")

  // For proapi, need load <YOUR-API-KEY>
  // client := NewClient("https://proapi.robtex.com", "", "<YOUR-API-KEY>")

  // ipquery
  ipInfo, err := client.IPQuery("199.19.54.1")
  if err != nil {
    log.Println(err)
  }
  log.Println(ipInfo)// => {"status":"ok","city":"Toronto","country":"Canada","as":12041,...

  // asquery
  asn, err := client.AsQuery(1234)
  if err != nil {
    log.Println(err)
  }
  log.Println(asn)

  // pdns forward
  passiveDNSForward, err := client.PDNSForward("a.iana-servers.net")
  if err != nil {
    log.Println(err)
  }
  log.Println(passiveDNSForward)

  // pdns reverse
  passiveDNSReverse, err := client.PDNSReverse("199.43.132.53")
  if err != nil {
    log.Println(err)
  }
  log.Println(passiveDNSReverse)
}
```
