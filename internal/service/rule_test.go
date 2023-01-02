package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRule(t *testing.T) {
	type testCase struct {
		notifType int
		err       bool
	}
	tcs := []testCase{
		{
			notifType: -1,
			err:       true,
		},
		{
			notifType: NotificationTypeStatus,
		},
		{
			notifType: NotificationTypeNews,
		},
		{
			notifType: NotificationTypeMarketing,
		},
	}

	for _, tc := range tcs {
		res, err := GetRule(tc.notifType)
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, res.Limit)
			assert.NotEmpty(t, res.Duration)
		}
	}
}
