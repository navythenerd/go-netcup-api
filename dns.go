package netcup

type DNSRecord struct {
	Id          string `json:"id"`
	Hostname    string `json:"hostname"`
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Priority    string `json:"priority"`
	State       string `json:"state"`
	Delete      bool   `json:"deleterecord"`
}
