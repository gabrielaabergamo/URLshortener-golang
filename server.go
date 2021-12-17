package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

//https://ichi.pro/pt/sua-primeira-api-rest-em-golang-com-mux-202836347743488
//https://blog.logrocket.com/making-http-requests-in-go/

func verificarURL(url string) bool {
	match, _ := regexp.MatchString("[(http(s)?):\\//\\//(www\\.)?a-zA-Z0-9\\+]{2,256}\\.[a-z]{2,6}", url)
	return match
}

func verificarCodigo(codigo string) bool {
	//match, _ := regexp.MatchString("^[a-zA-Z0-9]{6,6}$", codigo)
	//return match
	return true
}

func MetodoGet(router *mux.Router) {
	router.HandleFunc("/retrieve/", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		res.Header().Set("Content-Type", "application/json")
		verificacao := verificarCodigo(vars["nameog"])
		if verificacao {
			fmt.Fprint(res, URLGet(vars["nameog"]))
		} else {
			fmt.Fprint(res, "Código inválido")
		}

	}).Methods("GET").Queries("url", "{nameog}") //retorna URL original
}

func MetodoPost(router *mux.Router) {
	router.HandleFunc("/send/", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		res.Header().Set("Content-Type", "application/json")
		//fmt.Fprint(res, vars, vars["name"])
		//res.WriteHeader(http.StatusOK)
		longURL := vars["name"]
		verificacao := verificarURL(longURL)
		// valid := validation.Validate(longURL, validation.Required, validation.Length(5, 100), is.URL)
		// log.Println(valid, longURL)
		if verificacao {
			fmt.Fprint(res, URLPost(longURL))
		} else {
			fmt.Fprint(res, "URL inválida")
		}
		//URLPost(longURL)
	}).Methods("POST").Queries("url", "{name}")
}

func Servidor() {
	router := mux.NewRouter()   //instância de mux
	const port string = ":8000" //port do server

	MetodoGet(router)

	MetodoPost(router)

	log.Println("servidor up")
	log.Fatal(http.ListenAndServe(port, router))
	//ListenAndServe inicia um server HTTP com um endereço e um handler
}
