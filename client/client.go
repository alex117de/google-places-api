package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"places-client/request"
	"places-client/response"
)

type Client struct {
	config Config
}

const (
	HeaderApiKey      = "X-Goog-Api-Key"
	HeaderContentType = "Content-Type"
	HeaderFieldMask   = "X-Goog-FieldMask"
)

func (c Client) GetNearByPlaces(req request.Request) (*response.Response, error) {
	res := &response.Response{}

	httpRes, err := c.doReq(req)

	if err != nil {
		return res, err
	}

	err = c.readJson(*httpRes, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c Client) readJson(httpResponse http.Response, response *response.Response) error {
	body, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, response)
}

func (c Client) doReq(request request.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(request.Body)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", c.config.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set(HeaderContentType, "application/json")
	req.Header.Set(HeaderApiKey, c.config.ApiToken)
	req.Header.Set(HeaderFieldMask, request.FieldMask)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

func New() (*Client, error) {
	conf, err := NewConfig()

	if err != nil {
		return nil, err
	}

	return &Client{config: *conf}, nil
}
