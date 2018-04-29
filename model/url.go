package model

type Url struct {
	BaseModel
	SourceUrl string `gorm:"not null;"`
}
