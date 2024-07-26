package models

type Kontrak struct {
	Id            int64     `gorm: "primaryKey" json: "id"`
	KaryawanId    int64     `json:"id_karyawan"`
	Karyawan      *Karyawan `json: "karyawan"`
	UnitId			int64		`gorm:"type:int" json:"unit_id"`
	TglMulai      string    `gorm:"type:date" json:"tgl_mulai"`
	TglBerakir    string    `gorm:"type:date" json:"tgl_berakir"`
	No_kontrak    string    `gorm:"type:varchar(300)" json:"no_kontrak"`
	StatusKontrak int64     `gorm:"type:int" json:"status_kontrak"`
}
