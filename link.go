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

//função executada no método POST: checamos no banco de dados se tal URL já existe e caso contrário a adicionamos
func URLPost(url string) string {
	start := time.Now()

	achou := buscarURL(url) //busca no bd

	//caso url já tenha sido inserida
	if achou.ID != "" {
		achou.ProcessedAt = start
		achou.Duration = time.Since(start)
		return TransfJson(achou)
	}

	//caso não tenha sido inserida ainda
	ID, OriginalURL, ShortURL, CodigoSURL := URLCurta(url)
	aux := inserirURL(ID, OriginalURL, ShortURL, CodigoSURL) //add no bd
	aux.ProcessedAt = start
	aux.Duration = time.Since(start)
	return TransfJson(aux)
}

//função que encurta a URL
func URLCurta(txt string) (string, string, string, string) {
	ID := uuid.NewV4().String() //gerar ID único

	OriginalURL := txt

	CodigoSURL := uniuri.NewLen(6) //gerar código da URL encurtada único
	lista := verificarCodigoBD()
	for ChecarCodigo(CodigoSURL, lista) {
		CodigoSURL = uniuri.NewLen(6)
	}

	ShortURL := "go.io/" + CodigoSURL

	return ID, OriginalURL, ShortURL, CodigoSURL
}

//checa se o código gerado é único
func ChecarCodigo(codigo string, lista []string) bool {
	for _, value := range lista {
		match, _ := regexp.MatchString(value, codigo)
		if match {
			return true
		}
	}
	return false
}

//função executada no método GET: checamos em listURL qual struct desejamos retornar
func URLGet(url string) string {
	start := time.Now()

	//utilizando o código da url encurtada é feita a busca pelos dados da url
	achou := buscarURLCurta(url)

	//caso não exista tal url no banco
	if achou.ID == "" {
		log.Println("deu ruim família")
		return "URL não existe no banco de dados"
	}

	achou.ProcessedAt = start
	achou.Duration = time.Since(start)
	return TransfJson(achou)
}

//transforma a struct em JSON
func TransfJson(aux Url) string {
	auxjson, err := json.Marshal(aux)
	if err != nil {
		fmt.Println("erro", err)
	}
	return string(auxjson)
}
