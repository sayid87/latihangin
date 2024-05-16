package controllers

import (
	"latihangin/configs"
	"latihangin/models"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginUser() gin.HandlerFunc {
	return func(request *gin.Context) {
		email := request.PostForm("email")
		password := request.PostForm("password")

		cekUser := models.LoginUser(email, password)

		if cekUser {
			token, err := configs.CreateToken(email)
			if err != nil {
				log.Fatal(err.Error())
			}
			data := models.DetailUserByEmail(email)
			request.JSON(200, gin.H{"sukses": 1, "pesan": "Selamat Datang " + data.Nama, "data": data, "token": token})
		} else {
			request.JSON(200, gin.H{"sukses": 0, "pesan": "Email dan Password salah"})
		}
	}
}

func RegisterUser() gin.HandlerFunc {
	return func(request *gin.Context) {
		nama := request.PostForm("nama")
		email := request.PostForm("email")
		telp := request.PostForm("telp")
		password := request.PostForm("password")
		jenisKelamin := request.PostForm("kelamin")
		pekerjaan := request.PostForm("pekerjaan")
		hobi := request.PostFormArray("hobi")
		implodeHobi := strings.Join(hobi, ";")
		// fmt.Println("Hobi yang dipilih:", implodeHobi)
		tanggalLahir := request.PostForm("tanggal")
		warnaFavorit := request.PostForm("warna")
		alamat := request.PostForm("alamat")
		foto, errFoto := request.FormFile("foto")

		var pesanValidasi []string

		if nama == "" {
			pesanValidasi = append(pesanValidasi, "Nama harus diisi")
		}
		if email == "" || !configs.ValidateEmail(email) {
			pesanValidasi = append(pesanValidasi, "Email harus diisi dengan format yang benar")
		} else {
			cekUser := models.CekUser(email)
			if !cekUser {
				pesanValidasi = append(pesanValidasi, "Email sudah terdaftar")
			}
		}

		if telp == "" || !regexp.MustCompile(`\d`).MatchString(telp) || len(telp) < 8 || len(telp) > 13 {
			pesanValidasi = append(pesanValidasi, "Telp harus diisi antara 8 s/d 13 angka")
		}

		if password == "" {
			pesanValidasi = append(pesanValidasi, "Password harus diisi")
		}

		if jenisKelamin == "" {
			pesanValidasi = append(pesanValidasi, "Jenis kelamin harus dipilih")
		}

		if pekerjaan == "" {
			pesanValidasi = append(pesanValidasi, "Pekerjaan harus dipilih")
		}

		if len(hobi) == 0 {
			pesanValidasi = append(pesanValidasi, "Hobi harus dipilih")
		}

		if tanggalLahir == "" {
			pesanValidasi = append(pesanValidasi, "Tanggal lahir harus dipilih")
		}

		if warnaFavorit == "" {
			pesanValidasi = append(pesanValidasi, "Warna favorit harus dipilih")
		}

		if alamat == "" {
			pesanValidasi = append(pesanValidasi, "Alamat harus diisi")
		}

		if foto == nil {
			pesanValidasi = append(pesanValidasi, "Foto Harus dipilih")
		} else {
			mimeType := foto.Header.Get("Content-Type")
			if mimeType != "image/jpeg" && mimeType != "image/png" {
				pesanValidasi = append(pesanValidasi, "Foto Harus berupa file gambar jpg atau png")
			}
		}

		if len(pesanValidasi) > 0 {
			request.JSON(200, gin.H{"sukses": 0, "pesan": pesanValidasi})
		} else {
			if errFoto != nil {
				log.Fatal(errFoto.Error())
			}
			hasilTanggal, err := time.Parse("2006-01-02", tanggalLahir)
			if err != nil {
				log.Fatal(err.Error())
			}

			dst := "img/" + foto.Filename
			models.TambahUser(
				nama,
				email,
				telp,
				password,
				jenisKelamin,
				pekerjaan,
				implodeHobi,
				warnaFavorit,
				hasilTanggal,
				alamat,
				foto.Filename,
			)

			errUpload := request.SaveUploadedFile(foto, dst)
			if errUpload != nil {
				log.Fatal(errUpload.Error())
				return
			}

			request.JSON(200, gin.H{"sukses": 1, "pesan": "Data berhasil ditambah"})
		}

	}
}

func ListUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"sukses": 0, "pesan": "Unauthorized"})
		} else {
			c.JSON(200, gin.H{"sukses": 1, "data": models.LihatUser()})
		}

	}
}

func DetailUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(200, gin.H{"sukses": 0, "pesan": "Parameter bukan Angka"})
			//log.Fatal(err.Error())
		} else {
			c.JSON(200, gin.H{"sukses": 1, "data": models.DetailUser(id)})
		}

	}
}

func HapusUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(200, gin.H{"sukses": 0, "pesan": "Parameter bukan Angka"})
			//log.Fatal(err.Error())
		} else {
			models.HapusData(id)
			c.JSON(200, gin.H{"sukses": 1, "pesan": "Data dengan ID " + strconv.Itoa(id) + " telah dihapus"})
		}

	}
}

func UbahUser() gin.HandlerFunc {
	return func(request *gin.Context) {

		id, err := strconv.Atoi(request.Param("id"))
		if err != nil {
			request.JSON(200, gin.H{"sukses": 0, "pesan": "Parameter bukan Angka"})
			//log.Fatal(err.Error())
		} else {
			var pesanValidasi []string
			var namaFoto string
			nama := request.PostForm("nama")
			telp := request.PostForm("telp")
			jenisKelamin := request.PostForm("kelamin")
			pekerjaan := request.PostForm("pekerjaan")
			hobi := request.PostFormArray("hobi")
			implodeHobi := strings.Join(hobi, ";")
			// fmt.Println("Hobi yang dipilih:", implodeHobi)
			tanggalLahir := request.PostForm("tanggal")
			warnaFavorit := request.PostForm("warna")
			alamat := request.PostForm("alamat")
			foto, errFoto := request.FormFile("foto")

			if errFoto != nil {
				// log.Fatal(errFoto.Error())
				namaFoto = ""
			} else {
				mimeType := foto.Header.Get("Content-Type")
				if mimeType != "image/jpeg" && mimeType != "image/png" {
					pesanValidasi = append(pesanValidasi, "Foto Harus berupa file gambar jpg atau png")
				} else {
					namaFoto = foto.Filename
					dst := "img/" + namaFoto
					errUpload := request.SaveUploadedFile(foto, dst)
					if errUpload != nil {
						log.Fatal(errUpload.Error())
						return
					}
				}

			}

			if nama == "" {
				pesanValidasi = append(pesanValidasi, "Nama harus diisi")
			}

			if telp == "" || !regexp.MustCompile(`\d`).MatchString(telp) || len(telp) < 8 || len(telp) > 13 {
				pesanValidasi = append(pesanValidasi, "Telp harus diisi antara 8 s/d 13 angka")
			}

			if jenisKelamin == "" {
				pesanValidasi = append(pesanValidasi, "Jenis kelamin harus dipilih")
			}

			if pekerjaan == "" {
				pesanValidasi = append(pesanValidasi, "Pekerjaan harus dipilih")
			}

			if len(hobi) == 0 {
				pesanValidasi = append(pesanValidasi, "Hobi harus dipilih")
			}

			if tanggalLahir == "" {
				pesanValidasi = append(pesanValidasi, "Tanggal lahir harus dipilih")
			}

			if warnaFavorit == "" {
				pesanValidasi = append(pesanValidasi, "Warna favorit harus dipilih")
			}

			if alamat == "" {
				pesanValidasi = append(pesanValidasi, "Alamat harus diisi")
			}

			if len(pesanValidasi) > 0 {
				request.JSON(200, gin.H{"sukses": 0, "pesan": pesanValidasi})
			} else {

				hasilTanggal, err := time.Parse("2006-01-02", tanggalLahir)
				if err != nil {
					log.Fatal(err.Error())
				}

				models.UbahUser(
					id,
					nama,
					telp,
					jenisKelamin,
					pekerjaan,
					implodeHobi,
					warnaFavorit,
					hasilTanggal,
					alamat,
					namaFoto,
				)

				request.JSON(200, gin.H{"sukses": 1, "pesan": "Data berhasil diubah"})
			}
		}

	}
}
