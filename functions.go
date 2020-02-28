package assignment2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

//GetTheRequest function returns a respons from the api
func GetTheRequest(URL string, client *http.Client) *http.Response {
	//do a request on the url
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Println(err)
	}
	//getting response from the client
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	//returning the response
	return resp
}

//HandlerNil function that deals with the / endpoint
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid request, expected formating /repocheck/v1/", http.StatusBadRequest)
}

//CheckWebhook function checks if the event has a webhook, invokes if ture
func CheckWebhook(r *http.Request, event string) {
	//getting every registerd webhook form the database
	webhooks, err := ReturnAllWebhooks()
	if err != nil {
		fmt.Println("Error in HTTP request: " + err.Error())
	}
	//ranging over all webhooks and checks if the event is the same as in all the webhooks
	for _, v := range webhooks {
		if v.Event == event {
			CallURL(v.URL, "Event: "+v.Event+", Timestamp: "+time.Now().String()+", Parameters: "+r.URL.RawQuery)
		}
	}
}

//CallURL function calls given URL with given content and awaits response (status and body).
func CallURL(url string, content string) {
	//post on the url that was sent as a parameter
	res, err := http.Post(url, "string", bytes.NewReader([]byte(content)))
	if err != nil {
		fmt.Println("Error in HTTP request: " + err.Error())
		return
	}
	response, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		fmt.Println("Something is wrong with invocation response: " + err.Error())
		return
	}

	fmt.Println("Webhook invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(response))
}

//GetQueryLimit function gets the limit param of returns an error
func GetQueryLimit(r *http.Request) (int, error) {
	var limit string
	query := r.URL.Query()
	reg, _ := regexp.Compile("[0-9]+")

	if len(query["limit"]) > 0 {
		if reg.MatchString(query["limit"][0]) {
			limit = query["limit"][0]
		} else {
			limit = Defultlimit
		}
	} else {
		limit = Defultlimit
	}
	return strconv.Atoi(limit)
}

//Auth function that checks the authentication
func Auth(r *http.Request) string {
	query := r.URL.Query()
	if len(query["auth"]) > 0 {
		return query["auth"][0]
	}
	return ""
}
