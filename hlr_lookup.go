package doluna

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type HlrLookupResponse struct {
	HlrErrorCode           string `json:"hlr_error_code"`
	HlrImsi                string `json:"hlr_imsi"`
	HlrLookupdatetime      string `json:"hlr_lookupdatetime"`
	HlrNumber              string `json:"hlr_number"`
	HlrNumberLocation      string `json:"hlr_number_location"`
	HlrOperatorCountrycode string `json:"hlr_operator_countrycode"`
	HlrOperatorCountryname string `json:"hlr_operator_countryname"`
	HlrOperatorName        string `json:"hlr_operator_name"`
	HlrOperatorNetworkname string `json:"hlr_operator_networkname"`
	HlrPrefixMatch         string `json:"hlr_prefix_match"`
	HlrRemainingCredit     string `json:"hlr_remaining_credit"`
	HlrStatus              string `json:"hlr_status"`
	HlrTransRef            string `json:"hlr_trans_ref"`
}

func (self *DolunaClient) HlrLookup(phoneNumber string) (*HlrLookupResponse, error) {
	args := url.Values{}
	args.Set("output", "json")
	args.Set("hlr_mask", "1")
	args.Add("hlr_number", phoneNumber)
	args.Add("hlr_client_ref", self.ApiUsername)
	args.Add("api_service_key", self.ApiKey)

	response, err := http.Get(self.ApiHost + HLR_LOOKUP_URL + "?" + args.Encode())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Doluna returned %d status code", response.StatusCode)
	}

	decoder := json.NewDecoder(response.Body)

	var hlrLookupResponse HlrLookupResponse

	if err := decoder.Decode(&hlrLookupResponse); err != nil {
		return nil, err
	}

	if hlrLookupResponse.HlrStatus != "OK" {
		return &hlrLookupResponse, fmt.Errorf(hlrLookupResponse.HlrErrorCode)
	}

	return &hlrLookupResponse, nil
}
