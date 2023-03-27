package netcup

type DNSRecord struct {
	Hostname    string
	Type        string
	Destination string
	DomainId    string
	Priority    int
}
