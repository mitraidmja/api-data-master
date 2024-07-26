package models

type Status struct {
	Id        int64    `gorm:"primaryKey" json:"id"`
	InvoiceId int64    `json:"invoice_id"`
	Invoice   *Invoice `json:"invoice"`
	Memo      string   `gorm:"type:varchar(300)" json:"memo"`
	Tanggal   string   `gorm:"type:date" json:"tanggal"`
	St        int64    `gorm:"type:int" json:"st"`
}
