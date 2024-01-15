// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const linkedinURL = "https://www.linkedin.com/search/results/people/?heroEntityKey=urn%3Ali%3Aorganization%3A12577023&keywords=wivenn&origin=CLUSTER_EXPANSION&position=1&searchId=1ee62608-5740-4e35-92a3-f4c907067491&sid=VjS"
const apiURL = "https://recrutador.wivenn.com.br/api/profileInfo/get?path="

type APIResponse struct {
	Description string `json:"description"`
}

func main() {

	// Create a new CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // You might want to restrict this to specific origins in a production environment
		AllowedMethods: []string{"GET"},
	 })

	r := mux.NewRouter()
	r.HandleFunc("/api/profileInfo/get", GetProfileInfo).Methods("GET")

	// Serve static files (your Chrome Extension files)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Start the server
	port := ":8080"
	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	// Extract the dynamic user URI from the query parameters
	userURI := r.URL.Query().Get("path")

	// Construct the LinkedIn profile URL
	linkedinProfileURL := linkedinURL + userURI

	// Scrape the LinkedIn profile page
	description, err := scrapeLinkedInProfile(linkedinProfileURL)
	if err != nil {
		// Handle error (e.g., 404 not found)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Return the description in JSON format
	response := APIResponse{Description: description}
	json.NewEncoder(w).Encode(response)
}

func scrapeLinkedInProfile(url string) (string, error) {
	// Fetch the HTML content of the LinkedIn profile page
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	// Find and extract the description using the specified CSS selector
	description := doc.Find("div.entity-result__primary-subtitle.t-14.t-black.t-normal").Text()

	// Trim spaces and newlines from the description
	description = strings.TrimSpace(description)

	return description, nil
}
