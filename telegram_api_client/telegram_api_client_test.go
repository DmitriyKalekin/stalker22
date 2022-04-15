package telegram_api_client

import (
	"bytes"
	"encoding/json"
	// "encoding/json"
	"fmt"
	// dto "github.com/DmitriyKalekin/stalker22/telegram_api_client/json_dto"
	// log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc  func(req *http.Request) (*http.Response, error)
	GetFunc func(url string) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, nil
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	if m.GetFunc != nil {
		return m.GetFunc(url)
	}
	return &http.Response{}, nil
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func TestWrongTokenOnGetWebhookInfo(t *testing.T) {
	// --- given --------
	json_response := `{"description":"Not Found","error_code":404,"ok":false}`

	telegram_client := NewClient("WRONG_TOKEN")
	telegram_client.HttpClient = &MockClient{
		GetFunc: func(url string) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(json_response)))
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       r,
			}, nil
		},
	}

	// --- try --------
	info, err := telegram_client.GetWebhookInfo()

	// --- assert --------
	assert.NotNil(t, info)
	assert.Nil(t, err)
	assert.EqualValues(t, typeof(APIResponse{}), typeof(info))
	assert.EqualValues(t, false, info.Ok)
	assert.EqualValues(t, http.StatusNotFound, info.ErrorCode)

}

func TestGetWebhookInfo(t *testing.T) {
	// --- given --------
	json_response := `{"ok":true,"result":{"url":"","has_custom_certificate":false,"pending_update_count":0}}`

	telegram_client := NewClient("CORRECT_TELEGRAM_TOKEN")
	telegram_client.HttpClient = &MockClient{
		GetFunc: func(url string) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(json_response)))
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       r,
			}, nil
		},
	}

	// --- try --------
	info, err := telegram_client.GetWebhookInfo()
	var result WebhookInfo
	json.Unmarshal(*info.Result, &result)

	// --- assert --------
	assert.NotNil(t, info)
	assert.Nil(t, err)
	assert.EqualValues(t, typeof(APIResponse{}), typeof(info))
	assert.EqualValues(t, true, info.Ok)

	assert.EqualValues(t, "", result.Url)
	assert.EqualValues(t, false, result.HasCustomCertificate)
	assert.EqualValues(t, []string(nil), result.AllowedUpdates)
}

func TestWrongTokenOnSetWebhook(t *testing.T) {
	// --- given --------
	json_response := `{"description":"Not Found","error_code":404,"ok":false}`

	telegram_client := NewClient("WRONG_TOKEN")
	telegram_client.HttpClient = &MockClient{
		GetFunc: func(url string) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(json_response)))
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       r,
			}, nil
		},
	}

	// --- try --------
	url := "http://localhost:3000/tg"
	info, err := telegram_client.SetWebhook(url)

	// --- assert --------
	assert.NotNil(t, info)
	assert.Nil(t, err)
	assert.EqualValues(t, typeof(APIResponse{}), typeof(info))
	assert.EqualValues(t, false, info.Ok)
	assert.EqualValues(t, http.StatusNotFound, info.ErrorCode)
}

func TestNoHTTPSOnSetWebhook(t *testing.T) {
	// --- given --------
	json_response := `{"description":"Bad Request: bad webhook: HTTPS url must be provided for webhook", "error_code":400, "ok":false}`

	telegram_client := NewClient("CORRECT_TOKEN")
	telegram_client.HttpClient = &MockClient{
		GetFunc: func(url string) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(json_response)))
			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       r,
			}, nil
		},
	}

	// --- try --------
	url := "http://localhost:3000/tg"
	info, err := telegram_client.SetWebhook(url)

	// --- assert --------
	assert.NotNil(t, info)
	assert.Nil(t, err)
	assert.EqualValues(t, typeof(APIResponse{}), typeof(info))
	assert.EqualValues(t, false, info.Ok)
	assert.EqualValues(t, http.StatusBadRequest, info.ErrorCode)
}

func TestWasSetWebhook(t *testing.T) {
	// --- given --------
	json_response := `{"description":"Webhook was set", "ok":true, "result":true}`

	telegram_client := NewClient("CORRECT_TOKEN")
	telegram_client.HttpClient = &MockClient{
		GetFunc: func(url string) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(json_response)))
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       r,
			}, nil
		},
	}

	// --- try --------
	url := "https://15ba-217-25-217-143.eu.ngrok.io/tg"
	info, err := telegram_client.SetWebhook(url)
	var result bool
	json.Unmarshal(*info.Result, &result)

	// --- assert --------
	assert.NotNil(t, info)
	assert.Nil(t, err)
	assert.EqualValues(t, typeof(APIResponse{}), typeof(info))
	assert.EqualValues(t, true, info.Ok)
	assert.EqualValues(t, "Webhook was set", info.Description)
	assert.EqualValues(t, 0, info.ErrorCode)

	assert.EqualValues(t, true, result)
}

func TestAlreadySetWebhook(t *testing.T) {
	// --- given --------
	json_response := `{"ok":true,"description":"Webhook is already set","result":true}`

	telegram_client := NewClient("CORRECT_TOKEN")
	telegram_client.HttpClient = &MockClient{
		GetFunc: func(url string) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(json_response)))
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       r,
			}, nil
		},
	}

	// --- try --------
	url := "https://15ba-217-25-217-143.eu.ngrok.io/tg"
	info, err := telegram_client.SetWebhook(url)
	var result bool
	json.Unmarshal(*info.Result, &result)

	// --- assert --------
	assert.NotNil(t, info)
	assert.Nil(t, err)
	assert.EqualValues(t, typeof(APIResponse{}), typeof(info))
	assert.EqualValues(t, true, info.Ok)
	assert.EqualValues(t, "Webhook is already set", info.Description)
	assert.EqualValues(t, 0, info.ErrorCode)

	assert.EqualValues(t, true, result)
}
