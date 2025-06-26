package models

import "time"

type DataMahasiswa struct {
	ID     uint           `json:"id_data_mahasiswa" gorm:"primaryKey;column:id_data_mahasiswa"`
	UserID uint           `json:"userId" gorm:"type:bigint;not null"`
	User   *AkunMahasiswa `json:"user" gorm:"foreignKey:UserID"`
	TPSID  *uint          `json:"tps_id,omitempty" gorm:"type:bigint"`
	TPS    *TPS           `json:"tps,omitempty" gorm:"foreignKey:TPSID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	NamaDepan        string    `json:"nama_depan" gorm:"type:varchar(255);not null"`
	NamaBelakang     string    `json:"nama_belakang" gorm:"type:varchar(255)"`
	NoHp             string    `json:"no_hp" gorm:"type:varchar(255);unique;not null"`
	Alamat           string    `json:"alamat" gorm:"type:varchar(255);not null"`
	TanggalLahir     time.Time `json:"tanggalLahir" gorm:"type:date;not null"`
	NIK              string    `json:"nik" gorm:"type:varchar(20);unique;not null"`
	TempatLahir      string    `json:"tempatLahir" gorm:"type:varchar(100);not null"`
	Pekerjaan        string    `json:"pekerjaan" gorm:"type:varchar(100);not null"`
	StatusPerkawinan string    `json:"statusPerkawinan" gorm:"type:varchar(50);not null"`
	Agama            string    `json:"agama" gorm:"type:varchar(50);not null"`
	CreatedAt        time.Time `json:"createdAt" gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdatedAt        time.Time `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp;not null"`
	SudahMemilih     bool      `json:"sudah_memilih" gorm:"type:boolean;default:false;not null"`

	KandidatID *uint     `json:"kandidat_id,omitempty" gorm:"type:bigint"`
	Kandidat   *Kandidat `json:"kandidat,omitempty" gorm:"foreignKey:KandidatID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	ImageUsers []ImageDataMahasiswa `json:"imageUsers" gorm:"foreignKey:DataUserID"`
}
