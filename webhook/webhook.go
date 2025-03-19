package webhook

type Webhook interface {
	onWebhook()
}
