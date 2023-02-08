package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := http.Get("https://api.github.com/users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(list)
}

func GetDetail(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := strings.TrimPrefix(req.URL.Path, "/detail/")

	resp, err := http.Get("https://api.github.com/users/" + name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(list)
}

func StartRouter(port string) {
	http.HandleFunc("/list", GetList)
	http.HandleFunc("/detail/", GetDetail)

	fmt.Println("Application is running on port : " + port)
	fmt.Println("List of resource path")
	fmt.Println("1. /list")
	fmt.Println("2. /detail/:name")

	http.ListenAndServe(":"+port, nil)
}

func main() {
	commands := os.Args
	if len(commands) < 2 {
		fmt.Println("ERROR: Please specify the port (Recomended : 3000 - 10000)")
		return
	}

	port := commands[1]

	_, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("ERROR: Port should be number")
		return
	}

	StartRouter(port)
}
