package context

import (
	"flag"
	"os"
	"runtime/debug"

	"github.com/zulu-network/zulu-go-sdk/log"
)

// Run service
func Run(handle func(Context) error) {
	var h bool
	var c string
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&c, "c", "etc/bridge/conf.yml", "the configuration file")
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	ctx := NewContext(c)
	defer func() {
		if r := recover(); r != nil {
			ctx.Log().Error("service is stopped with panic", log.Any("panic", r), log.Any("stack", string(debug.Stack())))
		}
	}()

	pwd, _ := os.Getwd()
	ctx.Log().Info("service starting", log.Any("args", os.Args), log.Any("pwd", pwd))
	err := handle(ctx)
	if err != nil {
		ctx.Log().Error("service has stopped with error", log.Error(err))
	} else {
		ctx.Log().Info("service has stopped")
	}
}
