package matchers

import (
	"encoding/xml"
	"github.com/a3herrera/api-test/container/logger"
	"github.com/a3herrera/api-test/service/search"
	"github.com/tiaguinho/gosoap"
	"net/http"
	"time"
)

type crcindMatcher struct {
	client string
}

func init() {
	matcher := crcindMatcher{
		client: "http://www.crcind.com/csp/samples/SOAP.Demo.cls?WSDL",
	}
	search.Register("crcind", matcher)
}

type (
	sqlCRCIND struct {
		ID   int    `xml:"ID" json:"ID"`
		Name string `xml:"Name" json:"Name"`
		DOB  string `xml:"DOB" json:"DOB"`
		SSN  string `xml:"SSN" json:"SSN"`
	}

	listCRCIND struct {
		SQL []sqlCRCIND `xml:"SQL"'`
	}

	diffCRCIND struct {
		XMLName xml.Name   `xml:"urn:schemas-microsoft-com:xml-diffgram-v1 diffgram"`
		List    listCRCIND `xml:"ListByName"`
	}

	resultCRCIND struct {
		XMLName xml.Name `xml:"GetByNameResult"`
		Diff    diffCRCIND
	}

	responseCRCIND struct {
		XMLName xml.Name `xml:"GetByNameResponse"`
		Result  resultCRCIND
	}
)

func (m crcindMatcher) Search(searchValue string) (*search.Result, error) {
	logger.Log.Info("Start crcind search")
	httpClient := &http.Client{
		Timeout: 2500 * time.Millisecond,
	}

	soap, err := gosoap.SoapClient(m.client, httpClient)
	if err != nil {
		logger.Log.Errorf("SoapClient error: %s", err)
		return nil, err
	}
	params := gosoap.Params{
		"name": searchValue,
	}

	res, err := soap.Call("GetByName", params)
	if err != nil {
		logger.Log.Errorf("Call error: %s", err)
		return nil, err
	}
	var r responseCRCIND
	err = res.Unmarshal(&r)
	if err != nil {
		return nil, err
	}

	searchResult := search.Result{
		URI:    m.client,
		Exists: false,
	}

	if len(r.Result.Diff.List.SQL) > 0 {
		searchResult.Results = r.Result.Diff.List.SQL
		searchResult.Exists = true
	}
	return &searchResult, nil
}
