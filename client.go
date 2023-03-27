package netcup

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const netcupRestUri = "https://ccp.netcup.net/run/webservice/servers/endpoint.php?JSON"

type Client struct {
	apiKey         string
	apiPassword    string
	customerNumber string
	sessionId      string
	httpClient     http.Client
}

func NewClient(apiKey string, apiPassword string, customerNumber string) *Client {
	return &Client{
		apiKey:         apiKey,
		apiPassword:    apiPassword,
		customerNumber: customerNumber,
		httpClient:     http.Client{},
	}
}

func sendJson(data any) (*Response, error) {
	payload, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, netcupRestUri, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var decodedResponse Response
	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		return nil, err
	}

	return &decodedResponse, err
}

func unmarshalResponseData(resp *Response) (*ResponseData, error) {
	data := ResponseData{}
	err := json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) Login() error {
	req := Request{
		Action: ActionLogin,
	}

	req.Param.ApiKey = c.apiKey
	req.Param.ApiPassword = c.apiPassword
	req.Param.CustomerNumber = c.customerNumber

	resp, err := sendJson(req)

	if err != nil {
		return err
	}

	if resp.Status == "error" {
		return errors.New(resp.ShortMessage)
	}

	responseData, err := unmarshalResponseData(resp)

	if err != nil {
		return err
	}

	c.sessionId = responseData.ApiSessionId
	return nil
}

func (c *Client) Logout() error {
	req := Request{
		Action: ActionLogout,
	}

	req.Param.ApiKey = c.apiKey
	req.Param.ApiSessionId = c.sessionId
	req.Param.CustomerNumber = c.customerNumber

	resp, err := sendJson(req)

	if err != nil {
		return err
	}

	if resp.Status == "error" {
		return errors.New(resp.ShortMessage)
	}

	c.sessionId = ""
	return nil
}

func (c *Client) GetDNSRecords(domain string) ([]*DNSRecord, error) {
	req := Request{
		Action: ActionGetDNSRecords,
	}

	req.Param.ApiKey = c.apiKey
	req.Param.ApiSessionId = c.sessionId
	req.Param.CustomerNumber = c.customerNumber
	req.Param.DomainName = domain

	resp, err := sendJson(req)

	if err != nil {
		return nil, err
	}

	if resp.Status == "error" {
		return nil, errors.New(resp.ShortMessage)
	}

	responseData, err := unmarshalResponseData(resp)

	if err != nil {
		return nil, err
	}

	return responseData.DNSRecords, err
}

func (c *Client) UpdateDNSRecords(domain string, records []*DNSRecord) error {
	req := Request{
		Action: ActionUpdateDNSRecords,
	}

	req.Param.ApiKey = c.apiKey
	req.Param.ApiSessionId = c.sessionId
	req.Param.CustomerNumber = c.customerNumber
	req.Param.DomainName = domain
	req.Param.DNSRecordSet.Records = records

	resp, err := sendJson(req)

	if err != nil {
		return err
	}

	if resp.Status == "error" {
		return errors.New(resp.ShortMessage)
	}

	return nil
}
