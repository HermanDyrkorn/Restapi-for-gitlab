package main

import (
	"assignment2"
	"fmt"
	"log"
	"net/http"
)

func main() {

	assignment2.InitTime()
	assignment2.InitDatabase()
	defer assignment2.DB.Client.Close()

	port := "8080"
	http.HandleFunc("/", assignment2.HandlerNil)
	http.HandleFunc("/repocheck/v1/commits/", assignment2.HandlerCommits)
	http.HandleFunc("/repocheck/v1/languages/", assignment2.HandlerLanguages)
	http.HandleFunc("/repocheck/v1/status/", assignment2.HandlerStatus)
	http.HandleFunc("/repocheck/v1/webhooks/", assignment2.HandlerWebhook)
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
