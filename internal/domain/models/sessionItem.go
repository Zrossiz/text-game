package models

type SessionItems struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:100; unique; not null" json:"name"`
	SessionId uint
	Session   Session `gorm:"foreignKey:SessionId;references:ID"`
}
