package models

type Invoice struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	NoInv     string `gorm:"type:varchar(100)" json:"no_invc"`
	NoFak     string `gorm:"type:varchar(100)" json:"no_fak"`
	UnitId    int64  `json:"unit_id"`
	Unit      *Unit  `json:"unit"`
	Tanggal   string `gorm:"type:date" json:"tanggal"`
	TotalUpah int64  `gorm:"type:int" json:"total_upah"`
	Fee       int64  `gorm:"type:int" json:"fee"`
	Ppn       int64  `gorm:"type:int" json:"ppn"`
	Pph       int64  `gorm:"type:int" json:"pph"`
	Total     int64  `gorm:"type:int" json:"total"`
	TLunas    string `gorm:"type:date" json:"tgl_lunas"`
	St        int64  `gorm:"type:int" json:"st"`
}
