package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/konart/tft/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
