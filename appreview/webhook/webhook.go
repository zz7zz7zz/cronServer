package webhook

import "open.com/cronServer/appreview/models"

type Webhook interface {
	OnWebHook(appReviewRecord *models.AppReviewRecord)
}
