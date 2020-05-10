package restserver

import (
	"github.com/gorilla/mux"
	"github.com/hisitra/hedron/src/configs"
	"log"
	"net/http"
)

func New() Server {
	return &server{Port: configs.Server.RestPort}
}

func (s *server) Start() {
	log.Println("Starting Hedron Node:", configs.Node.Name, "REST Server at PORT:", configs.Server.RestPort)
	_ = http.ListenAndServe(":"+s.Port, s.getHandler())
}

func (s *server) getHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/create", s.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/read", s.Read).Methods("POST", "OPTIONS")
	router.HandleFunc("/update", s.Update).Methods("POST", "OPTIONS")
	router.HandleFunc("/delete", s.Delete).Methods("POST", "OPTIONS")

	return router
}

func (s *server) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *server) Read(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *server) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *server) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
