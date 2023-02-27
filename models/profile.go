package models

import (
	"time"
)

type Profile struct {
	ID              string    `db:"id" json:"id"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
	Username        string    `db:"username" json:"username"`
	AvatarURL       string    `db:"avatar_url" json:"avatar_url"`
	Bio             string    `db:"bio" json:"bio"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	Name            string    `db:"name" json:"name"`
	WorldMapURL     string    `db:"worldmap_url" json:"worldmap_url"`
	NumberFollowers int64     `db:"number_followers" json:"number_followers"`
	NumberFollowing int64     `db:"number_following" json:"number_following"`
}
