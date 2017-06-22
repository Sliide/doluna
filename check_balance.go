package doluna

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// {"billing_mode":"Money","decimal_balance":"493.6500","sms_credits":"-","hlr_credits":"-","tps_credits":"-"}

type CheckBalanceResponse struct {
	BillingMode    string  `json:"billing_mode"`
	DecimalBalance float64 `json:"decimal_balance"`
	SmsCredits     int     `json:"sms_credits"`
	HlrCredits     int     `json:"hlr_credits"`
	TpsCredits     int     `json:"tps_credits"`
}

func (c *CheckBalanceResponse) UnmarshalJSON(b []byte) error {
	// Doluna responds all the values as strings. For this reason we
	// do a custom unmarshaling of the response in order to get the
	// right response types.

	var untypedResponse map[string]string

	if err := json.Unmarshal(b, &untypedResponse); err != nil {
		return err
	}

	if billingMode, ok := untypedResponse["billing_mode"]; ok {
		c.BillingMode = billingMode
	}

	if decimalBalance, err := strconv.ParseFloat(untypedResponse["decimal_balance"], 64); err == nil {
		c.DecimalBalance = decimalBalance
	}

	if smsCredits, err := strconv.Atoi(untypedResponse["sms_credits"]); err == nil {
		c.SmsCredits = smsCredits
	}

	if hlrCredits, err := strconv.Atoi(untypedResponse["hlr_credits"]); err == nil {
		c.HlrCredits = hlrCredits
	}

	if tpsCredits, err := strconv.Atoi(untypedResponse["tps_credits"]); err == nil {
		c.TpsCredits = tpsCredits
	}

	return nil
}

func (self *DolunaClient) CheckBalance() (*CheckBalanceResponse, error) {
	args := url.Values{}
	args.Set("output", "json")
	args.Add("api_service_key", self.ApiKey)

	response, err := http.Get(self.ApiHost + CHECK_BALANCE_URL + "?" + args.Encode())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Doluna return %s status code when checking the balance", response.StatusCode)
	}

	var checkBalanceResponse CheckBalanceResponse
	if err := json.NewDecoder(response.Body).Decode(&checkBalanceResponse); err != nil {
		return &checkBalanceResponse, err
	}

	return &checkBalanceResponse, nil
}
