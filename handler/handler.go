package handler

import (
	"amexProject/app"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type handler struct{
	Router *mux.Router
	App    *app.App
}

func New(app *app.App) *handler {
	return &handler{
		Router: mux.NewRouter(),
		App: app,
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func (h *handler) InitialiseRoutes(){
	productPorts := NewProductPort(h.App.ProductSrv)
	h.Router.HandleFunc("/", helloWorldHandler)
	h.Router.HandleFunc("/api/product", productPorts.createProduct).Methods("POST")
	h.Router.HandleFunc("/api/product/{vendor}", productPorts.getAllProduct).Methods("GET")
	h.Router.HandleFunc("/api/product/{id:[0-9]+}", productPorts.updateProduct).Methods("PUT")
	h.Router.HandleFunc("/api/product/{id:[0-9]+}", productPorts.getProduct).Methods("GET")
	h.Router.HandleFunc("/api/product/{id:[0-9]+}", productPorts.deleteProduct).Methods("DELETE")
}

func (h *handler) Run(){
	srv := &http.Server{
		Handler: h.Router,
		Addr: ":8000",
	}
	log.Printf("connect to http://localhost:8000")
	log.Fatal(srv.ListenAndServe())
}

