package storage

import (
	"crypto/sha1"
	"fmt"
	"io"
	"tgbot/lib/e"
)

type Storage interface {
	Save(p *Page)
	PickRandom(userName string) (*Page, error)
	Remove() error
	IsExists() (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't caclculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't caclculate hash", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil

}
