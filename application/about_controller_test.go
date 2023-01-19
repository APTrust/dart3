package application_test

import (
	"dart/application"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAboutShow(t *testing.T) {
	app := application.GetAppInstance()
	resp := app.AboutShow()
	assert.NotNil(t, resp)
}
