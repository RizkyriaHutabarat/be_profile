package module

import (
	"context"
	"fmt"
	"github.com/rizkyriahutabarat/be_profile/model"
	"go.mongodb.org/mongo-driver/bson"
	"os"
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

func InsertUser(username string, email string, pass string) (InsertedID interface{}) {
	var user model.User
	user.Username = username
	user.Email = email
	user.Password = pass
	return InsertOneDoc("tugas_db", "user", user)
}

func GetUserFromEmail(pass string) (user model.User) {
	data_profile := MongoConnect("tugas_db").Collection("user")
	filter := bson.M{"pass": pass}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Printf("GetUserFromEmail: %v\n", err)
	}
	return user
}

func GetUserAll() (user []model.User) {
	data_profile := MongoConnect("tugas_db").Collection("user")
	filter := bson.D{}
	cur, err := data_profile.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetUserAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &user)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func InsertPendidikan(userid string, sekolah string, lulusan string,tahunmulai string, tahunselesai string) (InsertedID interface{}) {
	var pendidikan model.Pendidikan
	pendidikan.UserID = userid
	pendidikan.Sekolah = sekolah
	pendidikan.Lulusan = lulusan
	pendidikan.Tahunmulai = tahunmulai
	pendidikan.Tahunselesai = tahunselesai
	return InsertOneDoc("tugas_db",  "pendidikan", pendidikan)
}


func GetPendidikanFromSekolah(sekolah string) (pendidikan model.Pendidikan) {
	data_profile := MongoConnect("tugas_db").Collection("pendidikan")
	filter := bson.M{"sekolah": sekolah}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&pendidikan)
	if err != nil {
		fmt.Printf("GetPendidikanFromSekolah: %v\n", err)
	}
	return pendidikan
}


func GetPendidikanAll() (pendidikan []model.Pendidikan) {
	data_profile := MongoConnect("tugas_db").Collection("pendidikan")
	filter := bson.D{}
	cur, err := data_profile.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetProfileAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &pendidikan)
	if err != nil {
		fmt.Println(err)
	}
	return pendidikan
}


func InsertPengalaman(userid string, perusahaan string, jabatan string, deskripsi string, tahunmulai string, tahunselesai string) (InsertedID interface{}) {
	var pengalaman model.Pengalaman
	pengalaman.UserID = userid
	pengalaman.Perusahaan = perusahaan
	pengalaman.Jabatan = jabatan
	pengalaman.Deskripsi = deskripsi
	pengalaman.Tahunmulai = tahunmulai
	pengalaman.Tahunselesai = tahunselesai
	return InsertOneDoc("tugas_db", "pengalaman", pengalaman)
}

func GetPengalamanFromJabatan(jabatan string) (pengalaman model.Pengalaman) {
	data_profile := MongoConnect("tugas_db").Collection("pengalaman")
	filter := bson.M{"jabatan": jabatan}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&pengalaman)
	if err != nil {
		fmt.Printf("GetPengalamanFromJabatan: %v\n", err)
	}
	return pengalaman
}

func GetPengalamanAll() (pengalaman []model.Pengalaman) {
	data_profile := MongoConnect("tugas_db").Collection("pengalaman")
	filter := bson.D{}
	cur, err := data_profile.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetPengalamanAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &pengalaman)
	if err != nil {
		fmt.Println(err)
	}
	return pengalaman
}

func InsertSkill(nama string, level string) (InsertedID interface{}) {
	var skill model.Skill
	skill.Nama = nama
	skill.Level = level
	return InsertOneDoc("tugas_db", "skill", skill)
}

func GetSkillFromNama(nama string) (skill model.Skill) {
	data_profile := MongoConnect("tugas_db").Collection("skill")
	filter := bson.M{"nama": nama}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&skill)
	if err != nil {
		fmt.Printf("GetSkillFromNama: %v\n", err)
	}
	return skill
}

func GetSkillAll() (skill []model.Skill) {
	data_profile := MongoConnect("tugas_db").Collection("skill")
	filter := bson.D{}
	cur, err := data_profile.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetSkillAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &skill)
	if err != nil {
		fmt.Println(err)
	}
	return skill
}

func InsertProfile( nama_user string,  data_pendidikan []model.Pendidikan, data_pengalaman []model.Pengalaman, skills []model.Skill) (InsertedID interface{}) {
	var profile model.Profile
	profile.Nama_user = nama_user
	profile.Data_pendidikan = data_pendidikan
	profile.Data_pengalaman = data_pengalaman
	profile.Skills = skills
	return InsertOneDoc("tugas_db", "profile", profile)
}

func GetProfileFromNama_user(nama_user string) (profile model.Profile) {
	data_profile := MongoConnect("tugas_db").Collection("profile")
	filter := bson.M{"nama_user": nama_user}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&profile)
	if err != nil {
		fmt.Printf("GetProfileFromNama_user: %v\n", err)
	}
	return profile
}

func GetProfileAll() (profile []model.Profile) {
	data_profile := MongoConnect("tugas_db").Collection("profile")
	filter := bson.D{}
	cur, err := data_profile.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetProfileAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &profile)
	if err != nil {
		fmt.Println(err)
	}
	return profile
}


