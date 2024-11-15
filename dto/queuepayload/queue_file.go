package queuepayload

type QueueFileM3U8Payload struct {
	Path     string `json:"path"`
	IpServer string `json:"ipServer"`
	Uuid     string `json:"uuid"`
	Quantity string `json:"quantity"`
}
