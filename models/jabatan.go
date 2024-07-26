package models

type Jabatan struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	Nama_jabatan string `gorm:"type:varchar(300)" json:"nama_jabatan"`
}
