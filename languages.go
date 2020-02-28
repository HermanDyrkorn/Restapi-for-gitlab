package assignment2

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

//HandlerLanguages function that deals with the / endpoint
func HandlerLanguages(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	//declaring variables in use
	APIURL := "https://git.gvk.idi.ntnu.no/api/v4/projects/"
	client := http.DefaultClient
	var lang = []string{}
	var langNameAndCount = []Language{}
	var returnLang = ReturnLanguage{}
	var projectPayload = []int{}
	var projects []ProjectID

	//splitting the path
	parts := strings.Split(r.URL.Path, "/")

	//checks the authentication
	var authentication = false
	auth := Auth(r)
	if auth != "" {
		authentication = true
	}

	switch r.Method {
	case http.MethodGet:

		//gets the limit
		limit, err := GetQueryLimit(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//gettig respons on the given url
		resp := GetTheRequest(APIURL+"?per_page=100"+"&private_token="+auth, client)
		totalProjects, err1 := strconv.Atoi(resp.Header.Get("X-Total-Pages"))
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		//looping over pages with projects
		for i := 1; i <= totalProjects; i++ {
			pro := []ProjectID{}
			resp := GetTheRequest(APIURL+"?per_page=100"+"&page="+strconv.Itoa(i)+"&private_token="+auth, client)
			err := json.NewDecoder(resp.Body).Decode(&pro)
			//checking for errors
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			//appending every page with projects to an array that contains all the projects
			for k := range pro {
				projects = append(projects, pro[k])
			}
			defer resp.Body.Close()

		}

		//loop that first fills a map with programming languages
		//and then filling a slice of strings with evry programming language
		for _, v := range projects {
			lMap := make(map[string]float64)
			URL2 := APIURL + strconv.Itoa(v.ID) + "/languages/" + "?private_token=" + auth
			resp := GetTheRequest(URL2, client)

			//checking for errors
			err := json.NewDecoder(resp.Body).Decode(&lMap)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			for k := range lMap {
				lang = append(lang, k)
			}
			defer resp.Body.Close()

		}

		//counts evry unique language
		countLang := Dupcount(lang)

		//converting the map into an array with structs
		langNameAndCount = convertToArray(countLang, langNameAndCount)

		//sorting the array so that the language with most counts comes first and so on
		sort.Slice(langNameAndCount, func(i, j int) bool {
			return langNameAndCount[i].Count > langNameAndCount[j].Count
		})

		//makeing sure that the limit cant be higher than the amount of projects
		if limit > len(langNameAndCount) {
			limit = len(langNameAndCount)
		}

		//appending the languages to a based on the limit
		for i := 0; i < limit; i++ {
			returnLang.Languages = append(returnLang.Languages, langNameAndCount[i].Name)
		}
		returnLang.Auth = authentication
		//encoding the languages to the user
		err2 := json.NewEncoder(w).Encode(returnLang)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return
		}
		//invoking potensial webhooks
		CheckWebhook(r, parts[3])
	case http.MethodPost:

		//decode the body that was sent by the user
		json.NewDecoder(r.Body).Decode(&projectPayload)

		//loop that first fills a map with programming languages
		//and then filling a slice of strings with every programming language
		for _, v := range projectPayload {
			lMap := make(map[string]float64)
			URL2 := APIURL + strconv.Itoa(v) + "/languages/" + "?private_token=" + auth
			resp := GetTheRequest(URL2, client)

			//checking for errors
			err := json.NewDecoder(resp.Body).Decode(&lMap)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			for k := range lMap {
				lang = append(lang, k)
			}
			defer resp.Body.Close()

		}

		//counts evry unique language
		countLang := Dupcount(lang)

		//converting the map into an array with structs
		langNameAndCount = convertToArray(countLang, langNameAndCount)

		//sorting the array so that the language with most counts comes first and so on
		sort.Slice(langNameAndCount, func(i, j int) bool {
			return langNameAndCount[i].Count > langNameAndCount[j].Count
		})

		//appending all the languages on the return struct
		for i := 0; i < len(langNameAndCount); i++ {
			returnLang.Languages = append(returnLang.Languages, langNameAndCount[i].Name)
		}
		//filling the returnstruct with auth
		returnLang.Auth = authentication
		//encoding the languages to the user
		err2 := json.NewEncoder(w).Encode(returnLang)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return
		}

		//invoking potensial webhooks
		CheckWebhook(r, parts[3])

	default:
		http.Error(w, "Invalid method"+r.Method, http.StatusBadRequest)

	}
}

//Dupcount function that count unique languages
func Dupcount(list []string) map[string]int {
	duplicateFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicateFrequency[item]
		if exist {
			// increase counter by 1 if already in the map
			duplicateFrequency[item]++
		} else {
			// else start counting from 1
			duplicateFrequency[item] = 1
		}
	}
	return duplicateFrequency
}

//converts a map into an array with structs
func convertToArray(m map[string]int, nAndC []Language) []Language {
	for k, v := range m {
		singleLanguage := Language{}
		singleLanguage.Name = k
		singleLanguage.Count = v
		nAndC = append(nAndC, singleLanguage)
	}
	return nAndC
}
