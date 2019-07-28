package httpclient

import (
	"crypto/tls"
	"net/http"
)

// HTTPClient is base interface for all HttpClient implementation.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New creates a new instance of HttpClients.
// The allowInsecureConnection determines if the client needs to validate connection certificate.
func New(allowInsecureConnection bool) HTTPClient {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: allowInsecureConnection,
			},
		},
	}
}
