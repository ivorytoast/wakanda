package vision

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	"wakanda/common"
)

const (
	rateLimitRetryCount = 3
	rateLimitRetryDelay = time.Second
)

var (
	DefaultClient = NewClient(common.Credentials())
	paperBase = common.Configuration.Setup.PaperBaseUrl
	apiVersion = common.Configuration.Setup.PaperApiVersion
	clientTimeout = common.Configuration.Setup.PaperClientRequestTimeout * time.Second
	do = defaultDo
)

type Client struct {
	credentials *common.APIKey
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewClient(credentials *common.APIKey) *Client {
	return &Client{credentials: credentials}
}

func GetLastQuote(symbol string) (*LastQuoteResponse, error) {
	return DefaultClient.GetLastQuote(symbol)
}

func defaultDo(c *Client, req *http.Request) (*http.Response, error) {
	if c.credentials.OAuth != "" {
		req.Header.Set("Authorization", "Bearer " + c.credentials.OAuth)
	} else {
		req.Header.Set("APCA-API-KEY-ID", c.credentials.ID)
		req.Header.Set("APCA-API-SECRET-KEY", c.credentials.Secret)
	}

	client := &http.Client{
		Timeout: clientTimeout,
	}
	var resp *http.Response
	var err error
	for i := 0; ; i++ {
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusTooManyRequests {
			break
		}
		if i >= rateLimitRetryCount {
			break
		}
		time.Sleep(rateLimitRetryDelay)
	}

	if err = verify(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (client *Client) get(u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return do(client, req)
}

func (client *Client) GetLastQuote(symbol string) (*LastQuoteResponse, error) {

	log.Printf("Getting Last Quote")

	endpoint, err := url.Parse(fmt.Sprintf("%s/%s/assets/%s", paperBase, apiVersion, symbol))
	if err != nil {
		fmt.Println("Error parsing the url...")
		return nil, err
	}

	query := endpoint.Query()
	query.Set("symbol", symbol)

	endpoint.RawQuery = query.Encode()

	log.Printf(endpoint.Host)
	log.Printf(endpoint.Path)

	resp, err := client.get(endpoint)
	if err != nil {
		fmt.Println("Error during the call to the endpoint...")
		return nil, err
	}

	lastQuote := &LastQuoteResponse{}

	marshallingError := unmarshall(resp, &lastQuote)
	if marshallingError != nil {
		fmt.Println("Error during the marshalling process...")
		return nil, err
	}

	return lastQuote, nil
}

func verify(resp *http.Response) (err error) {
	if resp.StatusCode >= http.StatusMultipleChoices {
		var body []byte
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		apiErr := APIError{}

		err = json.Unmarshal(body, &apiErr)
		if err != nil {
			return fmt.Errorf("json unmarshal error: %s", err.Error())
		}
		err = &apiErr
	}

	return
}

func unmarshall(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}