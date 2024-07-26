package models

type TransferKaryawan struct {
	Id             	int64    	`gorm:"primaryKey" json:"id"`
	KaryawanId      int64     	`json:"karyawan_id"`
	Karyawan        *Karyawan 	`json:"karyawan"`
	UnitId			int64		`gorm:"type:int" json:"unit_id"`
	NomorSurat		string	  	`gorm:"type:varchar(100)" json:"nomor_surat"`
	Unit1			string		`gorm:"type:varchar(100)" json:"unit1"`
	Unit2			string		`gorm:"type:varchar(100)" json:"unit2"`
	Bagian1			string		`gorm:"type:varchar(100)" json:"bagian1"`
	Bagian2			string		`gorm:"type:varchar(100)" json:"bagian2"`
	Jabatan1		string		`gorm:"type:varchar(100)" json:"jabatan1"`
	Jabatan2		string		`gorm:"type:varchar(100)" json:"jabatan2"`
	JenisTransfer	int64		`gorm:"type:int" json:"jenis_transfer"`
}