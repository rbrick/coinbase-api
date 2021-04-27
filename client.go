package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// The base endpoint for the Coinbase API
	COINBASE_API_ENDPOINT = "https://api.coinbase.com/v2/"

	// The header for the access key
	CB_ACCESS_KEY_HEADER = "CB-ACCESS-KEY"
	// The header for the signature
	CB_ACCESS_SIGN_HEADER = "CB-ACCESS-SIGN"
	// The header for the access time stamp
	CB_ACCESS_TIMESTAMP = "CB-ACCESS-TIMESTAMP"
)

//Client is the main client for interacting with the Coinbase API
type Client struct {
	//ApiKey and ApiSecret are optionally used to authenticate
	// for your account
	ApiKey    string
	ApiSecret string

	// HttpClient is the current HTTP client
	httpClient *http.Client
}

//signRequest signs the current request if needed
// This adds the needed functions
func (c *Client) signRequest(req *http.Request) (err error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))

	req.Header.Add(CB_ACCESS_KEY_HEADER, c.ApiKey)
	req.Header.Add(CB_ACCESS_TIMESTAMP, timestamp)

	prehash := timestamp + strings.ToUpper(req.Method) + req.URL.Path

	b, err := ioutil.ReadAll(req.Body)

	if err != nil && err != io.EOF {
		// unexpected error
		return err
	}

	if err == nil {
		// if there was an error it was EOF meaning no body
		// if the error was null, assume there is a body
		prehash += string(b)
	}

	hmacHasher := hmac.New(sha256.New, []byte(c.ApiSecret))

	req.Header.Add(CB_ACCESS_SIGN_HEADER, string(hmacHasher.Sum([]byte(prehash))))

	return nil
}

func (c *Client) makeRequest(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, path, body)

	if err != nil {
		return nil, err
	}

	if c.ApiKey != "" && c.ApiSecret != "" {
		// sign this request
		e := c.signRequest(req)

		if e != nil {
			return nil, e
		}
	}

	response, reqError := c.httpClient.Do(req)

	if reqError != nil {
		return nil, reqError
	}

	return response, nil
}

//New creates a new client using an API key and API secret
func New(apiKey, apiSecret string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
	}
}
