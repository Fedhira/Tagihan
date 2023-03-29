package namapackage

import (
	"fmt"
	"testing"

	model "github.com/Fedhira/Tagihan/model"
	"github.com/Fedhira/Tagihan/module"
)

//NASABAH
func TestInsertNasabah(t *testing.T) {
	nama_nasabah := "Fedhira Syaila"
	email:= "sfedhira@gmail.com"
	phone_number := "62812345128342"
	alamat := "Bekasi"
		
	hasil:= module.InsertNasabah(module.MongoConn, "nasabah", nama_nasabah, email, phone_number, alamat)
	fmt.Println(hasil)
}

func TestGetNasabahFromNama(t *testing.T) {
	nama_nasabah := "Fedhira Syaila"
	biodata:= module.GetNasabahFromNama(nama_nasabah, "nasabah", module.MongoConn)
	fmt.Println(biodata)
}

func TestGetAllTagihanFromNama_nasabah(t *testing.T) {
	nama_nasabah := "Fedhira Syaila"
	data1 := module.GetAllTagihanFromNama_nasabah(nama_nasabah, module.MongoConn, "tagihan")
	fmt.Println(data1)
}

//PENAGIH
func TestInsertPenagih(t *testing.T) {
	nama_penagih := "Marlina"
	email := "123@gmail.com"
	phone_number := "0863890390"
	total_tagihan := model.Tagihan{
		Total_Tagihan : "900000",
		Deskripsi : "Tagihan Listrik Kos",
		Status : "Lunas",
		Tanggal_jatuhtempo : "11 november 2020",
		Biodata : model.Nasabah{
			Nama_nasabah : "Fedhira Syaila",
			Email : "sfedhira@gmail.com",
			Phone_number : "62812345128342",
			Alamat : "Bekasi",
		},
		Location : "Bekasi",
		Longitude : 98.345345,
		Latitude : 123.561651,
	}
	
	hasil:= module.InsertPenagih(module.MongoConn, "penagih", nama_penagih, email, phone_number, total_tagihan)
	fmt.Println(hasil)
}

func TestGetPenagihFromNama(t *testing.T) {
	nama_penagih := "Marlina"
	biodata:= module.GetPenagihFromNama(nama_penagih,  "penagih", module.MongoConn)
	fmt.Println(biodata)
}

// func TestGetPenagihAll(t *testing.T) {
// 	biodata := module.GetPenagihAll()
// 	fmt.Println(biodata)
// }


//TAGIHAN
func TestInsertTagihan(t *testing.T) {
	total_tagihan := "900000"
	deskripsi := "Tagihan Listrik Kos"
	status := "Lunas"
	tanggal_jatuhtempo := "11 november 2020"
	biodata := model.Nasabah{
		Nama_nasabah : "Fedhira Syaila",
		Email : "sfedhira@gmail.com",
		Phone_number : "62812345128342",
		Alamat : "Bekasi",
		}
	location := "Bekasi"
	longitude := 98.345345
	latitude := 123.561651

	hasil:= module.InsertTagihan(module.MongoConn, "tagihan", total_tagihan, deskripsi, status, tanggal_jatuhtempo, biodata, location, longitude, latitude )
	fmt.Println(hasil)
}

func TestGetTagihanFromNama_nasabah(t *testing.T) {
	nama_nasabah := "Fedhira Syaila"
	biodata:= module.GetTagihanFromNama_nasabah( nama_nasabah, "tagihan", module.MongoConn)
	fmt.Println(biodata)
}

// func TestGetTagihanAll(t *testing.T) {
// 	biodata := module.GetTagihanAll()
// 	fmt.Println(biodata)
// }

