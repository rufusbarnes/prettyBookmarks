package mysql

import (
	"database/sql"

	"github.com/rufusbarnes/prettyBookmarks/pkg/models"
)

type BookmarksModel struct {
	DB *sql.DB
}

func (m *BookmarksModel) Get(id int) (*models.Bookmarks, error) {
	// stmt := `
	// SELECT id, title, content, created, expires
	// FROM snippets
	// WHERE expires > UTC_TIMESTAMP()
	// AND id = ?`

	s := &models.Bookmarks{}

	// err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	var err error = nil
	// If the query returns no row, scan returns a sql.ErrNoRows error
	if err == sql.ErrNoRows {
		// Return our own error to encapsulate our application from any specific DB driver
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
