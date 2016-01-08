package beets

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Beets struct {
	LibraryPath string
	db          *sql.DB
}

func NewBeets(path string) *Beets {
	return &Beets{
		LibraryPath: path,
	}
}

func (b *Beets) Connect() error {
	db, err := sql.Open("sqlite3", b.LibraryPath)
	if err != nil {
		return err
	}
	b.db = db
	return nil
}

func (b *Beets) Close() error {
	return b.db.Close()
}

func (b *Beets) GetItemsByArtist(artist string) ([]Item, error) {
	rows, err := b.db.Query(`
		select `+itemColumns+` from items where artist = ?
	`, artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsItems(rows)
}

func (b *Beets) GetAlbumsByArtist(artist string) ([]Album, error) {
	rows, err := b.db.Query(`
		select `+albumColumns+` from albums where albumartist = ?
	`, artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsAlbums(rows)
}

func (b *Beets) SearchItems(query string) ([]Item, error) {
	rows, err := b.db.Query(`
		select `+itemColumns+` from items where
			title LIKE ? OR artist LIKE ? OR album LIKE ?
	`, query, query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsItems(rows)
}

func (b *Beets) GetAllItems() ([]Item, error) {
	rows, err := b.db.Query(`
		select ` + itemColumns + ` from items
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsItems(rows)
}

func (b *Beets) GetAllAlbums() ([]Album, error) {
	rows, err := b.db.Query(`
		select ` + albumColumns + ` from albums
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsAlbums(rows)
}

func (b *Beets) GetItemsInAlbum(albumID int) ([]Item, error) {
	rows, err := b.db.Query(`
		select `+itemColumns+` from items where album_id = ?
	`, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsItems(rows)
}

func (b *Beets) GetAlbum(albumID int) (Album, error) {
	album := Album{}

	rows, err := b.db.Query(`
		select `+albumColumns+` from albums where id = ?
	`, albumID)
	if err != nil {
		return album, err
	}
	defer rows.Close()

	rows.Next()

	if err = album.Scan(rows); err != nil {
		return album, err
	}

	return album, err
}
