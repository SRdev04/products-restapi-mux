package models

import "time"

//kita akan membuat model dari tabel products
type Products struct {
	Id        int64     `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Price     int64     `gorm:"type:bigint;not null" json:"price"`
	Quantity  int64     `gorm:"type:integer;default:0" json:"quantity"`
	Deskripsi string    `gorm:"type:text" json:"deskripsi"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}
