package driver

import (
	"net"

	"github.com/gin-gonic/gin"
)

type GinIface interface {
	DriverIface
	GetEngine() *gin.Engine
}

type ginServer struct {
	s   *gin.Engine
	lis net.Listener
}

func NewGinServer() GinIface {
	s := gin.New()
	return &ginServer{s: s}
}

func (gin *ginServer) GetEngine() *gin.Engine {
	return gin.s
}

func (gin *ginServer) SetListener(lis net.Listener) {
	gin.lis = lis
}

func (gin *ginServer) Serve() error {
	return gin.s.RunListener(gin.lis)
}
