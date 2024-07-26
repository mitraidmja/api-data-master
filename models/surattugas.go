package models

type SuratTugas struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	KaryawanId int64     `json:"karyawan_id"`
	Karyawan   *Karyawan `json:"karyawan"`
	UnitId     int64     `gorm:"type:int" json:"unit_id"`
}
