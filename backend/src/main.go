package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Description string `json:"deescription"`
}

type IndexPage struct {
	Title  string
	Detail map[string]RepoMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var a []Repos
	repo_map := make(map[string]RepoMap)
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	//err := json.NewDecoder(resp.Body).Decode(&a)
	resp.Body.Close()

	err := json.Unmarshal(bytes, &a)

	if err != nil {
		fmt.Println(err)
	}

	for _, data := range a {
		repo_map[fmt.Sprint(data.Id)] = RepoMap{Name: data.Name, Owner: data.Owner.Login, Description: data.Description}

	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bytes))
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

}
