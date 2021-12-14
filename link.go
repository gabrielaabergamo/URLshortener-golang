package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/dchest/uniuri"
	uuid "github.com/satori/go.uuid"
)

//struct com os dados da URL
type Url struct {
	ID          string
	ProcessedAt time.Time
	Duration    time.Duration
	OriginalURL string
	ShortURL    string
	CodigoSURL  string `json:"-"`
}

//lista de structs para armazenar dados da URL
var listaURL = make([]Url, 0)

//função executada no método POST: checamos em listaURL se tal URL já existe e caso contrário a adicionamos
func URLPost(url string) string {
	// inserirURL()
	start := time.Now()

	//checar se já existe essa URL
	achou, _ := ChecarURL(url) //substituir por uma consulta no bd

	if achou {
		//return "URL já está no banco de dados"
		_, matchIndice := ChecarURL(url)
		listaURL[matchIndice].ProcessedAt = start
		listaURL[matchIndice].Duration = time.Since(start)
		return TransfJson(listaURL[matchIndice])
	}

	//caso não exista:
	structURL := URLCurta(url)
	structURL.ProcessedAt = start
	structURL.Duration = time.Since(start)
	log.Println("nao existe")
	printslice()
	chamada(structURL.ID, structURL.OriginalURL, structURL.ShortURL, structURL.CodigoSURL) //add no bd
	return TransfJson(structURL)
}

//função que encurta a URL
func URLCurta(txt string) Url {
	aux := Url{}
	aux.ID = uuid.NewV4().String()
	aux.OriginalURL = txt

	codigo := uniuri.NewLen(6)
	for ChecarCodigo(codigo) {
		codigo = uniuri.NewLen(6)
	} //fica no slice ou bd? >bd

	aux.CodigoSURL = codigo
	aux.ShortURL = "go.io/" + aux.CodigoSURL
	listaURL = append(listaURL, aux)
	return aux
}

//checa se existe tal URL em listaURL e qual seu índice
func ChecarURL(url string) (bool, int) {
	for i, value := range listaURL {
		match, _ := regexp.MatchString(value.OriginalURL, url)
		if match {
			return true, i
		}
	}
	return false, -1
}

//checa se existe tal URL em listaURL e qual seu índice
func ChecarURLEncurtada(url string) (bool, int) {
	for i, value := range listaURL {
		match, _ := regexp.MatchString(value.CodigoSURL, url)
		if match {
			return true, i
		}
	}
	return false, -1
}

//checa se o código gerado é único
func ChecarCodigo(codigo string) bool {
	for _, value := range listaURL {
		match, _ := regexp.MatchString(value.CodigoSURL, codigo)
		if match {
			return true
		}
	}
	return false
}

//função executada no método GET: checamos em listURL qual struct desejamos retornar
func URLGet(url string) string {
	start := time.Now()
	//fazer busca para achar a struct que queremos
	_, indiceURL := ChecarURLEncurtada(url)

	if indiceURL == -1 {
		log.Println("deu ruim família")
		return "URL não existe no banco de dados"
	}

	listaURL[indiceURL].ProcessedAt = start
	listaURL[indiceURL].Duration = time.Since(start)
	return TransfJson(listaURL[indiceURL])
}

//transforma a struct em JSON
func TransfJson(aux Url) string {
	auxjson, err := json.Marshal(aux)
	if err != nil {
		fmt.Println("erro", err)
	}
	//printslice()
	return string(auxjson)
}

//teste para checar se os dados estão sendo salvos corretamente
func printslice() {
	for _, i := range listaURL {
		log.Println(i)
	}
}
