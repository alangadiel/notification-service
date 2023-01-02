package service

import (
	"errors"
	"time"

	"github.com/alangadiel/notification-service/internal/model"
)

const (
	NotificationTypeStatus = iota
	NotificationTypeNews
	NotificationTypeMarketing
)

func GetRule(notifType int) (model.Rule, error) {
	switch notifType {
	case NotificationTypeStatus:
		return model.Rule{
			Limit:    2,
			Duration: time.Minute,
		}, nil

	case NotificationTypeNews:
		return model.Rule{
			Limit:    1,
			Duration: time.Hour * 24,
		}, nil

	case NotificationTypeMarketing:
		return model.Rule{
			Limit:    3,
			Duration: time.Hour,
		}, nil
	default:
		return model.Rule{}, errors.New("invalid type")
	}
}
