package models

type UnitStaff struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	UnitId  int64  `json:"unit_id"`
	Unit    *Unit  `json:"unit"`
	StaffId int64  `json:"staff_id"`
	Staff   *Staff `json:"staff"`
}
