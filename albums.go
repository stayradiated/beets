package beets

import (
	"database/sql"
)

type Album struct {
	ASIN              string          `json:"asin,omitempty"`
	Added             float64         `json:"added,omitempty"`
	Album             string          `json:"album,omitempty"`
	AlbumArtist       string          `json:"album_artist,omitempty"`
	AlbumArtistCredit string          `json:"album_artist_credit,omitempty"`
	AlbumArtistSort   string          `json:"album_artist_sort,omitempty"`
	AlbumDisambig     string          `json:"album_disambig,omitempty"`
	AlbumStatus       string          `json:"album_status,omitempty"`
	AlbumType         string          `json:"album_type,omitempty"`
	ArtPath           *JSONNullString `json:"art_path,omitempty"`
	CatalogNum        string          `json:"catalog_num,omitempty"`
	Comp              int             `json:"comp,omitempty"`
	Country           string          `json:"country,omitempty"`
	Day               int             `json:"day,omitempty"`
	DiscTotal         int             `json:"disc_total,omitempty"`
	Genre             string          `json:"genre,omitempty"`
	ID                int             `json:"id,omitempty"`
	Label             string          `json:"label,omitempty"`
	Language          string          `json:"language,omitempty"`
	MBAlbumArtistID   string          `json:"mb_album_artist_id,omitempty"`
	MBAlbumID         string          `json:"mb_album_id,omitempty"`
	MBReleaseGroupID  string          `json:"mb_release_group_id,omitempty"`
	Month             int             `json:"month,omitempty"`
	OriginalDay       int             `json:"original_day,omitempty"`
	OriginalMonth     int             `json:"original_month,omitempty"`
	OriginalYear      int             `json:"original_year,omitempty"`
	RGAlbumGain       float64         `json:"rg_album_gain,omitempty"`
	RGAlbumPeak       float64         `json:"rg_album_peak,omitempty"`
	Script            string          `json:"script,omitempty"`
	Year              int             `json:"year,omitempty"`
	Items             []Item          `json:"items,omitempty"`
}

const albumColumns = `
	added, album, albumartist, albumartist_sort, albumtype, day, genre, id,
	label, month, year
`

func (a *Album) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&a.Added, &a.Album, &a.AlbumArtist, &a.AlbumArtistSort, &a.AlbumType,
		&a.Day, &a.Genre, &a.ID, &a.Label, &a.Month, &a.Year,
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
