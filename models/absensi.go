package models

type Absensi struct {
	Id             	int64    	`gorm:"primaryKey" json:"id"`
	KaryawanId      int64     	`json:"karyawan_id"`
	Karyawan        *Karyawan 	`json:"karyawan"`
	UnitId			int64		`gorm:"type:int" json:"unit_id"`
	Tgl				string		`gorm:"type:date" json:"tgl"`
	JamIn			string		`gorm:"type:datetime" json:"jam_in"`
	JamOut			string		`gorm:"type:datetime" json:"jam_out"`
	StIn			int64		`gorm:"type:int" json:"st_in"`
	StOut			int64		`gorm:"type:int" json:"st_out"`
}