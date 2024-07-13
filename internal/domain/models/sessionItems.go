package models

type SessionItems struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:100; unique; not null" json:"name"`
	LocationId uint
	Location   Location `gorm:"foreignKey:LocationId;references:ID"`
}
