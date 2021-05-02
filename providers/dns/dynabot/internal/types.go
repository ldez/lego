package internal

import "encoding/xml"

type SetDNSRequest struct {
	Key     string `url:"key"`
	Command string `url:"command"`

	Domain string `url:"domain"`

	MainRecord0     string `url:"main_record0"`
	MainRecordType0 string `url:"main_record_type0"`

	SubDomain0     string `url:"subdomain0"`
	SubRecordType0 string `url:"sub_record_type0"`
	SubRecord0     string `url:"sub_record0"`

	TTL int `url:"ttl"`
}

type SetDNSResponse struct {
	XMLName     xml.Name `xml:"SetDnsResponse"`
	SuccessCode int      `xml:"SetDnsHeader>SuccessCode"`
	Status      string   `xml:"SetDnsHeader>Status"`
	Error       string   `xml:"SetDnsHeader>Error"`
}

type IsProcessingRequest struct {
	Key     string `url:"key"`
	Command string `url:"command"`
}

type Response struct {
	Name         xml.Name `xml:"Response"`
	ResponseCode int      `xml:"ResponseHeader>ResponseCode"`
	ResponseMsg  string   `xml:"ResponseHeader>ResponseMsg"`
	Error        string   `xml:"ResponseHeader>Error"`
}
