package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error fetching artists:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	var artists []Artist

	//fmt.Println("JSON Response:", string(body))
	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Println("Error unmarshalling artists:", err)
	}
	//fmt.Println("Artists:", artists)

	filteredArtists := make([]Artist, 0, len(artists))
	for _, artist := range artists {
		if artist.ID != 21 {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	return filteredArtists, nil
}

func FetchRelation(artistID string) (Relation, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + artistID)

	if err != nil {
		return Relation{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Relation{}, err
	}

	var relation Relation

	//fmt.Println("JSON Response:", string(body))
	err = json.Unmarshal(body, &relation)
	if err != nil {
		return Relation{}, err
	}

	return relation, nil
}
