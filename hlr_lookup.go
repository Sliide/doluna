package doluna

import (
	"encoding/json"
	"errors"
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

var ErrHttpRequest = errors.New("Failed Http Request.")
var ErrNon200Status = errors.New("Doluna returned non 200 Http Status Code.")
var ErrJsonDecode = errors.New("Failed to decode json response.")
var ErrHlrFailed = errors.New("HLR lookup Failed.")

func (self *DolunaClient) HlrLookup(phoneNumber string) (*HlrLookupResponse, error) {
	args := url.Values{}
	args.Set("output", "json")
	args.Set("hlr_mask", "1")
	args.Add("hlr_number", phoneNumber)
	args.Add("hlr_client_ref", self.ApiUsername)
	args.Add("api_service_key", self.ApiKey)

	response, err := http.Get(self.ApiHost + HLR_LOOKUP_URL + "?" + args.Encode())
	if err != nil {
		return nil, ErrHttpRequest
	}

	if response.StatusCode != 200 {
		return nil, ErrNon200Status
	}

	var hlrLookupResponse HlrLookupResponse

	if err := json.NewDecoder(response.Body).Decode(&hlrLookupResponse); err != nil {
		return nil, ErrJsonDecode
	}

	if hlrLookupResponse.HlrStatus != "OK" {
		return &hlrLookupResponse, ErrHlrFailed
	}

	return &hlrLookupResponse, nil
}
