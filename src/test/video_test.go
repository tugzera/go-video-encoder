package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)


func TestIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()

	err := video.Validate()

	require.Error(t, err)
}


func TestIfVideoIDIsUUID(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "oab"

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = uuid.NewV4().String()
	video.FilePath = "video"
	video.CreatedDate = time.Now()

	err := video.Validate()

	require.Nil(t, err)

}