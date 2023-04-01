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
			Total_Tagihan : "3500000",
			Deskripsi : "Kredit Investasi",
			Status : "Lunas",
			Tanggal_jatuhtempo : "23 Maret 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Rahma",
				Email : "rahma@gmail.com",
				Phone_number : "0878888217283",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 95.312312,
		Latitude : 244.262626,
		}, {
			Total_Tagihan : "5000000",
			Deskripsi : "Kartu Kredit",
			Status : "Lunas",
			Tanggal_jatuhtempo : "12 Juni 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Bella",
				Email : "bella@gmail.com",
				Phone_number : "08627836788",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 96.424242,
		Latitude : 214.545454,
		}, {
			Total_Tagihan : "5000000",
			Deskripsi : "Kartu Kredit",
			Status : "Lunas",
			Tanggal_jatuhtempo : "11 Mei 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Asri",
				Email : "asri@gmail.com",
				Phone_number : "087883456211",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 97.848484,
		Latitude : 215.232323,
		}, {
			Total_Tagihan : "12000000",
			Deskripsi : "Kredit Perumahan",
			Status : "Belum Lunas",
			Tanggal_jatuhtempo : "13 Juni 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Fasha",
				Email : "fasha@gmail.com",
				Phone_number : "08345261780",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 99.464646,
		Latitude : 211.212121,
		}, {
			Total_Tagihan : "14000000",
			Deskripsi : "Kredit Perumahan",
			Status : "Lunas",
			Tanggal_jatuhtempo : "9 Agustus 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Silipa",
				Email : "silipa@gmail.com",
				Phone_number : "081245362178",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 91.123122,
		Latitude : 204.243244,
		}, {
			Total_Tagihan : "20000000",
			Deskripsi : "Kartu Kredit",
			Status : "Lunas",
			Tanggal_jatuhtempo : "20 September 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Rismaya",
				Email : "rismaya@gmail.com",
				Phone_number : "086736779300",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 90.09090,
		Latitude : 201.43535,
		}, {
			Total_Tagihan : "7000000",
			Deskripsi : "Kartu Kendaraan Bermotor",
			Status : "Lunas",
			Tanggal_jatuhtempo : "10 Oktober 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Stefany",
				Email : "stefany@gmail.com",
				Phone_number : "081235233627",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 90.19191,
		Latitude : 203.67676,
		}, 
	}

	daftar := []model.Penagih{
		{
			Nama_penagih : "Aryka",
			Email : "aryka@gmail.com",
			Phone_number : "08123675489",
			Total_Tagihan : model.Tagihan{
				Total_Tagihan : "3500000",
				Deskripsi : "Kredit Investasi",
				Status : "Belum Lunas",
				Tanggal_jatuhtempo : "23 Maret 2021",
				Biodata : model.Nasabah{
					Nama_nasabah : "Rahma",
					Email : "rahma@gmail.com",
					Phone_number : "0878888217283",
					Alamat : "Bandung",
				},
				Location : "Bandung",
				Longitude : 95.312312,
				Latitude : 244.262626,
			},
	}, {
		Nama_penagih : "Syaila",
		Email : "syaila@gmail.com",	
		Phone_number : "08123415228",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "6000000",
			Deskripsi : "Kartu Kredit",
			Status : "Lunas",
			Tanggal_jatuhtempo : "12 Juni 2021",
		Biodata : model.Nasabah{
			Nama_nasabah : "Bella",
			Email : "bella@gmail.com",
			Phone_number : "08627836788",
			Alamat : "Bandung",
	},
	Location : "Bandung",
	Longitude : 96.424242,
	Latitude : 214.545454,
		},
			
	}, {
		Nama_penagih : "Alnovianti",
		Email : "novi@gmail.com",
		Phone_number : "0812452678",
		Total_Tagihan : model.Tagihan{
				Total_Tagihan : "5000000",
				Deskripsi : "Kartu Kredit",
				Status : "Lunas",
				Tanggal_jatuhtempo : "11 Mei 2021",
		Biodata : model.Nasabah{
				Nama_nasabah : "Asri",
				Email : "asri@gmail.com",
				Phone_number : "087883456211",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 96.424242,
		Latitude : 214.545454,
	},

}, {
	Nama_penagih : "Saw",
		Email : "saw@gmail.com",	
		Phone_number : "08462562627",
		Total_Tagihan : model.Tagihan{
				Total_Tagihan : "12000000",
				Deskripsi : "Kredit Perumahan",
				Status : "Belum Lunas",
				Tanggal_jatuhtempo : "13 Juni 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Fasha",
				Email : "fasha@gmail.com",
				Phone_number : "08345261780",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 99.464646,
		Latitude : 211.212121,
},
}, {
	Nama_penagih : "Putri",
		Email : "putri@gmail.com",	
		Phone_number : "08638738990",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "14000000",
			Deskripsi : "Kredit Perumahan",
			Status : "Lunas",
			Tanggal_jatuhtempo : "9 Agustus 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Silipa",
				Email : "silipa@gmail.com",
				Phone_number : "081245362178",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 91.123122,
		Latitude : 204.243244,
},
}, {
	Nama_penagih : "Renjun",
		Email : "renjun@gmail.com",	
		Phone_number : "089172537492",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "20000000",
			Deskripsi : "Kartu Kredit",
			Status : "Lunas",
			Tanggal_jatuhtempo : "20 September 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Rismaya",
				Email : "rismaya@gmail.com",
				Phone_number : "086736779300",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 90.09090,
		Latitude : 201.43535,
},
}, {
	Nama_penagih : "Mark",
		Email : "mark@gmail.com",	
		Phone_number : "0812748912",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "7000000",
			Deskripsi : "Kartu Kendaraan Bermotor",
			Status : "Lunas",
			Tanggal_jatuhtempo : "10 Oktober 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Stefany",
				Email : "stefany@gmail.com",
				Phone_number : "081235233627",
				Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 90.19191,
		Latitude : 203.67676,
},
}, 
	}

	biodata := []model.Nasabah{
		{
			Nama_nasabah : "Rahma",
			Email : "rahma@gmail.com",
			Phone_number : "0878888217283",
			Alamat : "Bandung",
		}, {
			Nama_nasabah : "Bella",
			Email : "bella@gmail.com",
			Phone_number : "08627836788",
			Alamat : "Bandung",
		}, {
			Nama_nasabah : "Asri",
			Email : "asri@gmail.com",
			Phone_number : "087883456211",
			Alamat : "Bandung",
		}, {
			Nama_nasabah : "Fasha",
			Email : "fasha@gmail.com",
			Phone_number : "08345261780",
			Alamat : "Bandung",
		}, {
			Nama_nasabah : "Silipa",
			Email : "silipa@gmail.com",
			Phone_number : "081245362178",
			Alamat : "Bandung",
		}, {
			Nama_nasabah : "Rismaya",
			Email : "rismaya@gmail.com",
			Phone_number : "086736779300",
			Alamat : "Bandung",
		}, {
			Nama_nasabah : "Stefany",
			Email : "stefany@gmail.com",
			Phone_number : "081235233627",
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



func TestGetAllTagihan(t *testing.T) {
	biodata := module.GetAllTagihan( module.MongoConn, "tagihan")
	fmt.Println(biodata)
}
