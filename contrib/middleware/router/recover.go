package router

import (
	"github.com/sllt/booby"
	"github.com/sllt/booby/util"
)

// Recover returns the recovery middleware handler.
func Recover() booby.HandlerFunc {
	return func(ctx *booby.Context) {
		defer util.Recover()
		ctx.Next()
	}
}
