package models

type GroupUser struct{
	Id     		int64  `gorm:"primaryKey" json:"id"`
	NamaGroup 	string `gorm:"type:varchar(300)" json:"nama_group"`
	
}