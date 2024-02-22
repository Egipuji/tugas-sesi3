package main

import (
	"errors"
	"fmt"
	"os"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func findData(id string, data map[string]Biodata) (Biodata, error) {
	//melakukan perulang untuk mencari apakah ada key yg sama dengan id
	for key, value := range data {
		if key == id {
			//mereturn value dari key yang sama dan return error menjadi nil
			return value, nil
		}
	}
	//bila key tidak ditemukan mereturn value kosong/ struct kosong dan error message
	return Biodata{}, errors.New("data siswa tidak ada")
}

func main() {
	//mendapatkan argumen dari terminal
	args := os.Args[1:]

	siswaMap := map[string]Biodata{
		"1": {Nama: "Egi Puji Sutrisno", Alamat: "Kota Tangerang", Pekerjaan: "Ingin menjadi backend programmer", Alasan: "Ingin mengasah skill golang"},
		"2": {Nama: "Ilham", Alamat: "Jakarta", Pekerjaan: "Presiden", Alasan: "Mengasah skill"},
		"3": {Nama: "Kurniawan", Alamat: "Bandung", Pekerjaan: "CEO", Alasan: "Ingin tau aja"},
		"4": {Nama: "Kurniadi", Alamat: "Surabaya", Pekerjaan: "Pengusaha", Alasan: "Ingin kaya"},
	}

	data, err := findData(args[0], siswaMap)

	//cek apakah ada error
	if err != nil {
		panic(err)
	}

	fmt.Println("Data siswa dengan absen no", args[0])
	fmt.Println("Nama:", data.Nama)
	fmt.Println("Alamat:", data.Alamat)
	fmt.Println("Pekerjaan:", data.Pekerjaan)
	fmt.Println("Alasan:", data.Alasan)

}
