package models

type Bpjs struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	KaryawanId int64     `json:"karyawan_id"`
	Karyawan   *Karyawan `json:karyawan`
	UnitId			int64		`gorm:"type:int" json:"unit_id"`
	NoBpjsKes  string    `gorm:"type:varchar(300)" json:"no_bpjs_kes"`
	NoBpjsKet  string    `gorm:"type:varchar(300)" json:"no_bpjs_ket"`
	Paket      int64     `gorm:"type:int" json:"paket"`
}
