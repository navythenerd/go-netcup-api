package netcup

const (
	ActionLogin            = "login"
	ActionLogout           = "logout"
	ActionGetDNSRecords    = "infoDnsRecords"
	ActionUpdateDNSRecords = "updateDnsRecords"
)

type Request struct {
	Action string `json:"action"`

	Param struct {
		ApiKey         string `json:"apikey"`
		ApiPassword    string `json:"apipassword,omitempty"`
		ApiSessionId   string `json:"apisessionid,omitempty"`
		CustomerNumber string `json:"customernumber"`
		DomainName     string `json:"domainname"`
		DNSRecordSet   struct {
			Records []*DNSRecord `json:"dnsrecords"`
		} `json:"dnsrecordset,omitempty"`
	} `json:"param"`
}
