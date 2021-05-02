package internal

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// querystring

const defaultBaseURL = "https://api.dynadot.com/api3.xml"

const (
	commandSetDNS       = "set_dns2"
	commandIsProcessing = "is_processing"
)

const codeSuccess = 0
const msgOK = "yes"

type Client struct {
	key string

	baseURL    string
	HTTPClient *http.Client
}

func NewClient(key string) *Client {
	return &Client{
		key:        key,
		baseURL:    defaultBaseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) SetDNS(domain, subDomain, recordType, recordValue string, ttl int) (*SetDNSResponse, error) {
	req := SetDNSRequest{
		Key:            c.key,
		Command:        commandSetDNS,
		Domain:         domain,
		SubDomain0:     subDomain,
		SubRecordType0: recordType,
		SubRecord0:     recordValue,
		TTL:            ttl,
	}

	endpoint, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}

	values, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	endpoint.RawQuery = values.Encode()

	fmt.Println(endpoint)

	resp, err := c.HTTPClient.Get(endpoint.String())
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(respBody))
	}

	var setDNSResp SetDNSResponse
	err = xml.Unmarshal(respBody, &setDNSResp)
	if err != nil {
		var ue xml.UnmarshalError
		if errors.As(err, &ue) {
			var apiResp Response
			errResponse := xml.Unmarshal(respBody, &apiResp)
			if errResponse != nil {
				return nil, fmt.Errorf("failed to unmarshal response: %w", err)
			}

			return nil, fmt.Errorf("error: %s (%d) %s", apiResp.Error, apiResp.ResponseCode, apiResp.ResponseMsg)
		}

		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if setDNSResp.SuccessCode != codeSuccess {
		return nil, fmt.Errorf("error: %s (%d) %s", setDNSResp.Status, setDNSResp.SuccessCode, setDNSResp.Error)
	}

	return &setDNSResp, nil
}

func (c *Client) IsProcessing() (bool, error) {
	req := IsProcessingRequest{
		Key:     c.key,
		Command: commandIsProcessing,
	}

	endpoint, err := url.Parse(c.baseURL)
	if err != nil {
		return false, err
	}

	values, err := query.Values(req)
	if err != nil {
		return false, err
	}

	endpoint.RawQuery = values.Encode()

	fmt.Println(endpoint)

	resp, err := c.HTTPClient.Get(endpoint.String())
	if err != nil {
		return false, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("%d: %s", resp.StatusCode, string(respBody))
	}

	var isProcResp Response
	err = xml.Unmarshal(respBody, &isProcResp)
	if err != nil {
		return false, err
	}

	if isProcResp.ResponseCode != codeSuccess {
		return false, fmt.Errorf("error: %s (%d) %s", isProcResp.Error, isProcResp.ResponseCode, isProcResp.ResponseMsg)
	}

	return isProcResp.ResponseMsg == msgOK, nil
}
