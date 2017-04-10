package api

import "net/http"

// Quote -
func (api *API) Quote(w http.ResponseWriter, req *http.Request) {
	quote := api.quotes.RandomQuote()
	w.Write([]byte(quote.Text))
}

// SecretQuote -
func (api *API) SecretQuote(w http.ResponseWriter, req *http.Request) {
	quote := api.quotes.RandomQuote()
	w.Write([]byte(quote.Text))
}
