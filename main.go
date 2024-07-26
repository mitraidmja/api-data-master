package main

import (
	"github.com/adewputro/go-rest-api-mja/controllers/absensicontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/bagiancontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/bpjscontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/cmdcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/groupcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/invoicecontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/jabatancontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/jadwalcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/karyawancontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/kecelakaancontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/keluarcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/kontrakcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/logincontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/mesincontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/moucontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/staffcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/suratperingcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/unitcontroller"
	"github.com/adewputro/go-rest-api-mja/controllers/unitstaffcontroller"
	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true
	// port :=
	r.Use(cors.New(config))
	models.ConnectDatabase()
	//login
	r.POST("/login", logincontroller.Login)
	//bpjs
	r.GET("/api/bpjs", bpjscontroller.Index)
	r.GET("/api/bpjs/:id", bpjscontroller.Show)
	r.GET("/api/bpjs/unit/:id", bpjscontroller.ShowUnit)
	r.POST("/api/bpjs", bpjscontroller.Create)
	r.PUT("/api/bpjs/:id", bpjscontroller.Update)
	r.DELETE("/api/bpjs/:id", bpjscontroller.Delete)
	//kontrak
	r.GET("/api/kontrak", kontrakcontroller.Index)
	r.GET("/api/kontrak/:id", kontrakcontroller.Show)
	r.GET("/api/kontrak/karyawan/:id", kontrakcontroller.ShowKaryawan)
	r.GET("/api/kontrak/unit/:id", kontrakcontroller.ShowUnit)
	r.POST("/api/kontrak", kontrakcontroller.Create)
	r.PUT("/api/kontrak/:id", kontrakcontroller.Update)
	r.DELETE("/api/kontrak/:id", kontrakcontroller.Delete)
	//jabatan
	r.GET("/api/jabatan", jabatancontroller.Index)
	r.GET("/api/jabatan/:id", jabatancontroller.Show)
	r.POST("/api/jabatan", jabatancontroller.Create)
	r.PUT("/api/jabatan/:id", jabatancontroller.Update)
	r.DELETE("/api/jabatan/:id", jabatancontroller.Delete)
	//bagian
	r.GET("/api/bagian", bagiancontroller.Index)
	r.GET("/api/bagian/:id", bagiancontroller.Show)
	r.POST("/api/bagian", bagiancontroller.Create)
	r.PUT("/api/bagian/:id", bagiancontroller.Update)
	r.DELETE("/api/bagian/:id", bagiancontroller.Delete)
	//units
	r.POST("/api/unit", unitcontroller.Create)
	r.GET("/api/unit", unitcontroller.Index)
	r.GET("/api/unit/group/:id", unitcontroller.ShowGroup)
	r.GET("/api/unit/select", unitcontroller.SelectData)
	r.GET("/api/unit/:id", unitcontroller.Show)
	r.PUT("/api/unit/:id", unitcontroller.Update)
	r.DELETE("/api/unit/:id", unitcontroller.Delete)
	//karyawans
	r.GET("/api/karyawan/", karyawancontroller.Index)
	r.GET("/api/karyawan/unit/:id", karyawancontroller.ShowUnit)
	r.GET("/api/karyawan/:id", karyawancontroller.Show)
	r.POST("/api/karyawan/", karyawancontroller.Create)
	r.PUT("/api/karyawan/:id", karyawancontroller.Update)
	r.DELETE("/api/karyawan/:id", karyawancontroller.Delete)
	//absensi
	r.GET("/api/absensi", absensicontroller.Index)
	r.GET("/api/absensi/filter", absensicontroller.FilterDate)
	r.GET("/api/absensi/unit/:id", absensicontroller.ShowUnit)
	r.GET("/api/absensi/:id", absensicontroller.Show)
	r.POST("/api/absensi", absensicontroller.Create)
	r.PUT("/api/absensi/:id", absensicontroller.Update)
	r.DELETE("/api/absensi/:id", absensicontroller.Delete)
	//sp
	r.GET("/api/sp", suratperingcontroller.Index)
	r.GET("/api/sp/:id", suratperingcontroller.Show)
	r.GET("/api/sp/karyawan/:id", suratperingcontroller.ShowKaryawan)
	r.GET("/api/sp/unit/:id", suratperingcontroller.ShowUnit)
	r.POST("/api/sp", suratperingcontroller.Create)
	r.PUT("/api/sp/:id", suratperingcontroller.Update)
	r.DELETE("/api/sp/:id", suratperingcontroller.Delete)
	//keluar
	r.GET("/api/keluar", keluarcontroller.Index)
	r.GET("/api/keluar/:id", keluarcontroller.Show)
	r.GET("/api/keluar/karyawan/:id", keluarcontroller.ShowKaryawan)
	r.GET("/api/keluar/unit/:id", keluarcontroller.ShowUnit)
	r.POST("/api/keluar", keluarcontroller.Create)
	r.PUT("/api/keluar/:id", keluarcontroller.Update)
	r.DELETE("/api/keluar/:id", keluarcontroller.Delete)
	//staff
	r.GET("/api/staff/", staffcontroller.Index)
	r.GET("/api/staff/unit/:id", staffcontroller.ShowUnit)
	r.GET("/api/staff/:id", staffcontroller.Show)
	r.POST("/api/staff/", staffcontroller.Create)
	r.PUT("/api/staff/:id", staffcontroller.Update)
	r.DELETE("/api/staff/:id", staffcontroller.Delete)
	//mesin
	r.GET("/api/mesin", mesincontroller.Index)
	r.GET("/api/mesin/:id", mesincontroller.Show)
	r.GET("/api/mesin/unit/:id", mesincontroller.ShowUnit)
	r.POST("/api/mesin", mesincontroller.Create)
	r.PUT("/api/mesin/:id", mesincontroller.Update)
	r.DELETE("/api/mesin/:id", mesincontroller.Delete)
	//cmdlog
	r.POST("/api/cmd", cmdcontroller.Create)
	r.GET("/api/cmd", cmdcontroller.Index)
	//invoice
	r.GET("/api/invoice", invoicecontroller.Index)
	r.GET("/api/invoice/unit/:id", invoicecontroller.ShowUnit)
	r.GET("/api/invoice/status/:id", invoicecontroller.ShowStatus)
	r.POST("/api/invoice", invoicecontroller.Create)
	r.POST("/api/invoice/status", invoicecontroller.CreateStatus)
	r.DELETE("/api/invoice/:id", invoicecontroller.Delete)
	//groupuser
	r.GET("/api/group", groupcontroller.Index)
	r.GET("/api/group/:id", groupcontroller.Show)
	r.POST("/api/group", groupcontroller.Create)
	r.PUT("/api/group/:id", groupcontroller.Update)
	r.DELETE("/api/group/:id", groupcontroller.Delete)
	//Mou
	r.GET("/api/mou", moucontroller.Index)
	r.GET("/api/mou/:id", moucontroller.Show)
	r.POST("/api/mou", moucontroller.Create)
	r.PUT("/api/mou/:id", moucontroller.Update)
	r.DELETE("/api/mou/:id", moucontroller.Delete)
	//unitstaff
	r.GET("/api/unitstaff", unitstaffcontroller.Index)
	r.GET("/api/unitstaff/:id", unitstaffcontroller.Show)
	r.POST("/api/unitstaff", unitstaffcontroller.Create)
	r.DELETE("/api/unitstaff/:id", unitstaffcontroller.Delete)
	//jadwal
	r.GET("/api/jadwal", jadwalcontroller.Index)
	r.GET("/api/jadwal/:id", jadwalcontroller.Show)
	r.POST("/api/jadwal", jadwalcontroller.Create)
	r.PUT("/api/jadwal/:id", jadwalcontroller.Update)
	r.DELETE("/api/jadwal/:id", jadwalcontroller.Delete)
	//kecelakaan
	r.GET("/api/kecelakaan", kecelakaancontroller.Index)
	r.GET("/api/kecelakaan/:id", kecelakaancontroller.Show)
	r.POST("/api/kecelakaan", kecelakaancontroller.Create)
	r.PUT("/api/kecelakaan/:id", kecelakaancontroller.Update)
	r.DELETE("/api/kecelakaan/:id", kecelakaancontroller.Delete)
	//main
	r.Run(":54499")
}
