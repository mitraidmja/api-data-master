package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "sqlserver://mitraid:TerujicobA0409@103.245.39.238:1433?database=mitraid"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&Jadwal{})
	// db.AutoMigrate(&Mou{})
	// db.AutoMigrate(&Invoice{})
	// db.AutoMigrate(&Status{})
	// db.AutoMigrate(&SetInvc{})
	// db.AutoMigrate(&Mesin{})
	// db.AutoMigrate(&TransferKaryawan{})
	// db.AutoMigrate(&Kecelakaan{})
	// db.AutoMigrate(&Absensi{})
	// db.AutoMigrate(&KaryawanKeluar{})
	// db.AutoMigrate(&Jabatan{})
	// db.AutoMigrate(&UnitStaff{})
	// db.AutoMigrate(&Bagian{})
	// db.AutoMigrate(&Kontrak{})
	// db.AutoMigrate(&Bpjs{})
	// db.AutoMigrate(&SuratPeringatan{})
	// db.AutoMigrate(&Staff{})
	DB = db
}
