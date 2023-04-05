package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID        primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
    Username  string 		      `bson:"username,omitempty" json:"username,omitempty"`
    Email     string 	          `bson:"email,omitempty" json:"email,omitempty"`
    Password  string 	          `bson:"password,omitempty" json:"password,omitempty"`
}

type Pendidikan struct {
    ID            primitive.ObjectID    	`bson:"_id,omitempty" json:"_id,omitempty"`
    UserID        string   					`bson:"userid,omitempty" json:"userid,omitempty"`
    Sekolah       string 				 	`bson:"sekolah,omitempty" json:"sekolah,omitempty"`
    Lulusan       string  				  	`bson:"lulusan,omitempty" json:"lulusan,omitempty"`
    Tahunmulai    string 	                `bson:"tahunmulai,omitempty" json:"tahunmulai,omitempty"`
    Tahunselesai  string 	            	`bson:"tahunselesai,omitempty" json:"tahunselesai,omitempty"`
}

type Pengalaman struct {
    ID               primitive.ObjectID    	`bson:"_id,omitempty" json:"_id,omitempty"`
    UserID           string    				`bson:"userid,omitempty" json:"userid,omitempty"`
    Perusahaan       string                 `bson:"perusahaan,omitempty" json:"perusahaan,omitempty"`
    Jabatan          string                 `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
    Deskripsi        string                 `bson:"deskripsi,omitempty" json:"deksripsi,omitempty"`
    Tahunmulai       string                 `bson:"tahunmulai,omitempty" json:"tahunmulai,omitempty"`
    Tahunselesai     string                 `bson:"tahunselesai,omitempty" json:"tahunselesai,omitempty"`
}

type Skill struct {
    ID      primitive.ObjectID    	`bson:"_id,omitempty" json:"_id,omitempty"`
    Nama 	string 	                `bson:"nama,omitempty" json:"nama,omitempty"`
	Level 	string 	                `bson:"level,omitempty" json:"level,omitempty"`
}

// type Profile struct {
//     Nama_user        string           `bson:"nama_user,omitempty" json:"nama_user,omitempty"`
//     Data_pendidikan  []Pendidikan   `bson:"data_pendidikan,omitempty" json:"data_pendidikan,omitempty"`
//     Data_pengalaman  []Pengalaman   `bson:"data_pengalaman,omitempty" json:"data_pengalaman,omitempty"`
//     Skills      []Skill         `bson:"skills,omitempty" json:"skills,omitempty"`
// }

type Profile struct {
    Nama_user        string           `bson:"nama_user,omitempty" json:"nama_user,omitempty"`
    Data_pendidikan  Pendidikan   `bson:"data_pendidikan,omitempty" json:"data_pendidikan,omitempty"`
    Data_pengalaman  Pengalaman   `bson:"data_pengalaman,omitempty" json:"data_pengalaman,omitempty"`
    Skills           Skill         `bson:"skills,omitempty" json:"skills,omitempty"`
}