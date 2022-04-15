package handlers

import (
	"bytes"
	"encoding/json"
	// "encoding/json"
	// "encoding/json"
	"fmt"
	// dto "github.com/DmitriyKalekin/stalker22/telegram_api_client/json_dto"
	// log "github.com/sirupsen/logrus"
	// "github.com/go-chi/chi/v5"
	assert "github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	TELEGRAM_TOKEN = "MOCK_CORRECT_TOKEN"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func TestTgWebhookHandler(t *testing.T) {

	t.Run("200 OK on correct telegram JSON-request", func(t *testing.T) {
		// --- given --------
		json_response := `{"update_id":464725880, "message":{"message_id":11,"from":{"id":435627225,"is_bot":false,"first_name":"Дмитрий","last_name":"Калекин","username":"herr_horror","language_code":"en"},"chat":{"id":435627225,"first_name":"Дмитрий","last_name":"Калекин","username":"herr_horror","type":"private"},"date":1650009613,"text":"привет"}}`
		r := ioutil.NopCloser(bytes.NewReader([]byte(json_response)))

		// Create a request to pass to our handler
		req, err := http.NewRequest("POST", "/tg"+TELEGRAM_TOKEN, r)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(TgHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// log.Warn(rr.HeaderMap.Values("Content-Type"))

		// Check the response body is what we expect.
		expected := Response{Status: "OK"} //`{"status":"OK"}`

		var got_value Response
		json.NewDecoder(rr.Body).Decode(&got_value)

		assert.EqualValues(t, typeof(Response{}), typeof(got_value))
		assert.EqualValues(t, "OK", got_value.Status)
		assert.EqualValues(t, expected, got_value)

		// if rr.Body.String() != expected {
		// 	log.Warn([]byte(expected))
		// 	log.Warn([]byte(rr.Body.Bytes()))
		// 	t.Errorf("handler returned unexpected body: got '%v' want '%v'",
		// 		rr.Body.String(), expected)
		// }
	})

}
