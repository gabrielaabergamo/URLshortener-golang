package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dchest/uniuri"
	uuid "github.com/satori/go.uuid"
)

//struct com os dados da URL
type Url struct {
	ID          string
	ProcessedAt string
	Duration    string
	ShortURL    string
	OriginalURL string
}

var listaURL = make([]Url, 0)

//struct teste
var aux = Url{
	ShortURL:    "nenhuma",
	OriginalURL: "oi",
}

//função que encurta a URL
func URLCurta(txt string) {
	aux.ID = uuid.NewV4().String()
	aux.OriginalURL = txt
	aux.ShortURL = "go.io/" + uniuri.NewLen(6)
	TimeCalc()
	listaURL = append(listaURL, aux)
}

//calcula o dia e horário em que foi processado e a duração
func TimeCalc() {
	start := time.Now()
	aux.ProcessedAt = start.Format("2006-01-02 15:04:05")
	aux.Duration = time.Since(start).String()
}

//transforma a struct em JSON
func TransfJson() string {
	auxjson, err := json.Marshal(aux)
	if err != nil {
		fmt.Println("erro", err)
	}
	//printslice()
	return string(auxjson)
}

//teste para checar se os dados estão sendo salvos corretamente
// func printslice() {
// 	for _, i := range listaURL {
// 		log.Println(i)
// 	}
// }
