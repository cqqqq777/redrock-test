package model

import "time"

type Book struct {
	BookId      int64     `json:"book_id"`
	Score       int64     `json:"score"`
	Name        string    `json:"name"`
	Author      string    `json:"author,omitempty"`
	Cover       string    `json:"cover,omitempty"`
	Link        string    `json:"link,omitempty"`
	PublishTime time.Time `json:"publish_time,omitempty"`
}

type ApiBook struct {
	Started    bool   `json:"is_star"`
	Collected  bool   `json:"is_collected"`
	CommentNum int64  `json:"comment_num"`
	Tags       string `json:"tags"`
	Book       *Book  `json:"book"`
}
