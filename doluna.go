package doluna

var HLR_LOOKUP_URL = "/hlr/v2/sync"
var CHECK_BALANCE_URL = "/account/balance"

type DolunaClient struct {
	ApiHost     string
	ApiUsername string
	ApiKey      string
}

type Doluna interface {
	HlrLookup(string) (*HlrLookupResponse, error)
	CheckBalance() (*CheckBalanceResponse, error)
}

func New(apiHost string, apiUsername string, apiKey string) Doluna {
	return &DolunaClient{
		ApiHost:     apiHost,
		ApiUsername: apiUsername,
		ApiKey:      apiKey,
	}
}
