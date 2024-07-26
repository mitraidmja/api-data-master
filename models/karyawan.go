package models

type Karyawan struct {
	Id             int64    `gorm:"primaryKey" json:"id"`
	KodeKaryawan   string   `gorm:"type:varchar(300)" json:"kode_karyawan"`
	Nama           string   `gorm:"type:varchar(300)" json:"nama"`
	Jk             string   `gorm:"type:varchar(300)" json:"jk"`
	Nik            string   `gorm:"type:varchar(300)" json:"nik"`
	Nkk            string   `gorm:"type:varchar(300)" json:"nkk"`
	Alamat         string   `gorm:"type:varchar(300)" json:"alamat"`
	Dsn            string   `gorm:"type:varchar(300)" json:"dsn"`
	Rt             string   `gorm:"type:varchar(300)" json:"rt"`
	Rw             string   `gorm:"type:varchar(300)" json:"rw"`
	Desa           string   `gorm:"type:varchar(300)" json:"desa"`
	Kecamatan      string   `gorm:"type:varchar(300)" json:"kecamatan"`
	Kota           string   `gorm:"type:varchar(300)" json:"kota"`
	Agama          string   `gorm:"type:varchar(300)" json:"agama"`
	Notelp         string   `gorm:"type:varchar(300)" json:"no_telp"`
	TglLahir       string   `gorm:"type:date" json:"tgl_lahir"`
	TempatLahir    string   `gorm:"type:varchar(300)" json:"tempat_lahir"`
	StatusPersonal int64    `gorm:"type:int" json:"status_personal"`
	NamaIbu        string   `gorm:"type:varchar(300)" json:"nama_ibu"`
	UnitId         int64    `json:"unit_id"`
	Unit           *Unit    `json:"unit"`
	JabatanId      int64    `json:"jabatan_id"`
	Jabatan        *Jabatan `json:"jabatan"`
	BagianId       int64    `json:"bagian_id"`
	Bagian         *Bagian  `json:"bagian"`
	Sn             string   `gorm:"type:varchar(300)" json:"sn"`
	IsAktif        int64    `gorm:"type:int" json:"is_aktife"`
	TglBergabung   string   `gorm:"type:date" json:"tgl_bergabung"`
}
