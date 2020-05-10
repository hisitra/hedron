package restserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/grpcserver"
	iot "github.com/hisitra/hedron/src/iotranslator"
	"io/ioutil"
	"log"
	"net/http"
)

func New() Server {
	return &server{Port: configs.Server.RestPort}
}

func (s *server) Start() {
	log.Println("Info: Starting Hedron Node:", configs.Node.Name, "REST Server at PORT:", configs.Server.RestPort)
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
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse("Failed to decode request"))
		return
	}

	res, err := grpcserver.New().Create(
		context.Background(),
		&comcn.Message{Value: reqJSON})

	if err != nil {
		_, err = w.Write(iot.InternalServerResponse("").Marshal())
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}

	_, err = w.Write(res.Value)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (s *server) Read(w http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse("Failed to decode request"))
		return
	}

	res, err := grpcserver.New().Read(
		context.Background(),
		&comcn.Message{Value: reqJSON})

	if err != nil {
		_, err = w.Write(iot.InternalServerResponse("").Marshal())
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}

	_, err = w.Write(res.Value)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (s *server) Update(w http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse("Failed to decode request"))
		return
	}

	res, err := grpcserver.New().Update(
		context.Background(),
		&comcn.Message{Value: reqJSON})

	if err != nil {
		_, err = w.Write(iot.InternalServerResponse("").Marshal())
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}

	_, err = w.Write(res.Value)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (s *server) Delete(w http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintln(w, iot.BadRequestResponse("Failed to decode request"))
		return
	}

	res, err := grpcserver.New().Delete(
		context.Background(),
		&comcn.Message{Value: reqJSON})

	if err != nil {
		_, err = w.Write(iot.InternalServerResponse("").Marshal())
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}

	_, err = w.Write(res.Value)
	if err != nil {
		w.WriteHeader(500)
	}
}
