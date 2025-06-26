package models

type WilayahTPS struct {
	ID          uint        `json:"id_wilayah" gorm:"primaryKey;column:id_wilayah"`
	NamaWilayah string      `json:"nama_wilayah_tps" gorm:"type:varchar(50);not null;unique"`
	Daerah      []DaerahTPS `json:"-" gorm:"foreignKey:WilayahID"`
}
