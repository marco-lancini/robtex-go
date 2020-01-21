package robtex

// ---------------------------------------------------------------------------------------
// IP
// ---------------------------------------------------------------------------------------
type act struct {
	O string `json:"o"`
	T int64  `json:"t"`
}

// IPInfo is ipquery data
type IPInfo struct {
	Status    string `json:"status"`
	City      string `json:"city"`
	Country   string `json:"country"`
	As        int    `json:"as"`
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
type DNSRecord struct {
	Rrname    string `json:"rrname"`
	Rrdata    string `json:"rrdata"`
	Rrtype    string `json:"rrtype"`
	TimeFirst int64  `json:"time_first"`
	TimeLast  int64  `json:"time_last"`
	Count     int    `json:"count"`
}

type Pdns struct {
	Records []DNSRecord
}
