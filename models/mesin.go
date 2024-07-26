package models

type Mesin struct {
	Id           int64  `gorm: "primaryKey" json: "id"`
	NamaMesin    string `json:"nama_mesin"`
	SerialNumber string `json:"sn"`
	UnitId       int64  `json:"unit_id"`
	Unit         *Unit  `json:"unit"`
	CountKar     string `json:"count_kar"`
	CountClock   string `json:"count_clock`
}
