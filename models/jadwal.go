package models

type Jadwal struct {
	Id             int64  `gorm: "primaryKey" json: "id"`
	NamaJadwal     string `gorm:"type:varchar(100)" json:"nama_jadwal"`
	UnitId         int64  `json:"unit_id"`
	Unit           *Unit  `json:"unit"`
	Durasi         int64  `gorm:"type:int" json:"durasi"`
	MulaiJam       string `gorm:"type:varchar(50)" json:"mulai_jam"`
	JamMasuk       string `gorm:"type:varchar(50)" json:"jam_masuk"`
	AkhirJamMasuk  string `gorm:"type:varchar(50)" json:"akhir_jam_masuk"`
	MulaiKeluar    string `gorm:"type:varchar(50)" json:"mulai_keluar"`
	JamKeluar      string `gorm:"type:varchar(50)" json:"jam_keluar"`
	AkhirJamKeluar string `gorm:"type:varchar(50)" json:"akhir_keluar_jam"`
	Toleransi      int64  `gorm:"type:int" json:"toleransi"`
	IsKel          string `gorm:"type:varchar(50)" json:"is_kel"`
	IsMa           string `gorm:"type:varchar(50)" json:"is_ma"`
	MinLem         int64  `gorm:"type:int" json:"min_lem"`
	MaxLem         int64  `gorm:"type:int" json:"max_lem"`
}
