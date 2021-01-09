package contracts

// ShieldResponse is the payload that Shields.io expects to receive for custom shields endpoints:
// https://shields.io/endpoint
type ShieldResponse struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}
