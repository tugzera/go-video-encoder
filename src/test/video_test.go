package domain__test

import (
	"encoder/domain"
	"testing"

	"github.com/stretchr/testify/require"
)


func Test(t *testing.T) {
	video := domain.NewVideo()

	err := video.Validate()

	require.Error(t, err)
}