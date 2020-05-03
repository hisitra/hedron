package restserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/grpcserver"
	iot "github.com/hisitra/hedron/src/iotranslator"
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
	req := &comcn.Input{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse(""))
		return
	}

	output, err := grpcserver.New().Create(context.Background(), req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.InternalServerErrorResponse(err.Error()))
		return
	}
	_, _ = fmt.Fprintln(w, output)
}

func (s *server) Read(w http.ResponseWriter, r *http.Request) {
	req := &comcn.Input{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse(""))
		return
	}

	output, err := grpcserver.New().Read(context.Background(), req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.InternalServerErrorResponse(err.Error()))
		return
	}
	_, _ = fmt.Fprintln(w, output)
}

func (s *server) Update(w http.ResponseWriter, r *http.Request) {
	req := &comcn.Input{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse(""))
		return
	}

	output, err := grpcserver.New().Update(context.Background(), req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.InternalServerErrorResponse(err.Error()))
		return
	}
	_, _ = fmt.Fprintln(w, output)
}

func (s *server) Delete(w http.ResponseWriter, r *http.Request) {
	req := &comcn.Input{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse(""))
		return
	}

	output, err := grpcserver.New().Delete(context.Background(), req)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.InternalServerErrorResponse(err.Error()))
		return
	}
	_, _ = fmt.Fprintln(w, output)
}
