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

//struct teste
var aux = &Url{
	ShortURL:    "nenhuma",
	OriginalURL: "oi",
}

//função que encurta a URL
func URLCurta(txt string) {
	aux.ID = uuid.NewV4().String()
	aux.OriginalURL = txt
	aux.ShortURL = "go.io/" + uniuri.NewLen(6)
	TimeCalc()
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
	return string(auxjson)
}
