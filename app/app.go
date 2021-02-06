package app

import (
	"log"
	"net/http"

	"../controllers"
)

func StartApp() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":5050", nil); err != nil {
		//log.Fatalf("App Stop with: %v", err)
		panic(err)
	}
}
