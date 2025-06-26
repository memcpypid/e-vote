package models

import "time"

type Kandidat struct {
	ID           uint   `json:"id_kandidat" gorm:"primaryKey;column:id_kandidat"`
	NamaKandidat string `json:"nama_kandidat" gorm:"type:varchar(100);not null"`
	NIMKandidat  string `json:"nim_kandidat" gorm:"type:varchar(50);not null;unique"`

	NamaPasangan *string `json:"nama_pasangan,omitempty" gorm:"type:varchar(100)"`
	NIMPasangan  *string `json:"nim_pasangan,omitempty" gorm:"type:varchar(50)"`

	Visi      string    `json:"visi" gorm:"type:text;not null"`
	Misi      string    `json:"misi" gorm:"type:text;not null"`
	Foto      *string   `json:"foto,omitempty" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp"`
}
