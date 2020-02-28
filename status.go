package assignment2

import (
	"encoding/json"
	"net/http"
	"strings"
)

//HandlerStatus function that deals with the / endpoint
func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	//creating an empty status struct
	status := &Status{}

	//filling the status struct with info
	status.Uptime = Uptime()
	status.Version = "V1"
	status.Database, _ = GetStatusCode("https://console.firebase.google.com")
	status.Gitlab, _ = GetStatusCode("https://git.gvk.idi.ntnu.no/api")
	//encode the status struct back to the user
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
	}
	CheckWebhook(r, parts[3])
	
}

//GetStatusCode function
func GetStatusCode(URL string) (int, error) {
	//getrequest on a given url
	resp, err := http.Get(URL)
	if err != nil {
		return 0, err
	}
	//returning the statuscode of the response
	return resp.StatusCode, nil
}
