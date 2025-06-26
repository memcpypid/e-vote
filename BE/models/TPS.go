package models

type TPS struct {
	ID        uint        `json:"id_tps" gorm:"primaryKey;column:id_tps"`
	NamaTPS   string      `json:"nama_tps" gorm:"type:varchar(100);not null"`
	NoTPS     string      `json:"no_tps" gorm:"type:varchar(100);not null"`
	DaerahID  *uint       `json:"daerah_id"`
	WilayahID *uint       `json:"wilayah_id"`
	Daerah    *DaerahTPS  `json:"daerah" gorm:"foreignKey:DaerahID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Wilayah   *WilayahTPS `json:"wilayah" gorm:"foreignKey:WilayahID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	IsOpen    bool        `json:"is_open" gorm:"type:boolean;default:false"`
}
