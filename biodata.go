package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func findData(absen int, data map[int]Biodata) (*Biodata, error) {
	//mencari siswa berdasarkan no absen
	if value, ok := data[absen]; ok {
		return &value, nil
	}
	//bila key tidak ditemukan mereturn value kosong/ struct kosong dan error message
	return nil, errors.New("data siswa tidak ada")
}

func main() {
	//mendapatkan argumen dari terminal
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Harap masukan nomer absen")
		return
	}

	absen, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Masukan absen berupa number")
		return
	}

	siswaMap := map[int]Biodata{
		1: {Nama: "Egi Puji Sutrisno", Alamat: "Kota Tangerang", Pekerjaan: "Ingin menjadi backend programmer", Alasan: "Ingin mengasah skill golang"},
		2: {Nama: "Ilham", Alamat: "Jakarta", Pekerjaan: "Presiden", Alasan: "Mengasah skill"},
		3: {Nama: "Kurniawan", Alamat: "Bandung", Pekerjaan: "CEO", Alasan: "Ingin tau aja"},
		4: {Nama: "Kurniadi", Alamat: "Surabaya", Pekerjaan: "Pengusaha", Alasan: "Ingin kaya"},
	}

	data, err := findData(absen, siswaMap)

	//mengecek apakah find data mereturn error atau error nil
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data siswa dengan absen no", absen)
	fmt.Println("Nama:", data.Nama)
	fmt.Println("Alamat:", data.Alamat)
	fmt.Println("Pekerjaan:", data.Pekerjaan)
	fmt.Println("Alasan:", data.Alasan)

}
