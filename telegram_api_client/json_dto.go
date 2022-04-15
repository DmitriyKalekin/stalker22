/*********************************
*
* Some types taken from [https://github.com/go-telegram-bot-api/telegram-bot-api/]
**********************************/
package telegram_api_client

import "encoding/json"

// "encoding/json"
// "errors"
// "fmt"
// "net/url"
// "strings"
// "time"

// APIResponse is a response from the Telegram API with the result
// stored raw.
type APIResponse struct {
	Ok          bool                `json:"ok"`
	Result      *json.RawMessage    `json:"result,omitempty"` // can be bool, WebhookInfo or result with message
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// WebhookInfo is information about a currently set webhook.
type WebhookInfo struct {
	// URL webhook URL, may be empty if webhook is not set up.
	Url string `json:"url"`
	// HasCustomCertificate true, if a custom certificate was provided for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// PendingUpdateCount number of updates awaiting delivery.
	PendingUpdateCount int `json:"pending_update_count"`
	// IPAddress is the currently used webhook IP address
	//
	// optional
	IPAddress string `json:"ip_address,omitempty"`
	// LastErrorDate unix time for the most recent error
	// that happened when trying to deliver an update via webhook.
	//
	// optional
	LastErrorDate int `json:"last_error_date,omitempty"`
	// LastErrorMessage error message in human-readable format for the most recent error
	// that happened when trying to deliver an update via webhook.
	//
	// optional
	LastErrorMessage string `json:"last_error_message,omitempty"`
	// MaxConnections maximum allowed number of simultaneous
	// HTTPS connections to the webhook for update delivery.
	//
	// optional
	MaxConnections int `json:"max_connections,omitempty"`
	// AllowedUpdates is a list of update types the bot is subscribed to.
	// Defaults to all update types
	//
	// optional
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// IsSet returns true if a webhook is currently set.
func (info WebhookInfo) IsSet() bool {
	return info.Url != ""
}

// ResponseParameters are various errors that can be returned in APIResponse.
type ResponseParameters struct {
	// The group has been migrated to a supergroup with the specified identifier.
	//
	// optional
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	// In case of exceeding flood control, the number of seconds left to wait
	// before the request can be repeated.
	//
	// optional
	RetryAfter int `json:"retry_after,omitempty"`
}
