package client

import (
	"fmt"
	"io"

	"github.com/alangadiel/notification-service/internal/model"
)

type Gateway struct {
	OutputWriter io.Writer
}

func (g *Gateway) Send(userID model.UserID, message string) error {
	_, err := fmt.Fprintf(g.OutputWriter, "User: %d. Message: %s\n", userID, message)
	return err
}
