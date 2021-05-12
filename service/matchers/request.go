package matchers

import (
	"encoding/json"
	"github.com/a3herrera/api-test/container/logger"
	"net/http"
	"net/url"
)

func getRequest(uri string, results interface{}) {
	resp, err := http.Get(uri)
	if err != nil {
		logger.Log.Errorf("fail to retrieve information from itunes, %s", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(results)
	if err != nil {
		logger.Log.Errorf("fail to get body responseCRCIND, %s", err)
	}
}

func GET(baseURL string, relativeURL string, queryParams map[string]string, results interface{}) (string, error) {
	u, err := url.Parse(relativeURL)
	if err != nil {
		return "", err
	}

	if len(queryParams) > 0 {
		queryString := u.Query()
		for key, value := range queryParams {
			queryString.Set(key, value)
		}
		u.RawQuery = queryString.Encode()
	}
	basePath, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	fullPath := basePath.ResolveReference(u)
	resp, err := http.Get(fullPath.String())
	if err != nil {
		logger.Log.Errorf("fail to retrieve information from itunes, %s", err)
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(results)
	if err != nil {
		logger.Log.Errorf("fail to get body responseCRCIND, %s", err)
		return "", err
	}
	return fullPath.String(), nil
}
