package beets

import (
	"database/sql"
)

const albumColumns = `
	added, album, albumartist, albumartist_sort, albumtype, artpath, day, id,
	label, month, original_day, original_month, original_year, year
`

type Album struct {
	ASIN              string         `json:",omitempty"`
	Added             float64        `json:",omitempty"`
	Album             string         `json:",omitempty"`
	AlbumArtist       string         `json:",omitempty"`
	AlbumArtistCredit string         `json:",omitempty"`
	AlbumArtistSort   string         `json:",omitempty"`
	AlbumDisambig     string         `json:",omitempty"`
	AlbumStatus       string         `json:",omitempty"`
	AlbumType         string         `json:",omitempty"`
	ArtPath           JSONNullString `json:",omitempty"`
	CatalogNum        string         `json:",omitempty"`
	Comp              int            `json:",omitempty"`
	Country           string         `json:",omitempty"`
	Day               int            `json:",omitempty"`
	DiscTotal         int            `json:",omitempty"`
	Genre             string         `json:",omitempty"`
	ID                int            `json:",omitempty"`
	Label             string         `json:",omitempty"`
	Language          string         `json:",omitempty"`
	MBAlbumArtistID   string         `json:",omitempty"`
	MBAlbumID         string         `json:",omitempty"`
	MBReleaseGroupID  string         `json:",omitempty"`
	Month             int            `json:",omitempty"`
	OriginalDay       int            `json:",omitempty"`
	OriginalMonth     int            `json:",omitempty"`
	OriginalYear      int            `json:",omitempty"`
	RGAlbumGain       float64        `json:",omitempty"`
	RGAlbumPeak       float64        `json:",omitempty"`
	Script            string         `json:",omitempty"`
	Year              int            `json:",omitempty"`
}

func (a *Album) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&a.Added, &a.Album, &a.AlbumArtist, &a.AlbumArtistSort, &a.AlbumType,
		&a.ArtPath, &a.Day, &a.ID, &a.Label, &a.Month, &a.OriginalDay,
		&a.OriginalMonth, &a.OriginalYear, &a.Year,
	)
}

func ParseRowsAsAlbums(rows *sql.Rows) ([]Album, error) {
	albums := make([]Album, 0)
	for rows.Next() {
		album := Album{}

		if err := album.Scan(rows); err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}
