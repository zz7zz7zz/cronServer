package webhook

type Webhook interface {
	sendTextMessage()
}
