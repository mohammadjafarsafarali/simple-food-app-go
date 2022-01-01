package main

import (
	"github.com/joho/godotenv"
	"log"
	_ "log"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}


}
