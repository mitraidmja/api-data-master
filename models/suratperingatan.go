package models

type SuratPeringatan struct {
	Id              int64     `gorm:"primaryKey" json:"id"`
	NomorSurat		string	  `gorm:"type:varchar(100)" json:"nomor_surat"`
	KaryawanId      int64     `json:"karyawan_id"`
	Karyawan        *Karyawan `json:"karyawan"`
	UnitId			int64		`gorm:"type:int" json:"unit_id"`
	TglPeringatan   string    `gorm:"type:date" json:"tgl_peringatan"`
	StatusPeringtan int64     `gorm:"type:int" json:"status_peringatan"`
	Keterangan      string    `gorm:"type:varchar(300)" json:"keterangan"`
}
