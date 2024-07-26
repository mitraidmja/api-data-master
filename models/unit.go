package models

type Unit struct {
	Id           int64       `gorm:"primaryKey" json:"id"`
	Unit         string      `gorm:"type:varchar(300)" json:"unit"`
	Mou          *Mou        `json: "mou"`
	Karyawan     []*Karyawan `json: "karyawan"`
	UnitStaff    *UnitStaff  `json:"unit_staff"`
	Alamat       string      `gorm:"type:varchar(300)" json:"alamat"`
	Notelp       string      `gorm:"type:varchar(300)" json:"no_telp"`
	Nokontrak    string      `gorm:"type:varchar(300)" json:"no_kontrak"`
	TglBergabung string      `gorm:"type:date" json:"tgl_bergabung"`
	StatusUnit   int64       `gorm:"type:int" json:"status_unit"`
	NamaOwner    string      `gorm:"type:varchar(300)" json:"nama_owner"`
	BidangUsaha  string      `gorm:"type:varchar(300)" json:"bidang_usaha"`
	Groups       int64       `gorm:"type:int" json:"group"`
}
