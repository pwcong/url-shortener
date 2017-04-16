package model

import "time"

type Url struct {
	ID        int64     `gorm:"AUTO_INCREMENT;primary_key"`
	Source    string    `gorm:"size:255;not null;unique"`
	CreatedAt time.Time `gorm:"not null"`
}
