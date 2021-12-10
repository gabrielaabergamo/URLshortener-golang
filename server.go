package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//https://ichi.pro/pt/sua-primeira-api-rest-em-golang-com-mux-202836347743488
//https://blog.logrocket.com/making-http-requests-in-go/

func MetodoGet(router *mux.Router) {
	router.HandleFunc("/retrieve/{nameog}", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		fmt.Fprint(res, URLGet(vars["nameog"]))
	}).Methods("GET") //retorna URL original
}

func MetodoPost(router *mux.Router) {
	router.HandleFunc("/send/{name}", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		//fmt.Fprint(res, vars, vars["name"])
		//res.WriteHeader(http.StatusOK)
		longURL := vars["name"]
		fmt.Fprint(res, URLPost(longURL))
		//URLPost(longURL)
	}).Methods("POST") //retorna URL encurtada
}

func Servidor() {
	router := mux.NewRouter()   //instância de mux
	const port string = ":8000" //port do server

	MetodoGet(router)

	MetodoPost(router)

	log.Fatal(http.ListenAndServe(port, router))
	//ListenAndServe inicia um server HTTP com um endereço e um handler
}
