package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//https://ichi.pro/pt/sua-primeira-api-rest-em-golang-com-mux-202836347743488
//https://blog.logrocket.com/making-http-requests-in-go/

func main() {
	router := mux.NewRouter()   //instância de mux
	const port string = ":8000" //port do server

	//registando rota mapeando o path URL para handler
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Server is Up and Running...")
	}) //ResponseWriter é uma interface e Request é uma struct

	router.HandleFunc("/1", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, aux)
	}).Methods("GET") //retorna URL original

	router.HandleFunc("/send/{name}", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		fmt.Fprint(res, vars, vars["name"])
		//res.WriteHeader(http.StatusOK)
		teste := vars["name"]
		Testando(teste)
	}).Methods("POST") //retorna URL encurtada

	log.Fatal(http.ListenAndServe(port, router))
	//ListenAndServe inicia um server HTTP com um endereço e um handler
}
