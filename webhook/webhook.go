package webhook

import "cronServer/models"

type Webhook interface {
	OnWebHook(appReviewRecord *models.AppReviewRecord)
}
