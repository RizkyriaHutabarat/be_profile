package module

import (
	"context"
	"fmt"
	"github.com/rizkyriahutabarat/be_profile/model"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aiteung/atdb"

)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "profile_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}



func InsertUser(db *mongo.Database, col string,username string, email string, pass string) (InsertedID interface{}) {
	var user model.User
	user.Username = username
	user.Email = email
	user.Password = pass
	return InsertOneDoc(db, col, user)
}

func GetUserFromEmail(email string, db *mongo.Database, col string) (user model.User) {
	data_profile := db.Collection(col)
	filter := bson.M{"email": email}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Printf("getUserFromEmail: %v\n", err)
	}
	return user
}


func InsertPendidikan(db *mongo.Database, col string,userid string, sekolah string, lulusan string,tahunmulai string, tahunselesai string) (InsertedID interface{}) {
	var pendidikan model.Pendidikan
	pendidikan.UserID = userid
	pendidikan.Sekolah = sekolah
	pendidikan.Lulusan = lulusan
	pendidikan.Tahunmulai = tahunmulai
	pendidikan.Tahunselesai = tahunselesai
	return InsertOneDoc(db,col, pendidikan)
}


func GetPendidikanFromSekolah(sekolah string, db *mongo.Database, col string) (pendidikan model.Pendidikan) {
	data_profile := db.Collection(col)
	filter := bson.M{"sekolah": sekolah}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&pendidikan)
	if err != nil {
		fmt.Printf("getPendidikanFromSekolah: %v\n", err)
	}
	return pendidikan
}

func InsertPengalaman(db *mongo.Database, col string, userid string, perusahaan string, jabatan string, deskripsi string, tahunmulai string, tahunselesai string) (InsertedID interface{}) {
	var pengalaman model.Pengalaman
	pengalaman.UserID = userid
	pengalaman.Perusahaan = perusahaan
	pengalaman.Jabatan = jabatan
	pengalaman.Deskripsi = deskripsi
	pengalaman.Tahunmulai = tahunmulai
	pengalaman.Tahunselesai = tahunselesai
	return InsertOneDoc(db, col, pengalaman)
}


func GetPengalamanFromJabatan(jabatan string, db *mongo.Database, col string) (pengalaman model.Pengalaman) {
	data_profile := db.Collection(col)
	filter := bson.M{"jabatan": jabatan}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&pengalaman)
	if err != nil {
		fmt.Printf("getPengalamanFromJabatan: %v\n", err)
	}
	return pengalaman
}



func InsertSkill(db *mongo.Database, col string,nama string, level string) (InsertedID interface{}) {
	var skill model.Skill
	skill.Nama = nama
	skill.Level = level
	return InsertOneDoc(db,col, skill)
}



func GetSkillFromNama(nama string, db *mongo.Database, col string) (skill model.Skill) {
	data_profile := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&skill)
	if err != nil {
		fmt.Printf("getSkillFromNama: %v\n", err)
	}
	return skill
}



func InsertProfile( db *mongo.Database, col string,nama_user string,  data_pendidikan []model.Pendidikan, data_pengalaman []model.Pengalaman, skills []model.Skill) (InsertedID interface{}) {
	var profile model.Profile
	profile.Nama_user = nama_user
	profile.Data_pendidikan = data_pendidikan
	profile.Data_pengalaman = data_pengalaman
	profile.Skills = skills
	return InsertOneDoc(db,col, profile)
}




func GetProfileFromNama_user(nama_user string, db *mongo.Database, col string) (profile model.Profile) {
	data_profile := db.Collection(col)
	filter := bson.M{"nama_user": nama_user}
	err := data_profile.FindOne(context.TODO(), filter).Decode(&profile)
	if err != nil {
		fmt.Printf("getProfileFromNama_user: %v\n", err)
	}
	return profile
}


func GetAllProfileFromNama_user(nama_user string, db *mongo.Database, col string) (profile model.Profile) {
	data_profile := db.Collection(col)
	filter := bson.M{"nama_user": nama_user}
	cursor, err := data_profile.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllProfileFromNama_user: %v\n", err)
	}
	err = cursor.All(context.TODO(), &profile)
	if err != nil{
		fmt.Println(err)
	}
	return 
}

// func GetAllUser(db *mongo.Database, col string) (user model.User) {
// 	data_profile := db.Collection(col)
// 	filter := bson.M{}
// 	cursor, err := data_profile.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetALLUser :", err)
// 	}
// 	err = cursor.All(context.TODO(), &user)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

// func GetAllPendidikan(db *mongo.Database, col string) (pendidikan model.Pendidikan) {
// 	data_profile := db.Collection(col)
// 	filter := bson.M{}
// 	cursor, err := data_profile.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetALLPendidikan :", err)
// 	}
// 	err = cursor.All(context.TODO(), &pendidikan)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

// func GetAllPengalaman(db *mongo.Database, col string) (pengalaman model.Pengalaman) {
// 	data_profile := db.Collection(col)
// 	filter := bson.M{}
// 	cursor, err := data_profile.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetALLPengalaman :", err)
// 	}
// 	err = cursor.All(context.TODO(), &pengalaman)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }


func GetAllProfile(db *mongo.Database, col string) (profile model.Profile) {
	data_profile := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_profile.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllProfile :", err)
	}
	err = cursor.All(context.TODO(), &profile)
	if err != nil {
		fmt.Println(err)
	}
	return
}
