package client

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/alangadiel/notification-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	userID := model.UserID(123)
	message := "test"

	var b bytes.Buffer

	g := Gateway{
		OutputWriter: &b,
	}

	err := g.Send(userID, message)

	require.NoError(t, err)

	res := b.String()

	assert.Contains(t, res, message)
	assert.Contains(t, res, fmt.Sprint(userID))
}
