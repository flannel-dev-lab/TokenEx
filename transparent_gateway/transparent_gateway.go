// Package transparent_gateway handles sending detokenization requests to tokenex
package transparent_gateway

import (
	"bytes"
	"net/http"
)

// DetokenizeObject Contains all the fields needed to push data to tokenex
type DetokenizeObject struct {
	ContentType            string
	DetokenizeURL          string
	TokenExAPIKey          string
	TokenExID              string
	ExternalRequestURL     string
	ExternalRequestBody    []byte
	ExternalRequestHeaders map[string]string
}

// TransparentGatewayDetokenizeRequest Sends a request to tokenex for detokenization
func (detokenizeObject *DetokenizeObject) TransparentGatewayDetokenizeRequest() (*http.Response, error) {
	tokenExAPIRequest, err := http.NewRequest("POST", detokenizeObject.DetokenizeURL, bytes.NewBuffer(detokenizeObject.ExternalRequestBody))
	if err != nil {
		return nil, err
	}

	tokenExAPIRequest.Header.Set("Content-Type", detokenizeObject.ContentType)
	tokenExAPIRequest.Header.Set("TX_URL", detokenizeObject.ExternalRequestURL)
	tokenExAPIRequest.Header.Set("TX_APIKey", detokenizeObject.TokenExAPIKey)
	tokenExAPIRequest.Header.Set("TX_TokenExID", detokenizeObject.TokenExID)

	for headerKey, headerValue := range detokenizeObject.ExternalRequestHeaders {
		tokenExAPIRequest.Header.Set(headerKey, headerValue)
	}

	client := &http.Client{}
	return client.Do(tokenExAPIRequest)
}
