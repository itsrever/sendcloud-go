package sendcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-xray-sdk-go/xray"
)

type Payload interface {
	GetPayload() interface{}
}

type Response interface {
	GetResponse() interface{}
	SetResponse(body []byte) error
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Request string `json:"request"`
		Message string `json:"message"`
	} `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Request string `json:"request"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("request %s resulted in error code %d: %s", e.Request, e.Code, e.Message)
}

// Send a request to Sendcloud with given method, path, payload and credentials
func Request(ctx context.Context, method string, uri string, payload Payload, apiKey string, apiSecret string, r Response) error {
	var request *http.Request
	var err error

	if payload == nil {
		request, err = http.NewRequest(method, getUrl(uri), nil)
		if err != nil {
			return err
		}
	} else {
		body, err := json.Marshal(payload.GetPayload())
		if err != nil {
			return err
		}
		request, err = http.NewRequest(method, getUrl(uri), bytes.NewBuffer(body))
		if err != nil {
			return err
		}
	}

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "Sendcloud-Go/0.1 ("+apiKey+")")
	request.SetBasicAuth(apiKey, apiSecret)

	client := xray.Client(nil)
	request = request.WithContext(ctx)

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode > 299 || response.StatusCode < 200 {
		//Return error response
		errResponse := ErrorResponse{}
		err = json.Unmarshal(body, &errResponse)
		if err != nil {
			return err
		}
		return &Error{
			Code:    response.StatusCode,
			Request: errResponse.Error.Request,
			Message: errResponse.Error.Message,
		}
	}
	err = r.SetResponse(body)
	return err
}

// Return the full URL
func getUrl(uri string) string {
	var url string
	if strings.HasPrefix(uri, "https://") {
		url = uri
	} else if strings.HasPrefix(uri, "/") {
		url = "https://panel.sendcloud.sc" + uri
	} else {
		url = "https://panel.sendcloud.sc/" + uri
	}

	return url
}
