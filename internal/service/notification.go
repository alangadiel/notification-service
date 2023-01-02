package service

import (
	"fmt"
	"time"

	"github.com/alangadiel/notification-service/internal/model"
)

type gateway interface {
	Send(userID model.UserID, message string) error
}

type Notification struct {
	Gateway             gateway
	NotificationHistory model.NotificationHistory
}

func (n *Notification) Send(notifType int, userID model.UserID, message string) error {

	var rule model.Rule
	{
		var err error
		if rule, err = GetRule(notifType); err != nil {
			return fmt.Errorf("error getting rule: %w", err)
		}
	}

	nhk := model.NotificationHistoryKey{
		UserID:           userID,
		NotificationType: notifType,
	}

	n.NotificationHistory.Mutex.Lock()
	locked := true

	defer func() {
		if locked {
			n.NotificationHistory.Mutex.Unlock()
		}
	}()

	history, notFirstTime := n.NotificationHistory.Get[nhk]
	now := time.Now()
	if notFirstTime {
		if len(history) == rule.Limit {
			if !history[0].Add(rule.Duration).Before(now) {
				return fmt.Errorf("user %d exceeded limit for this notification type", userID)
			} else {
				n.NotificationHistory.Get[nhk] = append(n.NotificationHistory.Get[nhk][1:], now)
			}
		} else {
			n.NotificationHistory.Get[nhk] = append(n.NotificationHistory.Get[nhk], now)
		}
	} else {
		n.NotificationHistory.Get[nhk] = []time.Time{now}
	}

	locked = false
	n.NotificationHistory.Mutex.Unlock()

	return n.Gateway.Send(userID, message)
}
