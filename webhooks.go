package assignment2

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

//HandlerWebhook function
func HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case http.MethodPost:
		//Decoding the payload
		webhook := WebhookRegistration{}
		err := json.NewDecoder(r.Body).Decode(&webhook)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
		}
		//adding timestamp and saved the webhook to the database
		webhook.Timestamp = time.Now().String()
		Save(&webhook)
	case http.MethodGet:
		//if there is not a webhook id as parameters
		//every webhook is returned
		if parts[4] == "" {
			GetData(w)
		} else {
			//if there is a webhook id as parameter
			//this webhook is returned
			webhook, err := FindID(parts[4])
			if err != nil {
				http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
			}
			json.NewEncoder(w).Encode(webhook)
		}

	case http.MethodDelete:
		//delete the webhook with id in the parameter
		err := DeleteWebhook(parts[4])
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
		}

	}
}
