package server

import (
	"autobench-server/apis"
	
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	neg    *negroni.Negroni
	router *httprouter.Router
}

func New() (*Server, error) {
	sv := new(Server)
	sv.router = httprouter.New()
	sv.neg =  negroni.Classic()

	for i := 0; i < len(apis.APIs); i++ {
		apis.AddAPI(sv.router, apis.APIs[i])
	}

	sv.neg.UseHandler(sv.router)
	return sv, nil
}

func (s *Server) Run(port string) {
	s.neg.Run(port)
}
