package sqlite

import (
	"context"
	"database/sql"
	"tgbot/lib/e"
	"tgbot/storage"

	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, e.Wrap("can't open database", err)
	}
	if err := db.Ping(); err != nil {
		return nil, e.Wrap("can't connect to database", err)

	}
	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, p *storage.Page) error {
	q := `INSERT INTO pages (url, user_name) VALUES (?, ?)`
	_, err := s.db.ExecContext(ctx, q, p.URL, p.UserName)
	if err != nil {
		return e.Wrap("can't insert page", err)
	}
	return nil
}

func (s *Storage) PickRandom(ctx context.Context, userName string) (*storage.Page, error) {
	q := `SELECT url FROM pages WHERE user_name = ? ORDER BY RANDOM() LIMIT 1`
	var url string
	err := s.db.QueryRowContext(ctx, q, userName).Scan(&url)
	if err == sql.ErrNoRows {
		return nil, storage.ErrNoSavedPages
	}
	if err != nil {
		return nil, e.Wrap("can't pick random page", err)
	}

	return &storage.Page{
		URL:      url,
		UserName: userName}, nil

}

func (s *Storage) Remove(ctx context.Context, page *storage.Page) error {
	q := `DELETE FROM pages WHERE url = ? AND user_name = ?`
	if _, err := s.db.ExecContext(ctx, q, page.URL, page.UserName); err != nil {
		return e.Wrap("can't remove page", err)
	}
	return nil
}

func (s *Storage) IsExists(ctx context.Context, page *storage.Page) (bool, error) {
	q := `SELECT COUNT(*) FROM pages WHERE url = ? AND user_name = ?`
	var count int

	if err := s.db.QueryRowContext(ctx, q, page.URL, page.UserName).Scan(&count); err != nil {
		return false, e.Wrap("can't check if page exists s", err)
	}

	return count > 0, nil
}

func (s *Storage) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`
	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return e.Wrap("can't init table", err)
	}
	return nil
}
