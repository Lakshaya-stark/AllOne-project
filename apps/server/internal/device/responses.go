package device

type RegisterResponse struct {
    Message  string `json:"message"`
    DeviceID string `json:"device_id"`
}