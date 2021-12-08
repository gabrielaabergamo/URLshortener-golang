package main

import (
	"time"

	"github.com/dchest/uniuri"
	uuid "github.com/satori/go.uuid"
)

//struct com os dados da URL
type url struct {
	id          string
	processedAt string
	duration    string
	shortURL    string
	originalURL string
}

//struct teste
var aux = url{
	shortURL:    "nenhuma",
	originalURL: "oi",
}

//função que encurta a URL
func URLCurta(txt string) {
	aux.id = uuid.NewV4().String()
	aux.originalURL = txt
	aux.shortURL = "go.io/" + uniuri.NewLen(6)
	TimeCalc()
}

//calcula o dia e horário em que foi processado e a duração
func TimeCalc() {
	start := time.Now()
	aux.processedAt = start.Format("2006-01-02 15:04:05")
	aux.duration = time.Since(start).String()
}
