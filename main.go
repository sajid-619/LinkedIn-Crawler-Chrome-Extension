// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Profile represents the MongoDB document structure
type Profile struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	LinkedInURL string        `bson:"linkedin_url"`
	Description string        `bson:"description"`
}

var session *mgo.Session
var profilesCollection *mgo.Collection

func init() {
	var err error
	// Update the MongoDB connection string based on your setup
	session, err = mgo.Dial("mongodb://localhost:27017/linkedin_scraper")
	if err != nil {
		log.Fatal(err)
	}

	// Set up the "profiles" collection in MongoDB
	profilesCollection = session.DB("linkedin-scraper").C("profiles")
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/profileInfo/get", GetProfileInfo)

	http.ListenAndServe(":8080", r)
}

func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		http.Error(w, "Path parameter is required", http.StatusBadRequest)
		return
	}

	// Check if the profile is already in the database
	existingProfile, err := getProfileFromDatabase(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking existing profile: %s", err), http.StatusInternalServerError)
		return
	}

	if existingProfile != nil {
		// Profile already in the database, return the stored description
		response := map[string]string{"description": existingProfile.Description}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get LinkedIn profile information
	description, err := ScrapeLinkedInProfile(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scraping LinkedIn: %s", err), http.StatusInternalServerError)
		return
	}

	// Save the new profile to the database
	err = saveProfileToDatabase(path, description)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving profile to database: %s", err), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	response := map[string]string{"description": description}
	json.NewEncoder(w).Encode(response)
}

func ScrapeLinkedInProfile(path string) (string, error) {
	// Assuming LinkedIn URL is in the format "/in/dev-mirian-quispe/"
	// Modify as needed based on your actual URL format
	url := "https://www.linkedin.com" + path

	// Use Goquery to scrape LinkedIn profile
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	// Modify the CSS selector based on your actual requirement
	description := doc.Find("div.entity-result__primary-subtitle.t-14.t-black.t-normal").Text()

	return strings.TrimSpace(description), nil
}

func getProfileFromDatabase(path string) (*Profile, error) {
	var profile Profile
	err := profilesCollection.Find(bson.M{"linkedin_url": path}).One(&profile)
	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	return &profile, nil
}

func saveProfileToDatabase(path, description string) error {
	profile := Profile{
		LinkedInURL: path,
		Description: description,
	}

	err := profilesCollection.Insert(profile)
	return err
}
