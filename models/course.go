package models

import (
	"time"
)

type Course struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Slug        string    `gorm:"unique;not null" json:"slug"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	Thumbnail   *string   `json:"thumbnail,omitempty"`
	YoutubeUrl  *string   `gorm:"column:youtube_url" json:"youtube_url,omitempty"`
	CodeUrl     *string   `gorm:"column:code_url" json:"code_url,omitempty"`
	IsPublished bool      `gorm:"column:is_published;default:false" json:"is_published"`
	CreatedAt   time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
