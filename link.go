package main

import (
	"time"
)

type url struct {
	id          int
	processedAt string
	duration    string
	shortURL    string
	originalURL string
}

var aux = url{
	id:          1,
	shortURL:    "nenhuma",
	originalURL: "oi",
}

func Testando(txt string) {
	aux.shortURL = txt
	TimeCalc()
}

//calcula o dia e horário em que foi processado e a duração
func TimeCalc() {
	start := time.Now()
	aux.processedAt = start.Format("2006-01-02 15:04:05")
	aux.duration = time.Since(start).String()
}
