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

var listaURL = make([]Url, 0)

func URLPost(url string) {
	//checar se já existe essa URL
	achou, _ := ChecarURL(url)

	if !achou {
		//caso não exista:
		start := time.Now()
		structURL := URLCurta(url)
		structURL.ProcessedAt = start
		structURL.Duration = time.Since(start)
		log.Println("nao existe")
		printslice()
	}
}

//função que encurta a URL
func URLCurta(txt string) Url {
	aux := Url{}
	aux.ID = uuid.NewV4().String()
	aux.OriginalURL = txt
	aux.CodigoSURL = uniuri.NewLen(6)
	aux.ShortURL = "go.io/" + aux.CodigoSURL
	listaURL = append(listaURL, aux)
	return aux
}

func ChecarURL(url string) (bool, int) {
	for i, value := range listaURL {
		match, _ := regexp.MatchString(value.OriginalURL, url)
		if match {
			return true, i
		}
	}
	return false, -1
}

func URLGet(url string) string {
	start := time.Now()
	//fazer busca para achar a struct que queremos
	_, indiceURL := ChecarURL(url)

	if indiceURL == -1 {
		log.Println("deu ruim família")
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
