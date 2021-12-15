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

func inserirURL(id, OriginalURL, ShortURL, CodigoSURL string) Url {
	//adicionar url no banco
	stmt := `INSERT INTO urls(id, url_original, url_short, url_short_sufix) VALUES (?, ?, ?, ?)`
	res, _ := db.Exec(stmt, id, OriginalURL, ShortURL, CodigoSURL)
	log.Println(res)
	return buscarURL(OriginalURL)
}

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
func verificarCodigoBD() []string {
	//buscar códigos no bd
	rows, _ := db.Query("select url_short_sufix from urls")
	defer rows.Close()
	lista := make([]string, 0)
	for rows.Next() {
		aux := ""
		rows.Scan(aux)
		lista = append(lista, aux)
	}
	return lista
}

// func chamada(id, OriginalURL, ShortURL, CodigoSURL string) {
// 	// getEnvVars()
// 	// username := os.Getenv("USERNAME")
// 	// password := os.Getenv("PASSWORD")
// 	// //abrir banco
// 	// db, err := sql.Open("mysql", username+":"+password+"@/")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// //fechar banco
// 	// defer db.Close() //defer: is used to ensure that a function call is performed later in a program's execution
// 	// exec(db, "create database if not exists urlshortener")
// 	// exec(db, "use urlshortener")
// 	// exec(db, "drop table if exists urls")
// 	// exec(db, `create table urls (
// 	// 	id varchar(80),
// 	// 	url_original varchar(80) UNIQUE,
// 	// 	url_short varchar(80) UNIQUE,
// 	// 	url_short_sufix varchar(80) UNIQUE,
// 	// 	PRIMARY KEY (id)
// 	// 	)`)

// 	// stmt, _ := db.Prepare("insert into urls(id, url_original, url_short, url_short_sufix) values(?)")
// 	// stmt.Exec("teste", "teste", "teste", "teste")

// 	// //adicionar url no banco
// 	// stmt := `INSERT INTO urls(id, url_original, url_short, url_short_sufix) VALUES (?, ?, ?, ?)`
// 	// res, _ := db.Exec(stmt, id, OriginalURL, ShortURL, CodigoSURL)
// 	// log.Println(res)

// 	//buscar por url no banco
// 	// rows, _ := db.Query("select id, url_original, url_short, url_short_sufix from urls where id = ?", id)
// 	// //Query executes a query that returns rows, typically a SELECT.
// 	// defer rows.Close()

// 	// for rows.Next() {
// 	// 	var u Urlteste
// 	// 	rows.Scan(&u.ID, &u.OriginalURL, &u.ShortURL, &u.CodigoSURL)
// 	// 	log.Println(u)
// 	// }
// }
