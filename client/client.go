package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/alex117de/google-places-api-go/logger"
	"github.com/alex117de/google-places-api-go/request"
	"github.com/alex117de/google-places-api-go/response"
	"io"
	"log"
	"net/http"
)

type Client struct {
	config Config
	logger logger.Logger
}

const (
	HeaderApiKey      = "X-Goog-Api-Key"
	HeaderContentType = "Content-Type"
	HeaderFieldMask   = "X-Goog-FieldMask"
)

func (c Client) GetNearByPlaces(req request.Request) (*response.Response, error) {
	res := &response.Response{}

	httpRes, err := c.doReq(req)

	if httpRes != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				c.logger.Warning("closing body failed: %v", err)
			}
		}(httpRes.Body)
	}

	if err != nil || httpRes == nil {
		return res, err
	}

	if httpRes.StatusCode < 200 || httpRes.StatusCode >= 300 {
		return res, c.handleErrorResponse(*httpRes)
	}

	err = c.readJson(*httpRes, res)

	if err != nil {
		return nil, err
	}

	c.logger.Debug("Got Response from GetNearByPlaces", res)

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
		c.logger.Error("Error encoding JSON: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", c.config.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		c.logger.Error("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set(HeaderContentType, "application/json")
	req.Header.Set(HeaderApiKey, c.config.ApiToken)
	req.Header.Set(HeaderFieldMask, request.FieldMask)

	client := &http.Client{}

	c.logger.Debug("Sending request: %v", req)
	res, err := client.Do(req)

	if err != nil {
		c.logger.Error("Error making request: %v", err)
		return nil, err
	}

	return res, nil
}

func (c Client) handleErrorResponse(httpRes http.Response) error {
	body, err := io.ReadAll(httpRes.Body)

	if err != nil {
		c.logger.Error("Error reading error response body: %v", err)
	}

	c.logger.Warning("API responded with Status-Code: %v \n Response: %v", httpRes.StatusCode, string(body))

	return errors.New("API responded error")
}

func New(config Config, logService *log.Logger) Client {
	clientLogger := logger.New(logService, config.LogLevel)

	return Client{config: config, logger: clientLogger}
}
