package service

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/alangadiel/notification-service/internal/client"
	"github.com/alangadiel/notification-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendErrorGetRule(t *testing.T) {
	notifService := Notification{}

	invalidNotifType := -1
	userID := model.UserID(1)
	message := "test"

	err := notifService.Send(invalidNotifType, userID, message)

	assert.ErrorContains(t, err, "error getting rule")
}

func TestSendErrorLimitExceeded(t *testing.T) {
	now := time.Now()
	userID := model.UserID(1)
	message := "test"
	notifType := NotificationTypeMarketing

	notifService := Notification{
		NotificationHistory: model.NotificationHistory{
			Get: map[model.NotificationHistoryKey]model.NotificationHistoryValue{
				{UserID: userID, NotificationType: notifType}: []time.Time{now, now, now},
			},
		},
	}

	err := notifService.Send(notifType, userID, message)

	assert.ErrorContains(t, err, "exceeded limit for this notification type")
}

func TestSendFirstTime(t *testing.T) {
	userID := model.UserID(1)
	message := "test"
	notifType := NotificationTypeMarketing
	var gatewayOutput bytes.Buffer

	nhk := model.NotificationHistoryKey{UserID: userID, NotificationType: notifType}

	notifService := Notification{
		NotificationHistory: model.NotificationHistory{
			Get: map[model.NotificationHistoryKey]model.NotificationHistoryValue{},
		},
		Gateway: &client.Gateway{OutputWriter: &gatewayOutput},
	}

	err := notifService.Send(notifType, userID, message)
	require.NoError(t, err)

	assert.Equal(t, 1, len(notifService.NotificationHistory.Get[nhk]))

	res := gatewayOutput.String()
	assert.Contains(t, res, message)
	assert.Contains(t, res, fmt.Sprint(userID))
}

func TestSendSuccessWithFullSlice(t *testing.T) {
	now := time.Now()
	userID := model.UserID(1)
	message := "test"
	notifType := NotificationTypeMarketing
	var gatewayOutput bytes.Buffer

	var oldTime time.Time
	{
		var err error
		oldTime, err = time.Parse("2006-01-02", "2006-01-02")
		if err != nil {
			t.Fail()
		}
	}

	nhk := model.NotificationHistoryKey{UserID: userID, NotificationType: notifType}

	notifService := Notification{
		NotificationHistory: model.NotificationHistory{
			Get: map[model.NotificationHistoryKey]model.NotificationHistoryValue{
				nhk: []time.Time{oldTime, now, now},
			},
		},
		Gateway: &client.Gateway{OutputWriter: &gatewayOutput},
	}

	err := notifService.Send(notifType, userID, message)
	require.NoError(t, err)

	assert.Equal(t, 3, len(notifService.NotificationHistory.Get[nhk]))

	res := gatewayOutput.String()
	assert.Contains(t, res, message)
	assert.Contains(t, res, fmt.Sprint(userID))
}

func TestSendSuccessWithoutFullSlice(t *testing.T) {
	now := time.Now()
	userID := model.UserID(1)
	message := "test"
	notifType := NotificationTypeMarketing
	var gatewayOutput bytes.Buffer

	nhk := model.NotificationHistoryKey{UserID: userID, NotificationType: notifType}

	notifService := Notification{
		NotificationHistory: model.NotificationHistory{
			Get: map[model.NotificationHistoryKey]model.NotificationHistoryValue{
				nhk: []time.Time{now, now},
			},
		},
		Gateway: &client.Gateway{OutputWriter: &gatewayOutput},
	}

	err := notifService.Send(notifType, userID, message)
	require.NoError(t, err)

	assert.Equal(t, 3, len(notifService.NotificationHistory.Get[nhk]))

	res := gatewayOutput.String()
	assert.Contains(t, res, message)
	assert.Contains(t, res, fmt.Sprint(userID))
}
