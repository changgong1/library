package util

import (
	"net/http"
	"net/http/pprof"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pprofSrv *http.Server

func InitPprof(prot int) {
	pprofServer(0, prot)
}

// swt-0-关闭 1-开启
func OpenOrClosePprof(swt, prot int) {
	pprofServer(swt, prot)
}

func ClosePprof() {
	pprofServer(0, 0)
}

// start-0 close pprofSrv
// start-1 open pprofSrv
// prot 开放端口
func pprofServer(start, prot int) {
	if start == 0 {
		if pprofSrv != nil && pprofSrv.Addr != "" {
			if err := pprofSrv.Close(); err != nil {
			}
			pprofSrv = nil
		}
		return
	} else {
		if pprofSrv != nil {
			return
		}
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//address := global.MyConfig.Read("pprof", "address")
	address := ":" + strconv.Itoa(prot)
	routeRegisterPprof(r)
	pprofSrv = &http.Server{
		Addr:    address,
		Handler: r,
	}
	SafeGo(func() {
		if err := pprofSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	})
}

func routeRegisterPprof(rg *gin.Engine) {
	prefix := "/debug/pprof"

	prefixRouter := rg.Group(prefix)
	{
		prefixRouter.GET("/", pprofHandler(pprof.Index))
		prefixRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
		prefixRouter.GET("/profile", pprofHandler(pprof.Profile))
		prefixRouter.POST("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/trace", pprofHandler(pprof.Trace))
		prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		prefixRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
