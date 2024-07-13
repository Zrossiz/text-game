package models

type Session struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:30; unique; not null" json:"name"`
	LocationId uint
	Location   Location `gorm:"foreignKey:LocationId;references:ID"`
}
