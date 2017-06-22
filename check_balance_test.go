package doluna_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sliide/doluna"
)

func TestDolunaClientCheckBalanceCreditsResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseJson := `{
			"billing_mode":"Credits",
			"decimal_balance":null,
			"sms_credits":"4714",
			"hlr_credits":"897",
			"tps_credits":"999549"}
		}`

		fmt.Fprintln(w, responseJson)
	}))
	defer testServer.Close()
	dolunaClient := doluna.New(testServer.URL, "", "")
	response, err := dolunaClient.CheckBalance()

	AssertEqual(t, err, nil)
	AssertEqual(t, response.BillingMode, "Credits")
	AssertEqual(t, response.SmsCredits, 4714)
	AssertEqual(t, response.HlrCredits, 897)
	AssertEqual(t, response.TpsCredits, 999549)
}

func TestDolunaClientCheckBalanceFinancialResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseJson := `{
			"billing_mode":"Money",
			"decimal_balance":"493.6500",
			"sms_credits":"-",
			"hlr_credits":"-",
			"tps_credits":"-"
		}`

		fmt.Fprintln(w, responseJson)
	}))
	defer testServer.Close()

	dolunaClient := doluna.New(testServer.URL, "", "")
	response, err := dolunaClient.CheckBalance()

	AssertEqual(t, err, nil)
	AssertEqual(t, response.BillingMode, "Money")
	AssertEqual(t, response.DecimalBalance, 493.65)
}
