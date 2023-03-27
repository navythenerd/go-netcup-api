package netcup

import "encoding/json"

type Response struct {
	Status       string           `json:"status"`
	ShortMessage string           `json:"shortmessage"`
	LongMessage  string           `json:"longmessage"`
	Data         *json.RawMessage `json:"responsedata"`
}

type ResponseData struct {
	ApiSessionId string       `json:"apisessionid,omitempty"`
	DNSRecords   []*DNSRecord `json:"dnsrecords,omitempty"`
}
