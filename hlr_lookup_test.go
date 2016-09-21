package doluna_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sliide/doluna"
)

func TestDolunaClientSuccessResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseJson := `{
				"hlr_error_code": "0",
				"hlr_imsi": "23420",
				"hlr_lookupdatetime": "21/09/2016 11:47:11",
				"hlr_number": "447775372611",
				"hlr_number_location": "44",
				"hlr_operator_countrycode": "44",
				"hlr_operator_countryname": "United Kingdom",
				"hlr_operator_name": "Hutchison UK",
				"hlr_operator_networkname": "3 UK",
				"hlr_prefix_match": "Vodafone UK",
				"hlr_remaining_credit": "12.6200",
				"hlr_status": "OK",
				"hlr_trans_ref": "0264a95a-7db0-4f10-a2ff-6b69b78d48f8"
		}`

		fmt.Fprintln(w, responseJson)
	}))
	defer testServer.Close()

	dolunaClient := doluna.New(testServer.URL, "", "")
	response, err := dolunaClient.HlrLookup("447775372611")

	AssertEqual(t, err, nil)
	AssertEqual(t, response.HlrNumber, "447775372611")
}

func TestDolunaClientErrorResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseJson := `{
			"hlr_error_code": "Error : E109: Invalid api_service_key provided.  This lookup has not been submitted",
			"hlr_imsi": "-",
			"hlr_lookupdatetime": "21/09/2016 15:59:33",
			"hlr_number": null,
			"hlr_number_location": "-",
			"hlr_operator_countrycode": "-",
			"hlr_operator_countryname": "-",
			"hlr_operator_name": "-",
			"hlr_operator_networkname": "-",
			"hlr_prefix_match": "-",
			"hlr_remaining_credit": "",
			"hlr_status": "FAILED",
			"hlr_trans_ref": ""
		}`

		fmt.Fprintln(w, responseJson)
	}))
	defer testServer.Close()

	dolunaClient := doluna.New(testServer.URL, "", "")
	response, err := dolunaClient.HlrLookup("447775372611")

	AssertNotEqual(t, err, nil)
	AssertEqual(t, response.HlrErrorCode, fmt.Sprintf("%s", err.Error()))
}

func TestDolunaClientNon200StatusCode(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseJson := `{
			"hlr_error_code": "Error : E109: Invalid api_service_key provided.  This lookup has not been submitted",
			"hlr_imsi": "-",
			"hlr_lookupdatetime": "21/09/2016 15:59:33",
			"hlr_number": null,
			"hlr_number_location": "-",
			"hlr_operator_countrycode": "-",
			"hlr_operator_countryname": "-",
			"hlr_operator_name": "-",
			"hlr_operator_networkname": "-",
			"hlr_prefix_match": "-",
			"hlr_remaining_credit": "",
			"hlr_status": "FAILED",
			"hlr_trans_ref": ""
		}`
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(w, responseJson)
	}))
	defer testServer.Close()

	dolunaClient := doluna.New(testServer.URL, "", "")
	_, err := dolunaClient.HlrLookup("447775372611")
	AssertNotEqual(t, err, nil)
}

func TestServerWithHlrResponse(t *testing.T) {
	testServer := doluna.ServerWithHlrResponse(doluna.HlrLookupResponse{
		HlrStatus: "OK",
		HlrNumber: "447775372611",
	})

	dolunaClient := doluna.New(testServer.URL, "", "")
	response, err := dolunaClient.HlrLookup("447775372611")

	AssertEqual(t, err, nil)
	AssertEqual(t, response.HlrStatus, "OK")
	AssertEqual(t, response.HlrNumber, "447775372611")
}
