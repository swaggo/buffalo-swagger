package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/swaggo/buffalo-swagger/example/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
