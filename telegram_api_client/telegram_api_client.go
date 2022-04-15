package telegram_api_client

import (
	"encoding/json"
	//  "github.com/DmitriyKalekin/stalker22/pkg/json_dto"
	// "io/ioutil"
	// log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	BASE_URL = "https://api.telegram.org/bot"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

type TelegramApiClient struct {
	Token      string
	Url        string
	HttpClient HTTPClient
}

func NewClient(token string) *TelegramApiClient {
	c := new(TelegramApiClient)
	c.Token = token
	c.Url = BASE_URL + token + "/"
	c.HttpClient = Client
	return c
}

// Use this method to get current webhook status.
// Requires no parameters.
// On success, returns a WebhookInfo object.
// If the bot is using getUpdates, will return an object with the url field empty.
func (client *TelegramApiClient) GetWebhookInfo() (APIResponse, error) {
	var res APIResponse
	resp, err := client.HttpClient.Get(client.Url + "getWebhookInfo")

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

// Use this method to specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts.
// Returns True on success.
// If you'd like to make sure that the Webhook request comes from Telegram,
// we recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot's token, you can be pretty sure it's us.
func (client *TelegramApiClient) SetWebhook(wh_url string) (APIResponse, error) {
	var res APIResponse
	url := client.Url + "setWebhook?url=" + wh_url +
		"&allowed_updates=[\"message\", \"edited_message\", \"channel_post\", \"edited_channel_post\", \"inline_query\", \"chosen_inline_result\", \"callback_query\", \"my_chat_member\", \"chat_member\", \"poll\", \"poll_answer\"]"

	resp, err := client.HttpClient.Get(url)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

// async def setWebhook(self, wh_url: str) -> bool:
// """
// Use this method to specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts.
// Returns True on success.
// If you'd like to make sure that the Webhook request comes from Telegram,
// we recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot's token, you can be pretty sure it's us.
// :param wh_url:
// :raises TgException
// :return:
// """
// response = await self._tg.setWebhook(wh_url)
// a = self._try_parse_result(response)
// return response.payload["result"]

// async def deleteWebhook(self) -> bool:
// """
// Use this method to remove webhook integration if you decide to switch back to getUpdates.
// Returns True on success.
// :raises TgException
// :return:
// """
// response = await self._tg.deleteWebhook()
// self._try_parse_result(response)
// return response.payload["result"]
