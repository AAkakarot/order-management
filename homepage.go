package main

import (
	"log"
	"net/http"
)

func HomePage(http.ResponseWriter, *http.Request) {

	log.Println("Welcome to the HomePage!")
}
