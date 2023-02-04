package model

type Tag struct {
	Tid int64  `json:"tag_id" db:"tid"`
	Tag string `json:"tag" db:"tag"`
}
