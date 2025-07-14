package server

import "github.com/gin-gonic/gin"

func RunHTTPServer(addr string, wrapper func(router *gin.Engine)) {
	if addr == "" {
		panic("empty server addr")
	}
	RunHTTPServerOnAddr(addr, wrapper)
}

func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.Default()
	middleware := []gin.HandlerFunc{}
	SetMiddleware(apiRouter, middleware...)
	wrapper(apiRouter)
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}

func SetMiddleware(r *gin.Engine, middleware ...gin.HandlerFunc) {
	if len(middleware) == 0 {
		return
	}
	for _, m := range middleware {
		r.Use(m)
	}
}
