package models

type Command struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:30; unique; not null" json:"name"`
}
