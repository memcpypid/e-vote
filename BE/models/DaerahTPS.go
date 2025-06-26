// models/daerah.go
package models

type DaerahTPS struct {
	ID         uint        `json:"id_daerah" gorm:"primaryKey;column:id_daerah"`
	NamaDaerah string      `json:"nama_daerah_tps" gorm:"type:varchar(50);not null"`
	WilayahID  uint        `json:"wilayah_id" gorm:"type:bigint;not null"`
	Wilayah    *WilayahTPS `json:"wilayah,omitempty" gorm:"foreignKey:WilayahID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
