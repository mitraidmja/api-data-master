package models

type Cmdlog struct {
	Id     int64  `gorm:"primaryKey" json:"id"`
	Cmd    string `gorm:"type:varchar(300)" json:"cmd"`
	UnitId int64  `json:"unit_id"`
	SN     string `gorm:"type:varchar(50)" json:"sn"`
	Rv     int64  `gorm:"type:int" json:"rv"`
}
