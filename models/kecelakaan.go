package models

type Kecelakaan struct {
	Id            int64     `gorm:"primaryKey" json:"id"`
	KaryawanId    int64     `json:"karyawan_id"`
	Karyawan      *Karyawan `json:"karyawan"`
	UnitId        int64     `gorm:"type:int" json:"unit_id"`
	TglKecelakaan string    `gorm:"type:date" json:"tgl_kecelakaan"`
	Tidakan       string    `gorm:"type:varchar(300)" json:"tindakan"`
	Keterangan    string    `gorm:"type:varchar(300)" json:"keterangan"`
}
