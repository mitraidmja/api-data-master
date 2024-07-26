package models

type KaryawanKeluar struct {
	Id              int64     `gorm:"primaryKey" json:"id"`
	NomorSurat		string 	  `gorm:"type:varchar(100)" json:"nomor_surat"`
	KaryawanId      int64     `json:"karyawan_id"`
	Karyawan        *Karyawan `json:"karyawan"`
	UnitId			int64		`gorm:"type:int" json:"unit_id"`
	TglKeluar   	string    `gorm:"type:date" json:"tgl_keluar"`
	JenisKeluar		int64 	  `gorm:"type:int" json:"jenis_keluar"`
	Keterangan		string	  `gorm:"type:varchar(200)" json:"keterangan"`

}