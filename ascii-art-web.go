package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Define a struct to hold page variables
type PageVariables struct {
	Response string
	Input    string
}

func main() {
	fmt.Println("Server is running on http://localhost:8080")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/ascii-art", PostHandler)
	http.ListenAndServe(":8080", nil)
}

// homePage serves the main HTML page
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.Method != http.MethodGet {
		//http.Error(w, "Bad Request", http.StatusBadRequest) // 400
		badRequest(w)
		return
	}
	if r.URL.Path != "/" {
		notFound(w)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		internalServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	log.Printf("Response Status: %d\n", http.StatusOK)
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

// handleAsciiArt processes the POST request to generate ASCII art
func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.Method != http.MethodPost {
		//http.Error(w, "Bad Request", http.StatusBadRequest) // 400
		badRequest(w)
		return
	}
	if !(r.URL.Path == "/" || r.URL.Path == "/ascii-art") {
		notFound(w)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		internalServerError(w)
		return
	}
	r.ParseForm()
	inputText := r.FormValue("inputField")
	banner := "banners/"
	banner += r.FormValue("banner")
	response := AsciiArt(inputText, banner)
	if strings.HasPrefix(response, "Error reading file:") { // Example error check
		//http.Error(w, response, http.StatusNotFound) // 404
		notFound(w)
		return
	}
	if response == "Invalid input" {
		badRequest(w)
		return
	}
	vars := PageVariables{Response: response, Input: inputText}
	w.WriteHeader(http.StatusOK) // 200
	log.Printf("Response Status: %d\n", http.StatusOK)
	tmpl.Execute(w, vars)
}

func notFound(w http.ResponseWriter) {
	log.Printf("Response Status: %d\n", http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)
	tmpl, err := template.ParseFiles("templates/notFound.html")
	if err != nil {
		internalServerError(w)
		return
	}
	tmpl.ExecuteTemplate(w, "notFound.html", nil)
}

func badRequest(w http.ResponseWriter) {
	log.Printf("Response Status: %d\n", http.StatusBadRequest)
	w.WriteHeader(http.StatusBadRequest)
	tmpl, err := template.ParseFiles("templates/badRequest.html")
	if err != nil {
		internalServerError(w)
		return
	}
	tmpl.ExecuteTemplate(w, "badRequest.html", nil)
}

func internalServerError(w http.ResponseWriter) {
	log.Printf("Response Status: %d\n", http.StatusInternalServerError)
	w.WriteHeader(http.StatusInternalServerError)
	tmpl, err := template.ParseFiles("templates/internalServer.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}
	tmpl.ExecuteTemplate(w, "internalServer.html", nil)
}
