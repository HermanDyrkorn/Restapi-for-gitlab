package assignment2

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

//Defultlimit const
const Defultlimit = "5"

//HandlerCommits function that deals with the
func HandlerCommits(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		http.Header.Add(w.Header(), "content-type", "application/json")
		//splitting path into parts
		parts := strings.Split(r.URL.Path, "/")
		//gets the limit
		limit, err := GetQueryLimit(r)
		var authentication = false
		auth := Auth(r)
		if auth != "" {
			authentication = true
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//gettig respons on the given url
		APIURL := "https://git.gvk.idi.ntnu.no/api/v4/projects/"
		client := http.DefaultClient
		resp := GetTheRequest(APIURL+"?per_page=100"+"&private_token="+auth, client)
		totalPages, err1 := strconv.Atoi(resp.Header.Get("X-Total-Pages"))
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
		}

		//projects array with structs
		var projects []Project

		for i := 1; i <= totalPages; i++ {
			pro := []Project{}
			resp := GetTheRequest(APIURL+"?per_page=100"+"&page="+strconv.Itoa(i)+"&private_token="+auth, client)
			err := json.NewDecoder(resp.Body).Decode(&pro)

			//checking for errors
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			for k := range pro {
				projects = append(projects, pro[k])
			}

		}

		//loops over all the projects and query their commits
		for k, v := range projects {
			URL2 := APIURL + strconv.Itoa(v.ID) + "/repository/commits/" + "?private_token=" + auth
			resp := GetTheRequest(URL2, client)
			projects[k].Commits, _ = strconv.Atoi(resp.Header.Get("X-Total"))
		}

		//sorting the array with structs by its count, highest count = first in array
		sort.Slice(projects, func(i, j int) bool {
			return projects[i].Commits > projects[j].Commits
		})

		//makeing sure that the limit cant be higher than the amount of projects
		if limit > len(projects) {
			limit = len(projects)
		}

		//declaring a output array that takes projectstrucks and it is limit long
		var projectOutput = make([]Projectoutput, limit)

		//filling the output array with structs
		for k := range projectOutput {
			projectOutput[k].Name = projects[k].Name
			projectOutput[k].CommitCount = projects[k].Commits
		}

		//filling an array inside a struct with the right output
		repo := &RepoOutput{projectOutput, authentication}
		err2 := json.NewEncoder(w).Encode(repo)

		//checking for errors
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
		}
		defer resp.Body.Close()

		//invoking potensial urls
		CheckWebhook(r, parts[3])
	default:
		http.Error(w, "Invalid method"+r.Method, http.StatusBadRequest)

	}
}
