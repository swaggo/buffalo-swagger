package grifts

import (
	"github.com/cippaciong/mw-swaggo/example/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
