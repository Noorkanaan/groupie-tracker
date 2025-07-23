package internal

import (
	"net/http"
	"strings"
)

var artists []Artist

func LoadArtists() error {
	var err error
	artists, err = FetchArtists()
	return err
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	query := strings.ToLower(r.FormValue("query"))
	var results []string

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, artist.Name+" - artist/band")
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				results = append(results, member+" - member")
			}
		}
		// Assuming locations and other fields are handled similarly
	}

}
