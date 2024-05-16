package models

import (
	"fmt"
	"latihangin/configs"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type DataUser struct {
	Id_user      uint
	Nama         string
	Email        string
	Password     string
	JenisKelamin string
	Pekerjaan    string
	Hobi         string
	WarnaFavorit string
	TanggalLahir time.Time
	Alamat       string
	Foto         string
	Telp         string
}

func LihatUser() []DataUser {
	var list []DataUser
	querySQL := `SELECT 
	id_user,
	nama,
	email,
	telp,
	jenis_kelamin,
	pekerjaan,hobi,
	warna_favorit,
	tanggal_lahir,
	alamat,
	foto 
	FROM users`

	result := configs.QueryRead(querySQL)
	defer result.Close()
	if result != nil {
		for result.Next() {
			dataUser := DataUser{}
			err := result.Scan(
				&dataUser.Id_user,
				&dataUser.Nama,
				&dataUser.Email,
				&dataUser.Telp,
				&dataUser.JenisKelamin,
				&dataUser.Pekerjaan,
				&dataUser.Hobi,
				&dataUser.WarnaFavorit,
				&dataUser.TanggalLahir,
				&dataUser.Alamat,
				&dataUser.Foto,
			)
			if err != nil {
				log.Fatal(err.Error())
			}
			// tanggalLahirParsed, err := time.Parse("2006-01-02", dataUser.TanggalLahir.Format("2006-01-02"))
			// if err != nil {
			// 	log.Fatal(err.Error())
			// }
			// dataUser.TanggalLahir = tanggalLahirParsed
			list = append(list, dataUser)
		}
		return list
	} else {
		return nil
	}
}

func TambahUser(
	nama string,
	email string,
	telp string,
	password string,
	jenisKelamin string,
	pekerjaan string,
	hobi string,
	warnaFavorit string,
	tanggalLahir time.Time,
	alamat string,
	foto string,
) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}

	querySQL := `INSERT INTO users
	(
		email, 
		password, 
		nama,
		telp,
		jenis_kelamin,
		hobi,
		pekerjaan,
		foto,
		tanggal_lahir,
		warna_favorit,
		alamat
	) 
	VALUES
	(
		'` + email + `', 
		'` + string(hashedPassword) + `', 
		'` + nama + `',
		'` + telp + `',
		'` + jenisKelamin + `',
		'` + hobi + `',
		'` + pekerjaan + `',
		'` + foto + `',
		'` + tanggalLahir.Format("2006-01-02") + `',
		'` + warnaFavorit + `',
		'` + alamat + `'
	)`

	configs.QueryWrite(querySQL)
	fmt.Println("Data berhasil ditambahkan")
}

func CekUser(email string) bool {
	querySQL := `SELECT id_user FROM users WHERE email = '` + email + `'`

	result := configs.QueryRead(querySQL)
	if result != nil {
		count := 0
		for result.Next() {
			count += 1
		}

		if count > 0 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func DetailUser(id_user int) DataUser {
	querySQL := `SELECT id_user,
	nama,
	email,
	telp,
	jenis_kelamin,
	pekerjaan,hobi,
	warna_favorit,
	tanggal_lahir,
	alamat,
	foto  
	FROM users WHERE id_user = ` + strconv.Itoa(id_user)

	result := configs.QueryRead(querySQL)
	defer result.Close()
	if result != nil {
		result.Next()
		dataUser := DataUser{}
		result.Scan(
			&dataUser.Id_user,
			&dataUser.Nama,
			&dataUser.Email,
			&dataUser.Telp,
			&dataUser.JenisKelamin,
			&dataUser.Pekerjaan,
			&dataUser.Hobi,
			&dataUser.WarnaFavorit,
			&dataUser.TanggalLahir,
			&dataUser.Alamat,
			&dataUser.Foto,
		)

		return dataUser
	} else {
		return DataUser{} //struct tidak bisa di nil kan
	}
}

func DetailUserByEmail(email string) DataUser {
	querySQL := `SELECT id_user,
	nama,
	email,
	telp,
	jenis_kelamin,
	pekerjaan,hobi,
	warna_favorit,
	tanggal_lahir,
	alamat,
	foto  
	FROM users WHERE email = '` + email + `'`

	result := configs.QueryRead(querySQL)
	defer result.Close()
	if result != nil {
		result.Next()
		dataUser := DataUser{}
		result.Scan(
			&dataUser.Id_user,
			&dataUser.Nama,
			&dataUser.Email,
			&dataUser.Telp,
			&dataUser.JenisKelamin,
			&dataUser.Pekerjaan,
			&dataUser.Hobi,
			&dataUser.WarnaFavorit,
			&dataUser.TanggalLahir,
			&dataUser.Alamat,
			&dataUser.Foto,
		)

		return dataUser
	} else {
		return DataUser{} //struct tidak bisa di nil kan
	}
}

func HapusData(nomor int) {
	querySQL := "DELETE FROM users WHERE id_user = " + strconv.Itoa(nomor)

	configs.QueryWrite(querySQL)
	fmt.Println("Data dengan ID " + strconv.Itoa(nomor) + " telah dihapus")

}

func UbahUser(
	id_user int,
	nama string,
	telp string,
	jenisKelamin string,
	pekerjaan string,
	hobi string,
	warnaFavorit string,
	tanggalLahir time.Time,
	alamat string,
	foto string,
) {
	var querySQL string
	if foto == "" {
		querySQL = `UPDATE users SET 
		nama = '` + nama + `' ,
		telp = '` + telp + `',
		jenis_kelamin = '` + jenisKelamin + `',
		pekerjaan = '` + pekerjaan + `',
		hobi = '` + hobi + `',
		warna_favorit = '` + warnaFavorit + `',
		tanggal_lahir = '` + tanggalLahir.Format("2006-01-02") + `',
		alamat = '` + alamat + `'
		WHERE id_user = ` + strconv.Itoa(id_user)
	} else {
		querySQL = `UPDATE users SET 
		nama = '` + nama + `' ,
		telp = '` + telp + `',
		jenis_kelamin = '` + jenisKelamin + `',
		pekerjaan = '` + pekerjaan + `',
		hobi = '` + hobi + `',
		warna_favorit = '` + warnaFavorit + `',
		tanggal_lahir = '` + tanggalLahir.Format("2006-01-02") + `',
		alamat = '` + alamat + `',
		foto = '` + foto + `'
		WHERE id_user = ` + strconv.Itoa(id_user)
	}

	configs.QueryWrite(querySQL)
	fmt.Println("Data berhasil diubah")

}

func LoginUser(email string, password string) bool {

	var hashedPassword string

	querySQL := "SELECT password FROM users WHERE email = '" + email + "'"
	result := configs.QueryRead(querySQL)
	//fmt.Println(result == nil)
	defer result.Close()
	if result != nil {
		if result.Next() {
			errScan := result.Scan(&hashedPassword)
			if errScan != nil {
				log.Fatal("Errornya:" + errScan.Error())
			}
			err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
			if err != nil {
				fmt.Println(err.Error())
				return false
			} else {

				return true
			}
		} else {
			return false
		}
	} else {
		return false
	}

}
