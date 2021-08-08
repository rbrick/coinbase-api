package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

	var body []byte
	if req.Body != nil {
		body, err = ioutil.ReadAll(req.Body)

		if err != nil && err != io.EOF {
			// unexpected error
			return err
		}

	}

	if err == nil {
		// if there was an error it was EOF meaning no body
		// if the error was null, assume there is a body
		prehash += string(body)
	}

	hmacHasher := hmac.New(sha256.New, []byte(c.ApiSecret))

	hmacHasher.Write([]byte(prehash))

	key := hex.EncodeToString(hmacHasher.Sum(nil))

	req.Header.Add(CB_ACCESS_SIGN_HEADER, key)

	return nil
}

func (c *Client) makeRequest(method, path string, urlValues url.Values) (*http.Request, error) {
	path = COINBASE_API_ENDPOINT + path

	encoded := urlValues.Encode()
	var body io.Reader = nil

	if !hasFormBody(method) {
		path += "?" + encoded
	} else {
		body = bytes.NewReader([]byte(encoded))
	}

	req, err := http.NewRequest(method, path, body)

	if body != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))
	}

	if err != nil {
		return nil, err
	}
	return req, nil
}

// generics when?
//execute takes a new request
func (c *Client) execute(req *http.Request, decode interface{}) error {
	if c.ApiKey != "" && c.ApiSecret != "" {
		// sign this request
		c.signRequest(req)
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	decoded, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	// Coinbase returns all it's responses in the form of
	// {"data": <some json format>}
	// We can parse it into an intermediary map
	// containing "data" -> raw bytes
	// and then grab the mapped bytes & parse it into the
	// Unmarshaler to decode our interface
	var intermediary map[string]json.RawMessage
	if err = json.Unmarshal(decoded, &intermediary); err != nil {
		return err
	}

	if _, ok := intermediary["data"]; ok {
		if err = json.Unmarshal(intermediary["data"], decode); err != nil {
			return err
		}
	} else {
		return json.Unmarshal(decoded, decode)
	}

	return nil
}

//New creates a new client using an API key and API secret
func New(apiKey, apiSecret string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
	}
}
