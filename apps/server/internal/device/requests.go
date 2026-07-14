package device

type RegisterRequest struct {
	Name       string `json:"name"`
	Platform   string `json:"platform"`
	DeviceType string `json:"device_type"`
	PublicKey  string `json:"public_key"`
}