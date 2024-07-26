package models

type Bagian struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Nama_bagian string `gorm:"type:varchar(300)" json:"nama_bagian"`
}
