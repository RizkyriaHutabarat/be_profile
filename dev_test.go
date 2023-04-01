package profile

import (
	"fmt"
	"testing"
	model "github.com/rizkyriahutabarat/be_profile/model"
	module "github.com/rizkyriahutabarat/be_profile/module"
)

func TestInsertUser(t *testing.T) {
	username := "Marlina"
	email := "marlina@gmail.com"
	pass := "123456"
	hasil := module.InsertUser(module.MongoConn, "user",username, email, pass)
	fmt.Println(hasil)
}

func TestInsertPendidikan(t *testing.T) {
	userid := "176"
	sekolah := "ULBI"
	lulusan := "2025"
	tahunmulai := "2021"
	tahunselesai := "2025"
	hasil := module.InsertPendidikan(module.MongoConn, "pendidikan",userid, sekolah, lulusan, tahunmulai, tahunselesai)
	fmt.Println(hasil)
}

func TestInsertPengalaman(t *testing.T) {
	userid := "1527"
	perusahaan := "PT POS Indonesia"
	jabatan := "Mahasiswa Magang"
	deskripsi := "Mahasiswa Semester 4"
	tahunmulai := "2023"
	tahunselesai := "2023"
	hasil := module.InsertPengalaman(module.MongoConn, "pengalaman",userid, perusahaan, jabatan, deskripsi, tahunmulai, tahunselesai)
	fmt.Println(hasil)
}

func TestInsertSkill(t *testing.T) {
	nama := "Go"
	level := "8"
	hasil := module.InsertSkill(module.MongoConn, "skill",nama, level)
	fmt.Println(hasil)
}


func TestInsertProfile(t *testing.T) {
	nama_user := "Yanti"
	data_pendidikan := []model.Pendidikan{
		// {
		// 	UserID: "12",
		// 	Sekolah : "ULBI",
		// 	Lulusan : "S1",
		// 	Tahunmulai : "2020",
		// 	Tahunselesai : "2023",
		// }, {
		// 	UserID: "13",
		// 	Sekolah : "IPB",
		// 	Lulusan : "S2",
		// 	Tahunmulai : "2023",
		// 	Tahunselesai : "2025",
		// }, {
		// 	UserID: "14",
		// 	Sekolah : "UGM",
		// 	Lulusan : "S3",
		// 	Tahunmulai : "2025",
		// 	Tahunselesai : "2026",
		// },
		 {
			UserID: "196",
			Sekolah : "UKDW",
			Lulusan : "S1",
			Tahunmulai : "1990",
			Tahunselesai : "1995",
		}, {
			UserID: "197",
			Sekolah : "USU",
			Lulusan : "S2",
			Tahunmulai : "2000",
			Tahunselesai : "2002",
		}, 
		// {
		// 	UserID: "56",
		// 	Sekolah : "USU",
		// 	Lulusan : "S3",
		// 	Tahunmulai : "2022",
		// 	Tahunselesai : "2023",
		// }, 
	}
	data_pengalaman := []model.Pengalaman{
		// {
		// 	UserID  : "1234",
		// 	Perusahaan : "PT POS Indonesia",
		// 	Jabatan : "Direktur",
		// 	Deskripsi : "Direktur PT POS Indonesia",
		// 	Tahunmulai : "2023",
		// 	Tahunselesai : "2024",
		// }, {
		// 	UserID : "1345",
		// 	Perusahaan : "PT BTN Indonesia",
		// 	Jabatan : "Ceo",
		// 	Deskripsi : "Ceo PT BTN Indonesia",
		// 	Tahunmulai : "2024",
		// 	Tahunselesai : "2025",
		// }, {
		// 	UserID  : "1452",
		// 	Perusahaan : "PT BCA Indonesia",
		// 	Jabatan : "Ceo",
		// 	Deskripsi : "Ceo PT BCA Indonesia",
		// 	Tahunmulai : "2025",
		// 	Tahunselesai : "2026",
		// }, 
		{
			UserID  : "1577",
			Perusahaan : "PT Citra Hutata",
			Jabatan : "Direktur",
			Deskripsi : "Direktur PT Curta Hutata",
			Tahunmulai : "2002",
			Tahunselesai : "2010",
		}, {
			UserID  : "1578",
			Perusahaan : "PT BCA Indonesia",
			Jabatan : "CEO Divisi 4",
			Deskripsi : "CEO Divisi 4 PT BCA Indonesia",
			Tahunmulai : "2011",
			Tahunselesai : "2023",
		}, 
		// {
		// 	UserID  : "1093",
		// 	Perusahaan : "PT BTN Indonesia",
		// 	Jabatan : "Mahasiswa Magang",
		// 	Deskripsi : "Laporan Internship  ",
		// 	Tahunmulai : "2022",
		// 	Tahunselesai : "2023",
		// },

	}
	skills := []model.Skill{
	// 	 {
	// 		Nama : "GO",
	// 		Level : "9",
	// }, 
	// 	{
	// 	Nama : "HTML",
	// 	Level : "8",
	// },
	// 	{
	// 	Nama : "CSS",
	// 	Level : "9",
	// }, 
	{
		Nama : "Komunikasi",
		Level : "9",
	}, {
		Nama : "Marketing",
		Level : "9",
	}, {
		Nama : "Sales",
		Level : "8",
	}, 

	}
	
	hasil:=module.InsertProfile(module.MongoConn, "profile",nama_user , data_pendidikan , data_pengalaman, skills)
	fmt.Println(hasil)
}


func TestGetUserFromEmail(t *testing.T) {
	email := "marlina@gmail.com"
	biodata := module.GetUserFromEmail(email,module.MongoConn, "user")
	fmt.Println(biodata)
}




func TestGetPendidikanFromSekolah(t *testing.T) {
	sekolah := "ULBI"
	biodata := module.GetPendidikanFromSekolah(sekolah,module.MongoConn, "pendidikan")
	fmt.Println(biodata)
}



func TestGetPengalamanFromJabatan(t *testing.T) {
	jabatan := "Mahasiswa"
	biodata := module.GetPengalamanFromJabatan(jabatan,module.MongoConn, "pengalaman")
	fmt.Println(biodata)
}



func TestGetSkillFromNama(t *testing.T) {
	nama := "Ptyhon"
	biodata := module.GetSkillFromNama(nama,module.MongoConn, "skill")
	fmt.Println(biodata)
}



func TestGetProfileFromNama_user(t *testing.T) {
	nama_user := "pandapotan"
	biodata := module.GetProfileFromNama_user(nama_user,module.MongoConn, "profile")
	fmt.Println(biodata)
}

// func TestGetAllProfile(t *testing.T) {
// 	nama_user := "kia"
// 	biodata := module.GetProfileFromNama_user(nama_user,module.MongoConn, "profile")
// 	fmt.Println(biodata)
// }

func TestGetAllProfile(t *testing.T) {
	biodata := module.GetAllProfile(module.MongoConn, "profile")
	fmt.Println(biodata)
}


