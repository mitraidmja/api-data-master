package models

type Mou struct {
	Id         int64  `gorm: "primaryKey" json: "id"`
	UnitId     int64  `json:"unit_id"`
	Unit       *Unit  `json: "unit"`
	Keterangan string `gorm:"type:varchar(100)" json:"keterangan"`
	Nomer      string `gorm:"type:varchar(100)" json:"nomer"`
	TglMulai   string `gorm:"type:date" json:"tgl_mulai"`
	TglBerakir string `gorm:"type:date" json:"tgl_berakir"`
	Status     int64  `json:"status_terdaftar"`
}
