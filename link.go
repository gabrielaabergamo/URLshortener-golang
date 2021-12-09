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

var aux = make([]Url, 10)

//struct teste
// var aux = Url{
// 	ShortURL:    "nenhuma",
// 	OriginalURL: "oi",
// }

//função que encurta a URL
func URLCurta(txt string) {
	aux[0].ID = uuid.NewV4().String()
	aux[0].OriginalURL = txt
	aux[0].ShortURL = "go.io/" + uniuri.NewLen(6)
	TimeCalc()
}

//calcula o dia e horário em que foi processado e a duração
func TimeCalc() {
	start := time.Now()
	aux[0].ProcessedAt = start.Format("2006-01-02 15:04:05")
	aux[0].Duration = time.Since(start).String()
}

//transforma a struct em JSON
func TransfJson() string {
	auxjson, err := json.Marshal(aux[0])
	if err != nil {
		fmt.Println("erro", err)
	}
	return string(auxjson)
}
