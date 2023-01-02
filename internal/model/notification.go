package model

import (
	"sync"
	"time"
)

type Rule struct {
	Limit    int
	Duration time.Duration
}

type UserID int64

type NotificationHistoryKey struct {
	UserID
	NotificationType int
}

type NotificationHistoryValue []time.Time

type NotificationHistory struct {
	Mutex sync.Mutex
	Get   map[NotificationHistoryKey]NotificationHistoryValue
}
