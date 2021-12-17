package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

var db *sql.DB

func getEnvVars() {
	err := godotenv.Load("credentials.env")

	if err != nil {
		log.Fatal("deu RUIMZAO")
	}
}

//inserir queries no sql
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func inicializaBD() {
	getEnvVars()
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	//abrir banco
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@/")
	if err != nil {
		panic(err)
	}
	exec(db, "create database if not exists urlshortener")
	exec(db, "use urlshortener")
	exec(db, "drop table if exists urls")
	exec(db, `create table urls (
		id varchar(80),
		url_original varchar(80) UNIQUE,
		url_short varchar(80) UNIQUE,
		url_short_sufix varchar(80) UNIQUE,
		PRIMARY KEY (id)
		)`)
}

//adicionar URL no banco de dados
func inserirURL(id, OriginalURL, ShortURL, CodigoSURL string) (Url, error) {
	stmt := `INSERT INTO urls(id, url_original, url_short, url_short_sufix) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(stmt, id, OriginalURL, ShortURL, CodigoSURL)

	if err != nil {
		aux := Url{}
		return aux, err
	}

	return buscarURL(OriginalURL), nil
}

//buscar por URL a partir de seu link original
func buscarURL(OriginalURL string) Url {
	//buscar por url no banco
	rows, _ := db.Query("select id, url_original, url_short, url_short_sufix from urls where url_original = ?", OriginalURL)
	//Query executes a query that returns rows, typically a SELECT.
	defer rows.Close()
	var u Url
	for rows.Next() {
		rows.Scan(&u.ID, &u.OriginalURL, &u.ShortURL, &u.CodigoSURL)
	}
	return u
}

//buscar por URL encurtada a partir de seu código
func buscarURLCurta(CodigoSURL string) Url {
	//buscar por url no banco
	rows, _ := db.Query("select id, url_original, url_short, url_short_sufix from urls where url_short_sufix = ?", CodigoSURL)
	//Query executes a query that returns rows, typically a SELECT.
	defer rows.Close()
	var u Url
	for rows.Next() {
		rows.Scan(&u.ID, &u.OriginalURL, &u.ShortURL, &u.CodigoSURL)
	}
	return u
}

//retorna um slice com todos os códigos de urls que estão salvas
func verificarCodigoBD(codigo string) string {
	//buscar códigos no bd
	// rows, _ := db.Query("select url_short_sufix from urls")
	// defer rows.Close()
	// lista := make([]string, 0)
	// for rows.Next() {
	// 	aux := ""
	// 	rows.Scan(&aux)
	// 	lista = append(lista, aux)
	// }

	// return lista
	rows, _ := db.Query("select url_short_sufix from urls where url_short_sufix= ?", codigo)
	defer rows.Close()
	aux := ""
	for rows.Next() {
		rows.Scan(&aux)
	}
	return aux
}
