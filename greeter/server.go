package greeter

import (
	"encoding/json"
	"fmt"
	"github.com/apetesh/gothings/greeter/api"

	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {}

func (srv *Server) Start(port string) error {
	router := http.NewServeMux()
	router.HandleFunc("/greet", srv.greet)
	log.Printf("get your greetings on port %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func (srv *Server) greet(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer func() {
		if err = r.Body.Close(); err != nil {
			log.Print("error closing response body")
		}
	}()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	body := &api.GreetRequest{}
	err = json.Unmarshal(b, body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseBody, err := json.Marshal(&api.GreetResponse{Message: fmt.Sprintf("Well, hello there %s", body.Name)})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ , err = w.Write(responseBody)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}