package models

type Location struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:30; unique; not null" json:"name"`
}
