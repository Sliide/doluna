package doluna

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func ServerWithHlrResponse(response HlrLookupResponse) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(HLR_LOOKUP_URL, func(w http.ResponseWriter, req *http.Request) {
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Write(jsonResponse)
	})
	return httptest.NewServer(mux)
}
