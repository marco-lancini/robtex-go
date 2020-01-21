package robtex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"net/url"
	"strings"
)

// Client for interacting with Robtex API.
type Client struct {
	BaseURL *url.URL
	// APIKey is the Robtex key that identifies the user making the requests.
	APIKey string
	// Agent is a string included in the User-Agent header of every request
	Agent      string
	httpClient *http.Client
}

// NewClient is constructor
func NewClient(baseurl string, agent string, key string) *Client {
	if agent == "" {
		agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Firefox/58.0.1"
	}
	u, _ := url.Parse(baseurl)
	return &Client{
		BaseURL:    u,
		Agent:      agent,
		httpClient: &http.Client{},
		APIKey:     key,
	}
}

// sendRequest sends a HTTP request to the Robtex API.
func (c *Client) sendRequest(method, path string, body io.Reader) (*http.Response, error) {
	// Compose URL
	rel := &url.URL{Path: path}
	targetURL := c.BaseURL.ResolveReference(rel)

	// Write body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// New HTTP GET request

	req, err := http.NewRequest(method, targetURL.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.Agent)

	// Add api key to query
	if c.APIKey != "" {
		q := url.Values{}
		q.Add("key", c.APIKey)
		req.URL.RawQuery = q.Encode()
	}
	// log.Printf("Doing request: %s", targetURL.String())
	return (c.httpClient).Do(req)
}

// IPQuery call api ipquery
func (c *Client) IPQuery(ip string) (*IPInfo, error) {
	path := fmt.Sprintf("/ipquery/%s", ip)
	httpResp, err := c.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	result := &IPInfo{}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AsQuery call api asquery
func (c *Client) AsQuery(asn int) (*ASN, error) {
	path := fmt.Sprintf("/asquery/%d", asn)
	httpResp, err := c.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	result := &ASN{}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PDNSForward call api Passive DNS forward
func (c *Client) PDNSForward(domain string) (*Pdns, error) {
	path := fmt.Sprintf("/pdns/forward/%s", domain)
	httpResp, err := c.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	result := &Pdns{}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	items := strings.Split(string(body), "\n")
	for _, i := range items {
		if i != "" {
			var temp DNSRecord
			err := json.Unmarshal([]byte(i), &temp)
			if err != nil {
				return nil, err
			}
			result.Records = append(result.Records, temp)
		}
	}

	return result, nil
}

// PDNSReverse call api Passive DNS reverse
func (c *Client) PDNSReverse(ip string) (*Pdns, error) {
	path := fmt.Sprintf("/pdns/reverse/%s", ip)
	httpResp, err := c.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	result := &Pdns{}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	items := strings.Split(string(body), "\n")
	for _, i := range items {
		if i != "" {
			var temp DNSRecord
			err := json.Unmarshal([]byte(i), &temp)
			if err != nil {
				return nil, err
			}
			result.Records = append(result.Records, temp)
		}
	}

	return result, nil
}
