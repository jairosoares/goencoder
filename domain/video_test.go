package domain_test

import (
	"encoder/domain"
	"testing"
	"github.com/stretchr/testify/require"
	"time"
	uuid "github.com/satori/go.uuid"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotUUID(t *testing.T) {
	v := domain.NewVideo()
	v.ID = "abc"
	v.ResourceID = "a"
	v.FilePath = "path"
	v.CreatedAt = time.Now()
	err := v.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	v := domain.NewVideo()
	v.ID = uuid.NewV4().String()
	v.ResourceID = "a"
	v.FilePath = "path"
	v.CreatedAt = time.Now()
	err := v.Validate()
	require.Nil(t, err)
}