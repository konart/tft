package middleware_test

import (
	"testing"
	"github.com/gobuffalo/buffalo"
	"github.com/konart/tft/middleware"
)


func TestGetMetaInfo(t *testing.T) {
	_ = middleware.GetMetaInfo(func(c buffalo.Context) error {
		return nil
	})
}
