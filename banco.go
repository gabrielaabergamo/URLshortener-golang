package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//inserir queries no sql
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func chamada() {
	//abrir banco
	db, err := sql.Open("mysql", "root:<senha>@/")
	if err != nil {
		panic(err)
	}
	//fechar banco
	defer db.Close() //defer: is used to ensure that a function call is performed later in a program's execution
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
