package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yexk/src/api/dispatch"
)

// Server 服务
type Server struct {
	baseServer *gin.Engine
	version    string
}

// Webserver 提供web服务
func Webserver(port int) *Server {
	s := Server{
		baseServer: gin.Default(),
		version:    "v1.0.0",
	}
	// init webserver
	s.init(port)

	return &s
}

func (s *Server) init(port int) {
	/* 添加路由  */
	// moreCredit Server Provider
	s.addrouter("/emts/easycreditrs/reqdispatch", dispatch.Reqdispath, "POST")

	ports := fmt.Sprintf(":%d", port)
	s.baseServer.Run(ports)
}

func (s *Server) addrouter(url string, callback func(c *gin.Context), method string) {
	switch method {
	case "GET":
		s.baseServer.GET(url, callback)
	case "POST":
		s.baseServer.POST(url, callback)
	}
}
