package matchers

import (
	"encoding/json"
	"github.com/a3herrera/api-test/container/logger"
	"net/http"
)

func getRequest(uri string, results interface{}) {
	resp, err := http.Get(uri)
	if err != nil {
		logger.Log.Errorf("fail to retrieve information from itunes, %s", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(results)
	if err != nil {
		logger.Log.Errorf("fail to get body response, %s", err)
	}
}
