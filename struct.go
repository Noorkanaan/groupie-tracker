package internal

type ArtistDetails struct {
	RetArtist   Artist
	RetRelation Relation
}

type Artist struct {
	ID              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsUrl    string   `json:"locations"`
	ConcertDatesUrl string   `json:"concertDates"`
	RelationsUrl    string   `json:"relations"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
