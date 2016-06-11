package beets

import (
	"database/sql"
)

type Artist struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	MBArtistID string `json:"mb_artist_id,omitempty"`
}

const artistColumns = `
	albumartist, mb_albumartistid
`

func (a *Artist) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&a.Name, &a.MBArtistID,
	)
}

func ParseRowsAsArtists(rows *sql.Rows) ([]Artist, error) {
	artists := make([]Artist, 0)
	for rows.Next() {
		artist := Artist{}

		if err := artist.Scan(rows); err != nil {
			return nil, err
		}
		artist.ID = artist.Name

		artists = append(artists, artist)
	}

	return artists, nil
}
