package doluna

var HLR_LOOKUP_URL = "/hlr/v2/"

type DolunaClient struct {
	ApiHost     string
	ApiUsername string
	ApiKey      string
}

type Doluna interface {
	HlrLookup(string) (*HlrLookupResponse, error)
}

func New(apiHost string, apiUsername string, apiKey string) Doluna {
	return &DolunaClient{
		ApiHost:     apiHost,
		ApiUsername: apiUsername,
		ApiKey:      apiKey,
	}
}
