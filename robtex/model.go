package robtex

// ---------------------------------------------------------------------------------------
// IP
// ---------------------------------------------------------------------------------------
type act struct {
	O string `json:"o"`
	T int64  `json:"t"`
}

type IpInfo struct {
	Status    string `json:"status"`
	City      string `json:"city"`
	Country   string `json:"country"`
	As        int `json:"as"`
	Asname    string `json:"asname"`
	Whoisdesc string `json:"whoisdesc"`
	Routedesc string `json:"routedesc"`
	Bgproute  string `json:"bgproute"`
	Act       []act  `json:"act"`
	Acth      []act  `json:"acth"`
	Pas       []act  `json:"pas"`
	Pash      []act  `json:"pash"`
}

// ---------------------------------------------------------------------------------------
// ASN
// ---------------------------------------------------------------------------------------
type net struct {
	N     string `json:"n"`
	InBgp int    `json:"inbgp"`
}

type ASN struct {
	Status string `json:"status"`
	Nets   []net  `json:"nets"`
}

// ---------------------------------------------------------------------------------------
// PDNS
// ---------------------------------------------------------------------------------------
type DnsRecord struct {
	Rrname     string `json:"rrname"`
	Rrdata     string `json:"rrdata"`
	Rrtype     string `json:"rrtype"`
	Time_first int64  `json:"time_first"`
	Time_last  int64  `json:"time_last"`
	Count      int    `json:"count"`
}

type Pdns struct {
	Records []DnsRecord
}