func TestInsertBank(t *testing.T) {
	nama_bank := "bank abc"
	lokasi := "Bandung"
	total_tagihan:= []model.Tagihan{
		{
			Total_Tagihan : "900000",
			Deskripsi : "Pinjaman",
			Status : "Lunas",
			Tanggal_jatuhtempo : "11 november 2020",
			Biodata : model.Nasabah{
				Nama_nasabah : "Fedhira Syaila",
				Email : "sfedhira@gmail.com",
				Phone_number : "0812345128342",
				Alamat : "Bekasi",
		},
			Location : "Bekasi",
			Longitude : 98.345345,
			Latitude : 123.561651,
		}, {
			Total_Tagihan : "1000000",
			Deskripsi : "Kartu Kredit",
			Status : "Belum Lunas",
			Tanggal_jatuhtempo : "23 Maret 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Auliyah Safana",
				Email : "aol@gmail.com",
				Phone_number : "08946737892",
				Alamat : "Depok",
		},
		Location : "Depok",
		Longitude : 89.567899,
		Latitude : 124.781781,
		}, {
			Total_Tagihan : "2500000",
			Deskripsi : "Kredit Kendaraan Bermotor",
			Status : "Belum Lunas",
			Tanggal_jatuhtempo : "16 Juni 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Juwita",
				Email : "juwita@gmail.com",
				Phone_number : "08123478910",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 94.321321,
		Latitude : 234.678678,
		},
	}

	daftar := []model.Penagih{
		{
			Nama_penagih : "Marlina",
			Email : "123@gmail.com",
			Phone_number : "0863890390",
			Total_Tagihan : model.Tagihan{
				Total_Tagihan : "900000",
				Deskripsi : "Tagihan Listrik Kos",
				Status : "Lunas",
				Tanggal_jatuhtempo : "11 november 2020",
				Biodata : model.Nasabah{
					Nama_nasabah : "Fedhira Syaila",
					Email : "sfedhira@gmail.com",
					Phone_number : "62812345128342",
					Alamat : "Bekasi",
				},
				Location : "Bekasi",
				Longitude : 98.345345,
				Latitude : 123.561651,
			},
	}, {
		Nama_penagih : "Rizkyria",
		Email : "12@gmail.com",
		Phone_number : "0813456789",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "1000000",
			Deskripsi : "Kartu Kredit",
			Status : "Belum Lunas",
			Tanggal_jatuhtempo : "23 Maret 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Auliyah Safana",
				Email : "aol@gmail.com",
				Phone_number : "08946737892",
				Alamat : "Depok",
		},
		Location : "Depok",
		Longitude : 89.567899,
		Latitude : 124.781781,
		},
			
	}, {
		Nama_penagih : "Agita",
		Email : "yellow@gmail.com",
		Phone_number : "08224568990",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "2500000",
			Deskripsi : "Kredit Kendaraan Bermotor",
			Status : "Belum Lunas",
			Tanggal_jatuhtempo : "16 Juni 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Juwita",
				Email : "juwita@gmail.com",
				Phone_number : "08123478910",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 94.321321,
		Latitude : 234.678678,
	},

},
	}

	biodata := []model.Nasabah{
		{
			Nama_nasabah : "Fedhira Syaila",
			Email : "sfedhira@gmail.com",
			Phone_number : "0812345128342",
			Alamat : "Bekasi",
		}, {
			Nama_nasabah : "Auliyah Safana",
			Email : "aol@gmail.com",
			Phone_number : "08946737892",
			Alamat : "Depok",
		}, {
			Nama_nasabah : "Juwita",
			Email : "juwita@gmail.com",
			Phone_number : "08123478910",
			Alamat : "Bandung",
		},
	}
	
	hasil:= module.InsertBank(module.MongoConn, "bank", nama_bank, lokasi, total_tagihan, daftar, biodata)
	fmt.Println(hasil)
}

func TestGetBankFromDaftar(t *testing.T) {
	nama_penagih := "Marlina"
	biodata:= module.GetBankFromDaftar(nama_penagih, "bank", module.MongoConn)
	fmt.Println(biodata)
}

// func TestGetBankAll(t *testing.T) {
// 	biodata := module.GetTagihanAll()
// 	fmt.Println(biodata)
// }
