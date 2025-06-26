package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AkunMahasiswa struct {
	ID        uint       `json:"id_user" gorm:"primaryKey;column:id_user"`
	NIM       int        `json:"nim" gorm:"type:int;unique;not null"`
	PIC       string     `json:"pic" gorm:"type:varchar(255);not null"`
	Role      string     `json:"role" gorm:"type:varchar(20);not null;default:mahasiswa"`
	LastLogin *time.Time `json:"lastLogin,omitempty"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp"`
}

// HashPassword meng-hash nilai PIC (password)
func (user *AkunMahasiswa) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PIC), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PIC = string(hashedPassword)
	return nil
}

// ComparePassword membandingkan password input dengan hash tersimpan
func (user *AkunMahasiswa) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PIC), []byte(password))
	return err == nil
}
