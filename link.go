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
	ID, OriginalURL, ShortURL, CodigoSURL := URLCurta(url)
	aux, err := inserirURL(ID, OriginalURL, ShortURL, CodigoSURL) //add no bd
	if err != nil {
		match, _ := regexp.MatchString("Error 1062", err.Error())
		if match {
			aux = buscarURL(url)
		} else {
			return "erro desconhecido"
		}

	}
	aux.ProcessedAt = start
	aux.Duration = time.Since(start)
	return TransfJson(aux)
}

//função que encurta a URL
func URLCurta(txt string) (string, string, string, string) {
	ID := GeradorUUID().String() //gerar ID único

	OriginalURL := txt

	CodigoSURL := GeradorCodigo() //gerar código da URL encurtada único
	verificacao := verificarCodigoBD(CodigoSURL)
	for len(verificacao) > 0 {
		CodigoSURL = GeradorCodigo()
		verificacao = verificarCodigoBD(CodigoSURL)
	}

	ShortURL := GeradorURL(CodigoSURL)

	return ID, OriginalURL, ShortURL, CodigoSURL
}

func GeradorUUID() uuid.UUID {
	return uuid.NewV4()
}

func GeradorCodigo() string {
	return uniuri.NewLen(6)
}

func GeradorURL(CodigoSURL string) string {
	return ("http://go.io/" + CodigoSURL)
}

//função executada no método GET: checamos em listURL qual struct desejamos retornar
func URLGet(url string) string {
	start := time.Now()

	//utilizando a url encurtada para fazer a busca pelos dados da url
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
		fmt.Println("erro: não foi possível transformar a struct em JSON", err)
	}
	return string(auxjson)
}
