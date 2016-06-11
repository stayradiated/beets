package beets

import (
	"database/sql"
)

type Item struct {
	ASIN                string        `json:"asin,omitempty"`
	AcoustID            string        `json:"acoustid,omitempty"`
	AcoustIDFingerprint string        `json:"acoustid_fingerprint,omitempty"`
	Added               float64       `json:"added,omitempty"`
	Album               string        `json:"album,omitempty"`
	AlbumArtist         string        `json:"album_artist,omitempty"`
	AlbumArtistCredit   string        `json:"album_artist_credit,omitempty"`
	AlbumArtistSort     string        `json:"album_artist_sort,omitempty"`
	AlbumDisambig       string        `json:"album_disambig,omitempty"`
	AlbumID             int           `json:"album_id,omitempty"`
	AlbumStatus         string        `json:"album_status,omitempty"`
	AlbumType           string        `json:"album_type,omitempty"`
	Artist              string        `json:"artist,omitempty"`
	ArtistCredit        string        `json:"artist_credit,omitempty"`
	ArtistSort          string        `json:"artist_sort,omitempty"`
	Attributes          AttributeList `json:"attributes,omitempty"`
	BPM                 int           `json:"bpm,omitempty"`
	BitDepth            int           `json:"bitdepth,omitempty"`
	BitRate             int           `json:"bitrate,omitempty"`
	CatalogNum          string        `json:"catalog_num,omitempty"`
	Channels            int           `json:"channels,omitempty"`
	Comments            string        `json:"comments,omitempty"`
	Comp                int           `json:"comp,omitempty"`
	Composer            string        `json:"composer,omitempty"`
	Country             string        `json:"country,omitempty"`
	Day                 int           `json:"day,omitempty"`
	Disc                int           `json:"disc,omitempty"`
	DiscTitle           string        `json:"disc_title,omitempty"`
	DiscTotal           int           `json:"disc_total,omitempty"`
	Encoder             string        `json:"encoder,omitempty"`
	Format              string        `json:"format,omitempty"`
	Genre               string        `json:"genre,omitempty"`
	Grouping            string        `json:"grouping,omitempty"`
	ID                  int           `json:"id,omitempty"`
	InitialKey          string        `json:"initial_key,omitempty"`
	Label               string        `json:"label,omitempty"`
	Language            string        `json:"language,omitempty"`
	Length              float64       `json:"length,omitempty"`
	Lyrics              string        `json:"lyrics,omitempty"`
	MBAlbumArtistID     string        `json:"mb_album_artist_id,omitempty"`
	MBAlbumID           string        `json:"mb_album_id,omitempty"`
	MBArtistID          string        `json:"mb_artist_id,omitempty"`
	MBReleaseGroupID    string        `json:"mb_release_group_id,omitempty"`
	MBTrackID           string        `json:"mb_track_id,omitempty"`
	MTime               float64       `json:"mtime,omitempty"`
	Media               string        `json:"media,omitempty"`
	Month               int           `json:"month,omitempty"`
	OriginalDay         int           `json:"original_day,omitempty"`
	OriginalMonth       int           `json:"original_month,omitempty"`
	OriginalYear        int           `json:"original_year,omitempty"`
	Path                string        `json:"path,omitempty"`
	RGAlbumGain         float64       `json:"rg_album_gain,omitempty"`
	RGAlbumPeak         float64       `json:"rg_album_peak,omitempty"`
	RGTrackGain         float64       `json:"rg_track_gain,omitempty"`
	RGTrackPeak         float64       `json:"rg_track_peak,omitempty"`
	SampleRate          int           `json:"sample_rate,omitempty"`
	Script              string        `json:"script,omitempty"`
	Title               string        `json:"title,omitempty"`
	Track               int           `json:"track,omitempty"`
	TrackTotal          int           `json:"track_total,omitempty"`
	Year                int           `json:"year,omitempty"`
}

const itemColumns = `
		added, album, albumartist, albumartist_sort, album_id, albumtype, artist,
		artist_sort, bitrate, day, disc, disctotal, genre, id, label, length,
		month, title, track, tracktotal, year
`

func (i *Item) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&i.Added, &i.Album, &i.AlbumArtist, &i.AlbumArtistSort, &i.AlbumID,
		&i.AlbumType, &i.Artist, &i.ArtistSort, &i.BitRate, &i.Day, &i.Disc,
		&i.DiscTotal, &i.Genre, &i.ID, &i.Label, &i.Length, &i.Month, &i.Title,
		&i.Track, &i.TrackTotal, &i.Year,
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
