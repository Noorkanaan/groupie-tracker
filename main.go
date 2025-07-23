package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"main.go/internal"
)

var (
	tmpl     *template.Template
	artists  []internal.Artist
	relation internal.Relation
)

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", Home)
	http.HandleFunc("/details", Details)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	var err error
	artists, err = internal.FetchArtists()
	if err != nil {
		fmt.Println("Error handling request:", err)
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", artists)
}

func Details(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var selectedArtist internal.Artist
	for _, artist := range artists {
		if strconv.Itoa(artist.ID) == id {
			selectedArtist = artist
			break
		}
	}

	relation, _ = internal.FetchRelation(id)
	artistDetails := internal.ArtistDetails{
		RetArtist:   selectedArtist,
		RetRelation: relation,
	}
	// fmt.Println("artistDetails:", artistDetails)
	// fmt.Println("Artist:", artistDetails.RetArtist)
	// fmt.Println("Relation:", artistDetails.RetRelation)

	tmpl.ExecuteTemplate(w, "details.html", artistDetails)
}

// func searchHandler(w http.ResponseWriter, r *http.Request) {
// 	query := strings.ToLower(r.URL.Query().Get("query"))
// 	var results []string
// 	for _, artist := range artists {

// 		if strings.Contains(strings.ToLower(artist.Name), query) {
// 			results = append(results, fmt.Sprintf("%s - artist/band", artist.Name))
// 		}

// 		for _, member := range artist.Members {
// 			if strings.Contains(strings.ToLower(member), query) {
// 				results = append(results, fmt.Sprintf("%s - member", member))
// 			}
// 		}

// 		for _, location := range artist.LocationsUrl {
// 			if strings.Contains(strings.ToLower(artist.LocationsUrl), query) {
// 				results = append(results, fmt.Sprintf("%s - locationUrl", location))
// 			}
// 		}

// 		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
// 			results = append(results, fmt.Sprintf("%s - first album date", artist.FirstAlbum))
// 		}

// 		if strings.Contains(strings.ToLower(artist.ConcertDatesUrl), query) {
// 			results = append(results, fmt.Sprintf("%s - ConcertDatesUrl", artist.CreationDate))
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(results)
// }
