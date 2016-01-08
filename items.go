package beets

import (
	"database/sql"
)

const itemColumns = `
		added, album, albumartist, albumartist_sort, album_id, albumtype, artist,
		artist_sort, bitrate, day, disc, disctotal, id, label, length, month, path,
		original_day, original_month, original_year, title, track, tracktotal, year
`

type Item struct {
	ASIN                string  `json:",omitempty"`
	AcoustID            string  `json:",omitempty"`
	AcoustIDFingerprint string  `json:",omitempty"`
	Added               float64 `json:",omitempty"`
	Album               string  `json:",omitempty"`
	AlbumArtist         string  `json:",omitempty"`
	AlbumArtistCredit   string  `json:",omitempty"`
	AlbumArtistSort     string  `json:",omitempty"`
	AlbumDisambig       string  `json:",omitempty"`
	AlbumID             int     `json:",omitempty"`
	AlbumStatus         string  `json:",omitempty"`
	AlbumType           string  `json:",omitempty"`
	Artist              string  `json:",omitempty"`
	ArtistCredit        string  `json:",omitempty"`
	ArtistSort          string  `json:",omitempty"`
	BPM                 int     `json:",omitempty"`
	BitDepth            int     `json:",omitempty"`
	BitRate             int     `json:",omitempty"`
	CatalogNum          string  `json:",omitempty"`
	Channels            int     `json:",omitempty"`
	Comments            string  `json:",omitempty"`
	Comp                int     `json:",omitempty"`
	Composer            string  `json:",omitempty"`
	Country             string  `json:",omitempty"`
	Day                 int     `json:",omitempty"`
	Disc                int     `json:",omitempty"`
	DiscTitle           string  `json:",omitempty"`
	DiscTotal           int     `json:",omitempty"`
	Encoder             string  `json:",omitempty"`
	Format              string  `json:",omitempty"`
	Genre               string  `json:",omitempty"`
	Grouping            string  `json:",omitempty"`
	ID                  int     `json:",omitempty"`
	InitialKey          string  `json:",omitempty"`
	Label               string  `json:",omitempty"`
	Language            string  `json:",omitempty"`
	Length              float64 `json:",omitempty"`
	Lyrics              string  `json:",omitempty"`
	MBAlbumArtistID     string  `json:",omitempty"`
	MBAlbumID           string  `json:",omitempty"`
	MBArtistID          string  `json:",omitempty"`
	MBReleaseGroupID    string  `json:",omitempty"`
	MBTrackID           string  `json:",omitempty"`
	MTime               float64 `json:",omitempty"`
	Media               string  `json:",omitempty"`
	Month               int     `json:",omitempty"`
	OriginalDay         int     `json:",omitempty"`
	OriginalMonth       int     `json:",omitempty"`
	OriginalYear        int     `json:",omitempty"`
	Path                string  `json:",omitempty"`
	RGAlbumGain         float64 `json:",omitempty"`
	RGAlbumPeak         float64 `json:",omitempty"`
	RGTrackGain         float64 `json:",omitempty"`
	RGTrackPeak         float64 `json:",omitempty"`
	SampleRate          int     `json:",omitempty"`
	Script              string  `json:",omitempty"`
	Title               string  `json:",omitempty"`
	Track               int     `json:",omitempty"`
	TrackTotal          int     `json:",omitempty"`
	Year                int     `json:",omitempty"`
}

func (i *Item) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&i.Added, &i.Album, &i.AlbumArtist, &i.AlbumArtistSort, &i.AlbumID,
		&i.AlbumType, &i.Artist, &i.ArtistSort, &i.BitRate, &i.Day, &i.Disc,
		&i.DiscTotal, &i.ID, &i.Label, &i.Length, &i.Month, &i.Path,
		&i.OriginalDay, &i.OriginalMonth, &i.OriginalYear, &i.Title, &i.Track,
		&i.TrackTotal, &i.Year,
	)
}

func ParseRowsAsItems(rows *sql.Rows) ([]Item, error) {
	items := make([]Item, 0)
	for rows.Next() {
		item := Item{}

		if err := item.Scan(rows); err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return items, nil
}
