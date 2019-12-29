package server

import (
	"autobench-server/apis"
	
	"github.com/go-martini/martini"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	mar    *martini.martini
	router *httprouter.Router
}

func New() (*Server, error) {
	sv := new(Server)
	sv.router = httprouter.New()
	sv.mar =  martini.Classic()

	for i := 0; i < len(apis.APIs); i++ {
		apis.AddAPI(sv.router, apis.APIs[i])
	}

	sv.mar.UseHandler(sv.router)
	return sv, nil
}

func (s *Server) Run(port string) {
	s.mar.Run(port)
}
