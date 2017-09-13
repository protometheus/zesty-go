package main

import (
	"os"
    "fmt"
    "net/http"
	"encoding/json"
	"github.com/rs/cors"
)

type Name struct {
	First 	string	`json:"first"`
	Last 	string	`json:"last"`
}

type PersonalInfo struct {
	Name 				Name	`json:"name"`
	Github_repo_link 	string	`json:"github_repo_link"`
	Website 			string	`json:"website"`
	Email				string	`json:"email"`
}

func base_handler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/code/challenge", 301)
}

func endpoint_handler(w http.ResponseWriter, r *http.Request) {
	personal_info := PersonalInfo{Name{"Mohammad", "Oweis"}, "https://github.com/Protometheus/zesty-go",  "https://github.com/Protometheus/zesty-go", "oweismo.applicant@gmail.com"}

	js, _ := json.Marshal(personal_info)
	
	w.WriteHeader(201)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
}

func main() {
	mux := http.NewServeMux()
	
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	
	port = ":" + port
	
	fmt.Println(port)

    mux.HandleFunc("/", base_handler)
	mux.HandleFunc("/code/challenge", endpoint_handler)
	
	handler := cors.Default().Handler(mux)
	
    http.ListenAndServe(port, handler)
}
