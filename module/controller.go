package module

import (
	"context"
	"fmt"
	"os"
	
    model "github.com/Fedhira/Tagihan/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// NASABAH
func InsertNasabah(nama_nasabah string, email string, phone_number string, alamat string) (InsertedID interface{}) {
	var nasabah model.Nasabah
	nasabah.Nama_nasabah = nama_nasabah
	nasabah.Email = email
	nasabah.Phone_number = phone_number
	nasabah.Alamat = alamat
	return InsertOneDoc("db_tagihan", "nasabah", nasabah)
}

func GetNasabahFromNama(nama_nasabah string) (nasabah model.Nasabah) {
	data_nasabah := MongoConnect("db_tagihan").Collection("nasabah")
	filter := bson.M{"nama_nasabah": nama_nasabah}
	err := data_nasabah.FindOne(context.TODO(), filter).Decode(&nasabah)
	if err != nil {
		fmt.Printf("getNasabahFromNama: %v\n", err)  
	}
	return nasabah
}


func GetNasabahAll() (nasabah []model.Nasabah) {
	data_nasabah := MongoConnect("db_tagihan").Collection("nasabah")
	filter := bson.D{}
	// var results []Nasabah
	cur, err := data_nasabah.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetNasabahFromNama: %v\n", err)
	}
	err = cur.All(context.TODO(), &nasabah)
	if err != nil {
		fmt.Println(err)
	}
	return nasabah
}

// PENAGIH
func InsertPenagih(nama_penagih string, email string, phone_number string, total_tagihan  model.Tagihan) (InsertedID interface{}) {
	var penagih model.Penagih
	penagih.Nama_penagih = nama_penagih
	penagih.Email = email
	penagih.Phone_number = phone_number
	penagih.Total_Tagihan = total_tagihan
	return InsertOneDoc("db_tagihan", "penagih", penagih)
}

func GetPenagihFromNama(nama_penagih string) (penagih model.Penagih) {
	data_penagih := MongoConnect("db_tagihan").Collection("penagih")
	filter := bson.M{"nama_penagih": nama_penagih}
	err := data_penagih.FindOne(context.TODO(), filter).Decode(&penagih)
	if err != nil {
		fmt.Printf("getPenagihFromNama: %v\n", err)  
	}
	return penagih
}

func GetPenagihAll() (penagih []model.Penagih) {
	data_penagih := MongoConnect("db_tagihan").Collection("penagih")
	filter := bson.D{}
	// var results []Penagih
	cur, err := data_penagih.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetPenagihFromNama: %v\n", err)
	}
	err = cur.All(context.TODO(), &penagih)
	if err != nil {
		fmt.Println(err)
	}
	return penagih
}

// TAGIHAN
func InsertTagihan(total_tagihan string, deskripsi string,  status string, tanggal_jatuhtempo string, biodata  model.Nasabah, location string, longitude float64, latitude float64  ) (InsertedID interface{}) {
	var tagihan model.Tagihan
	tagihan.Total_Tagihan = total_tagihan
	tagihan.Deskripsi = deskripsi
	tagihan.Status = status
	// tagihan.Tanggal_jatuhtempo = primitive.NewDateTimeFromTime(time.Now().UTC())
	tagihan.Biodata = biodata
	tagihan.Location = location
	tagihan.Longitude = longitude
	tagihan.Latitude = latitude
	return InsertOneDoc("db_tagihan", "tagihan", tagihan)
}

func GetTagihanFromNama_nasabah(nama_nasabah string) (tagihan model.Tagihan) {
	data_tagihan := MongoConnect("db_tagihan").Collection("tagihan")
	filter := bson.M{"biodata.nama_nasabah": nama_nasabah}
	err := data_tagihan.FindOne(context.TODO(), filter).Decode(&tagihan)
	if err != nil {
		fmt.Printf("getTagihanFromNama_nasabah: %v\n", err)  
	}
	return tagihan
}

func GetTagihanAll() (tagihan []model.Tagihan) {
	data_tagihan := MongoConnect("db_tagihan").Collection("tagihan")
	filter := bson.D{}
	// var results []Tagihan
	cur, err := data_tagihan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetTagihanFromNama_nasabah: %v\n", err)
	}
	err = cur.All(context.TODO(), &tagihan)
	if err != nil {
		fmt.Println(err)
	}
	return tagihan
}

// BANK
func InsertBank(nama_bank string, lokasi string, total_tagihan []model.Tagihan,  daftar []model.Penagih, biodata []model.Nasabah) (InsertedID interface{}) {
	var bank model.Bank
	bank.Nama_bank = nama_bank
	bank.Lokasi = lokasi
	bank.Total_Tagihan = total_tagihan
	bank.Daftar = daftar
	bank.Biodata = biodata
	return InsertOneDoc("db_tagihan", "bank", bank)
}

func GetBankFromDaftar(nama_penagih string) (bank model.Bank) {
	data_bank := MongoConnect("db_tagihan").Collection("bank")
	filter := bson.M{"daftar.nama_penagih": nama_penagih}
	err := data_bank.FindOne(context.TODO(), filter).Decode(&bank)
	if err != nil {
		fmt.Printf("getBankFromDaftar: %v\n", err)  
	}
	return bank
}

func GetBankAll() (bank []model.Bank) {
	data_bank := MongoConnect("db_tagihan").Collection("bank")
	filter := bson.D{}
	// var results []Bank
	cur, err := data_bank.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetBankFromDaftar: %v\n", err)
	}
	err = cur.All(context.TODO(), &bank)
	if err != nil {
		fmt.Println(err)
	}
	return bank
}
