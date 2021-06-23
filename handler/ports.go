package handler

import (
	"amexProject/app/interfaces"
	"amexProject/repository"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type port struct {
	productSrv interfaces.ProductSrv
}

func NewProductPort(prod interfaces.ProductSrv) *port {
	return &port{
		productSrv: prod,
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (p *port) createProduct(w http.ResponseWriter, r *http.Request){
	var payload repository.CreateProductPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if prod, err := p.productSrv.CreateProduct(context.Background(), &payload); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}else {
		respondWithJSON(w, http.StatusCreated, &prod)
	}
}

func (p *port) updateProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	var payload repository.UpdateProductPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if prod, err := p.productSrv.UpdateProduct(context.Background(), int64(productID), &payload); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}else {
		respondWithJSON(w, http.StatusOK, &prod)
	}
}

func (p *port) getProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	if prod, err := p.productSrv.GetProduct(context.Background(), int64(productID)); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}else {
		respondWithJSON(w, http.StatusOK, &prod)
	}
}

func (p *port) getAllProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	vendor := vars["vendor"]
	count, _ := strconv.Atoi(r.FormValue("count"))
	if count > 10 || count < 1 {
		count = 10
	}
	if prod, err := p.productSrv.GetAllProducts(context.Background(), vendor ,count); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}else {
		respondWithJSON(w, http.StatusOK, &prod)
	}
}

func (p *port) deleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	if err := p.productSrv.DeleteProduct(context.Background(), int64(productID)); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "selected product deleted"})
}
