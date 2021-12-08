package main

type url struct {
	id          int
	processedAt float64
	duration    float64
	shortURL    string
	originalURL string
}

var aux = url{
	id:          1,
	processedAt: 1.0,
	duration:    1.0,
	shortURL:    "nenhuma",
	originalURL: "oi",
}

func Testando(txt string) {
	aux.shortURL = txt
}
