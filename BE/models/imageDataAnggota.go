package models

import "time"

type ImageDataMahasiswa struct {
	ID         uint          `json:"id" gorm:"primaryKey;column:id_image_user"`
	DataUserID uint          `json:"-" gorm:"type:bigint;not null"`
	DataUser   DataMahasiswa `json:"-" gorm:"foreignKey:DataUserID"`
	ImageURL   string        `json:"imageUrl" gorm:"type:varchar(255);not null"`
	Keterangan string        `json:"keterangan" gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time     `json:"createdAt" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt  time.Time     `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp"`
}
