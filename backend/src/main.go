package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const url string = "https://api.github.com/users/enotgnik/repos"

type Repos struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       struct {
		Login string `json:"login"`
	} `json:"owner"`
}

type RepoMap struct {
	Name        string
	Owner       string
	Description string
}

type IndexPage struct {
	Title  string
	Detail map[string]RepoMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var a []Repos
	repo_map := make(map[string]RepoMap)
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	err := json.Unmarshal(bytes, &a)
	if err != nil {
		fmt.Println(err)
	}

	for _, data := range a {
		repo_map[fmt.Sprint(data.Id)] = RepoMap{data.Name, data.Owner.Login, data.Description}

	}
	fmt.Fprint(w, repo_map)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", r)
}
