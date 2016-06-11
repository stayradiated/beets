package beets

import (
	"database/sql"
	"sort"
	"strings"

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

	items, err := ParseRowsAsItems(rows)
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		items[i].Attributes, err = b.GetItemAttributes(item.ID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
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

func (b *Beets) FilterAlbums(params map[string]string) ([]Album, error) {
	query := "select " + albumColumns + " from albums where "
	values := make([]interface{}, 0)

	for key, value := range params {
		query += key + " = ? "
		values = append(values, value)
	}

	rows, err := b.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	return ParseRowsAsAlbums(rows)
}

func (b *Beets) FilterItems(params map[string]string) ([]Item, error) {
	queries := make([]string, 0)
	values := make([]interface{}, 0)

	for key, value := range params {
		queries = append(queries, key+" = ?")
		values = append(values, value)
	}

	query := "select " + itemColumns + " from items where "
	query += strings.Join(queries, " and ")

	rows, err := b.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	items, err := ParseRowsAsItems(rows)
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		items[i].Attributes, err = b.GetItemAttributes(item.ID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (b *Beets) FilterArtists(params map[string]string) ([]Artist, error) {
	query := "select " + artistColumns + " from albums where "
	values := make([]interface{}, 0)

	for key, value := range params {
		query += key + " = ? "
		values = append(values, value)
	}

	query += " group by albumartist"

	rows, err := b.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	return ParseRowsAsArtists(rows)
}

func (b *Beets) SearchItems(query string) ([]Item, error) {
	query = "%" + query + "%"
	rows, err := b.db.Query(`
		select `+itemColumns+` from items where
			title LIKE ? OR artist LIKE ? OR album LIKE ?
	`, query, query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items, err := ParseRowsAsItems(rows)
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		items[i].Attributes, err = b.GetItemAttributes(item.ID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (b *Beets) GetAllItems() ([]Item, error) {
	rows, err := b.db.Query(`
		select ` + itemColumns + ` from items
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items, err := ParseRowsAsItems(rows)
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		items[i].Attributes, err = b.GetItemAttributes(item.ID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (b *Beets) GetAllAlbums() ([]Album, error) {
	rows, err := b.db.Query(`
		select ` + albumColumns + ` from albums order by added desc
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsAlbums(rows)
}

func (b *Beets) GetAllArtists() ([]Artist, error) {
	q := `
		select ` + artistColumns + `
		from albums
		group by albumartist
		order by albumartist
	`
	rows, err := b.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsArtists(rows)
}

func (b *Beets) GetItemsInAlbum(albumID int) ([]Item, error) {
	rows, err := b.db.Query(`
		select `+itemColumns+` from items where album_id = ?
	`, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items, err := ParseRowsAsItems(rows)
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		items[i].Attributes, err = b.GetItemAttributes(item.ID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (b *Beets) GetItem(itemID int) (Item, error) {
	item := Item{}

	rows, err := b.db.Query(`
		select `+itemColumns+` from items where id = ?
	`, itemID)
	if err != nil {
		return item, err
	}
	defer rows.Close()

	rows.Next()

	if err = item.Scan(rows); err != nil {
		return item, err
	}

	return item, nil
}

func (b *Beets) GetItemAttributes(itemID int) (AttributeList, error) {
	query := `
		select ` + attributeColumns + ` from item_attributes where entity_id = ?
	`

	rows, err := b.db.Query(query, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ParseRowsAsAttributes(rows)
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

func (b *Beets) GetAlbumWithItems(albumID int) (Album, error) {
	album, err := b.GetAlbum(albumID)
	if err != nil {
		return album, err
	}

	album.Items, err = b.GetItemsInAlbum(albumID)
	if err != nil {
		return album, err
	}

	sort.Sort(ByTrackNumber(album.Items))

	return album, nil
}

func (b *Beets) GetAlbumArtPath(albumID int) (string, error) {
	rows, err := b.db.Query(`
		select artpath from albums where id = ?
	`, albumID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	rows.Next()

	var artpath sql.NullString
	if err := rows.Scan(&artpath); err != nil {
		return "", err
	}

	return artpath.String, nil
}

func (b *Beets) GetItemPath(itemID int) (string, error) {
	rows, err := b.db.Query(`
		select path from items where id = ?
	`, itemID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	rows.Next()

	var path sql.NullString
	if err := rows.Scan(&path); err != nil {
		return "", err
	}

	return path.String, nil
}
