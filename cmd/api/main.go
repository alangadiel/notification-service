package main

import (
	"os"

	"github.com/alangadiel/notification-service/internal/client"
	"github.com/alangadiel/notification-service/internal/model"
	"github.com/alangadiel/notification-service/internal/service"
)

func main() {
	notifService := service.Notification{
		Gateway: &client.Gateway{OutputWriter: os.Stdout},
		NotificationHistory: model.NotificationHistory{
			Get: map[model.NotificationHistoryKey]model.NotificationHistoryValue{},
		},
	}

	printIfError(notifService.Send(service.NotificationTypeMarketing, model.UserID(1), "test1"))
	printIfError(notifService.Send(service.NotificationTypeMarketing, model.UserID(1), "test2"))
	printIfError(notifService.Send(service.NotificationTypeMarketing, model.UserID(1), "test3"))
	printIfError(notifService.Send(service.NotificationTypeMarketing, model.UserID(1), "test4"))
	printIfError(notifService.Send(service.NotificationTypeMarketing, model.UserID(2), "test5"))
}

func printIfError(err error) {
	if err != nil {
		println(err.Error())
	}
}
