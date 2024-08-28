package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id         uuid.UUID `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uuid.UUID `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (post *Post) Prepare() error {

	if err := post.validate(); err != nil {
		return err
	}
	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title is mandatory and cannot be left blank")
	}
	if post.Content == "" {
		return errors.New("the content is mandatory and cannot be left blank")
	}
	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
