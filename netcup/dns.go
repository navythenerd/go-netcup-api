package netcup

type DNSRecord struct {
	Id          string `json:"id,omitempty"`
	Hostname    string `json:"hostname"`
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Priority    int    `json:"priority,omitempty"`
}
