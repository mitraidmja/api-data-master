package models

type SetInvc struct {
	Id  int64 `gorm:"primaryKey" json:"id"`
	Ppn int64 `gorm:"type:int" json:"ppn"`
	Pph int64 `gorm:"type:int" json:"pph"`
}
