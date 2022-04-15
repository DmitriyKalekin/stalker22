package telegram_api_client

type WebhookInfoResult struct {
	Url                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   uint32   `json:"pending_update_count"`
	LastErrorDate        uint64   `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	MaxConnections       uint32   `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

type WebhookInfo struct {
	Ok          bool              `json:"ok"`
	ErrorCode   uint32            `json:"error_code,omitempty"`
	Description string            `json:"description,omitempty"`
	Result      WebhookInfoResult `json:"result"`
}

type WebhookSetInfo struct {
	Ok          bool   `json:"ok"`
	ErrorCode   uint32 `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
	Result      bool   `json:"result"`
}
