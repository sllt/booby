package router

import (
	"time"

	"github.com/sllt/booby"
	"github.com/sllt/booby/log"
)

// Logger returns the logger middleware.
func Logger() booby.HandlerFunc {
	return func(ctx *booby.Context) {
		t := time.Now()

		ctx.Next()

		cmd := ctx.Message.Cmd()
		method := ctx.Message.Method()
		addr := ctx.Client.Conn.RemoteAddr()
		cost := time.Since(t).Milliseconds()

		switch cmd {
		case booby.CmdRequest, booby.CmdNotify:
			log.Info("'%v',\t%v,\t%v ms cost", method, addr, cost)
			break
		default:
			log.Error("invalid cmd: %d,\tdropped", cmd)
			ctx.Done()
			break
		}
	}
}
