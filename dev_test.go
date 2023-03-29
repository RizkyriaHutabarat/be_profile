package profile

import (
	"fmt"
	"testing"
	model "github.com/rizkyriahutabarat/be_profile/model"
	module "github.com/rizkyriahutabarat/be_profile/module"
)

func TestInsertUser(t *testing.T) {
	username := "kia"
	email := "kia12@gmail.com"
	pass := "cantik"
	hasil := module.InsertUser(module.MongoConn, "user",username, email, pass)
	fmt.Println(hasil)
}

func TestInsertPendidikan(t *testing.T) {
	userid := "123"
	sekolah := "ulbi"
	lulusan := "2020"
	tahunmulai := "2016"
	tahunselesai := "2020"
	hasil := module.InsertPendidikan(module.MongoConn, "pendidikan",userid, sekolah, lulusan, tahunmulai, tahunselesai)
	fmt.Println(hasil)
}

func TestInsertPengalaman(t *testing.T) {
	userid := "102"
	perusahaan := "PT.Pos"
	jabatan := "CEO"
	deskripsi := "Ini CEO Pos"
	tahunmulai := "2016"
	tahunselesai := "2020"
	hasil := module.InsertPengalaman(module.MongoConn, "pengalaman",userid, perusahaan, jabatan, deskripsi, tahunmulai, tahunselesai)
	fmt.Println(hasil)
}

func TestInsertSkill(t *testing.T) {
	nama := "go"
	level := "8"
	hasil := module.InsertSkill(module.MongoConn, "skill",nama, level)
	fmt.Println(hasil)
}


func TestInsertProfile(t *testing.T) {
	nama_user := "Rizkyria"
	data_pendidikan := []model.Pendidikan{
		{
			UserID: "12",
			Sekolah : "ULBI",
			Lulusan : "S1",
			Tahunmulai : "2020",
			Tahunselesai : "2023",
		}, {
			UserID: "13",
			Sekolah : "IPB",
			Lulusan : "S2",
			Tahunmulai : "2023",
			Tahunselesai : "2025",
		},
	}
	data_pengalaman := []model.Pengalaman{
		{
			UserID  : "1234",
			Perusahaan : "PT POS Indonesia",
			Jabatan : "Direktur",
			Deskripsi : "Direktur PT POS Indonesia",
			Tahunmulai : "2023",
			Tahunselesai : "2045",
		}, {
			UserID : "1345",
			Perusahaan : "PT BTN Indonesia",
			Jabatan : "Ceo",
			Deskripsi : "Ceo PT BTN Indonesia",
			Tahunmulai : "2023",
			Tahunselesai : "2045",
		}, {
			UserID  : "1452",
			Perusahaan : "PT BCA Indonesia",
			Jabatan : "Ceo",
			Deskripsi : "Ceo PT BCA Indonesia",
			Tahunmulai : "2012",
			Tahunselesai : "2023",
		},

	}
	skills := []model.Skill{
		 {
			Nama : " GO ",
			Level : "Medium",
	}, 
		{
		Nama : " HTML ",
		Level : "8",
	},
	
	}
	hasil:=module.InsertProfile(module.MongoConn, "profile",nama_user , data_pendidikan , data_pengalaman, skills)
	fmt.Println(hasil)
}

func TestGetUserFromEmail(t *testing.T) {
	email := "kia12@gmail.com"
	biodata := module.GetUserFromEmail(email,module.MongoConn, "user")
	fmt.Println(biodata)
}




func TestGetPendidikanFromSekolah(t *testing.T) {
	sekolah := "Politeknik POS Indonesia"
	biodata := module.GetPendidikanFromSekolah(sekolah,module.MongoConn, "pendidikan")
	fmt.Println(biodata)
}



func TestGetPengalamanFromJabatan(t *testing.T) {
	jabatan := "CEO"
	biodata := module.GetPengalamanFromJabatan(jabatan,module.MongoConn, "pengalaman")
	fmt.Println(biodata)
}



func TestGetSkillFromNama(t *testing.T) {
	nama := "CSS"
	biodata := module.GetSkillFromNama(nama,module.MongoConn, "skill")
	fmt.Println(biodata)
}



func TestGetProfileFromNama_user(t *testing.T) {
	nama_user := "kia"
	biodata := module.GetProfileFromNama_user(nama_user,module.MongoConn, "profile")
	fmt.Println(biodata)
}

func TestGetAllProfile(t *testing.T) {
	nama_user := "kia"
	biodata := module.GetProfileFromNama_user(nama_user,module.MongoConn, "profile")
	fmt.Println(biodata)
}


