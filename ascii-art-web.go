package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// PageVariables holds data to be passed to HTML templates
type PageVariables struct {
	Response       string
	Input          string
	SelectedBanner string
	SpecialTrigger bool
}

// main starts the HTTP server and registers routes
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Register route handlers
	http.HandleFunc("/", GetHandler) // Define the route and its handler function
	http.HandleFunc("/ascii-art", PostHandler)
	//start the server on port 8080
	log.Println("Starting server on: http://localhost:8080")
	log.Println("Status ok: ", http.StatusOK)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// GetHandler handles GET requests and serves the main page
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.Method != http.MethodGet { // Ensure the method is GET
		badRequestError(w)
		return
	}
	// If not "http://localhost:8080"
	if r.URL.Path != "/" {
		notFoundError(w)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		internalServerError(w)
		return
	}
	// Render the template safely with status 200
	log.Printf("Response Status: %d\n", http.StatusOK)
	err = safeRenderTemplate(w, tmpl, "index.html", http.StatusOK, nil)
	if err != nil {
		internalServerError(w)
		return
	}
}

// PostHandler handles POST requests to generate ASCII art
func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	if r.Method != http.MethodPost {
		badRequestError(w)
		return
	}
	if r.URL.Path != "/ascii-art" {
		notFoundError(w)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		internalServerError(w)
		return
	}
	// Read user input and banner selection
	r.ParseForm()
	inputText := r.FormValue("inputField")
	banner := r.FormValue("banner")
	var vars PageVariables
	log.Println("User input is: ", inputText, "Selected banner is: ", banner)
	// Generate ASCII art
	response := AsciiArt(inputText, banner)

	// Handle errors from AsciiArt function
	validBanners := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}
	if strings.HasPrefix(response, "Error reading file:") {
		if validBanners[banner] {
			internalServerError(w)
		} else {
			notFoundError(w)
		}
		return
	}
	if strings.HasPrefix(response, "Invalid characters") {
		vars = PageVariables{Response: response, Input: "\n" + inputText, SelectedBanner: banner, SpecialTrigger: true}
	} else {
		vars = PageVariables{Response: response, Input: "\n" + inputText, SelectedBanner: banner, SpecialTrigger: false}
	}
	// Render the template with ASCII art
	//vars = PageVariables{Response: response, Input: inputText, SelectedBanner: banner}
	log.Printf("Response Status: %d\n", http.StatusOK)
	err = safeRenderTemplate(w, tmpl, "index.html", http.StatusOK, vars)
	if err != nil {
		internalServerError(w)
		return
	}
}

// badRequestError serves a 400 error page
func badRequestError(w http.ResponseWriter) {
	log.Printf("Response Status: %d\n", http.StatusBadRequest)
	tmpl, err := template.ParseFiles("templates/badRequest.html")
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = safeRenderTemplate(w, tmpl, "badRequest.html", http.StatusBadRequest, nil)
	if err != nil {
		internalServerError(w)
		return
	}
}

// notFoundError serves a 404 error page
func notFoundError(w http.ResponseWriter) {
	log.Printf("Response Status: %d\n", http.StatusNotFound)
	tmpl, err := template.ParseFiles("templates/notFound.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	err = safeRenderTemplate(w, tmpl, "notFound.html", http.StatusNotFound, nil)
	if err != nil {
		internalServerError(w)
		return
	}
}

// internalServerError serves a 500 error page
func internalServerError(w http.ResponseWriter) {
	log.Printf("Response Status: %d\n", http.StatusInternalServerError)
	tmpl, err := template.ParseFiles("templates/internalServer.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = safeRenderTemplate(w, tmpl, "internalServer.html", http.StatusInternalServerError, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// safeRenderTemplate renders a template safely and writes to the response
func safeRenderTemplate(w http.ResponseWriter, tmpl *template.Template, templateName string, status int, data any) error {
	var buffer bytes.Buffer
	err := tmpl.ExecuteTemplate(&buffer, templateName, data)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	buffer.WriteTo(w)
	return nil
}
