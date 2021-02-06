package main

import (
	"./app"
	"log"
)

func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app.StartApp()
}