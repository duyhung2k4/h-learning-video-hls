package queuepayload

type QueueMp4QuantityPayload struct {
	Path     string `json:"path"`
	Uuid     string `json:"uuid"`
	IpServer string `json:"ipServer"`
}

type QueueUrlQuantityPayload struct {
	Url      string `json:"url"`
	Quantity string `json:"quantity"`
	Uuid     string `json:"uuid"`
}
