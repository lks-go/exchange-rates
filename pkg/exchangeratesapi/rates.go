package exchangeratesapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) Rates(symbols... string) (*RatesResponse, error) {
	res, err := http.Get(fmt.Sprintf(convertURL, c.cfg.AccessKey, strings.Join(symbols, ",")))
	if err != nil {
		return nil, fmt.Errorf("failed to get rates: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get rates: wrong status: %d %s", res.StatusCode, res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	rates := &RatesResponse{}
	if err := json.Unmarshal(b, rates); err != nil {
		fmt.Errorf("failed to unmarshal response body: %s", err)
	}

	return rates, nil
}

type RatesResponse struct {
	Success bool `json:"success"`
	Timestamp int `json:"timestamp"`
	Base string `json:"base"`
	Date string `json:"date"`
	Rates map[string]float64 `json:"rates"`
}