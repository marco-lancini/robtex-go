package robtex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ---------------------------------------------------------------------------------------
// CLIENT
// ---------------------------------------------------------------------------------------
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
	APIKey		 string
}

// Constructor
func NewClient(baseurl string, ua string, key string) *Client {
	if ua == "" {
		ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36"
	}
	u, _ := url.Parse(baseurl)
	return &Client{
		BaseURL:    u,
		UserAgent:  ua,
		httpClient: &http.Client{},
		APIKey:			key,
	}
}

func (c *Client) newRequest(method, path string, body interface{}) *http.Request {
	// Compose URL
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	// Write body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error while encoding request body: %s", err))
		}
	}

	// New HTTP GET request
	log.Printf("Calling: %s", u.String())
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while creating the request: %s", err))
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	// Add api key to query
	if c.APIKey != "" {
		q := url.Values{}
		q.Add("key", c.APIKey)
		req.URL.RawQuery = q.Encode()
	}

	return req
}

func (c *Client) do(req *http.Request) string {
	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while performing the request: %s", err))
	}
	defer resp.Body.Close()

	// Parse response body
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while parsing the response body: %s", err))
	}

	return string(contents)
}

// ---------------------------------------------------------------------------------------
// API ENDPOINTS
// ---------------------------------------------------------------------------------------
func (c *Client) IpQuery(ip string) IpInfo {
	target_url := fmt.Sprintf("/ipquery/%s", ip)
	req := c.newRequest("GET", target_url, nil)
	resp := c.do(req)

	var result IpInfo
	err := json.Unmarshal([]byte(resp), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (c *Client) AsQuery(asn int) ASN {
	target_url := fmt.Sprintf("/asquery/%d", asn)
	req := c.newRequest("GET", target_url, nil)
	resp := c.do(req)

	var result ASN
	err := json.Unmarshal([]byte(resp), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (c *Client) PassiveDNS(domain string) Pdns {
	target_url := fmt.Sprintf("/pdns/forward/%s", domain)
	req := c.newRequest("GET", target_url, nil)
	resp := c.do(req)

	result := Pdns{}
	items := strings.Split(resp, "\n")
	for _, i := range items {
		if i != "" {
			var temp DnsRecord
			err := json.Unmarshal([]byte(i), &temp)
			if err != nil {
				log.Fatal(err)
			}
			result.Records = append(result.Records, temp)
		}
	}

	return result
}
