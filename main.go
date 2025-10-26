package main

import "fmt"

func main() {
	datas := [...]int{1, 4, 7, 8, 9, 10, 15, 17, 27, 30, 31, 34, 35, 37, 29, 40, 50, 55, 56, 57, 59, 60, 70, 75, 85}

	cari := 55

	awal := 0
	akhir := len(datas) - 1

	for {
		var tengah int = (awal + akhir) / 2

		// menghentikan pencarian ketika data tidak ditemukan
		if awal > akhir {
			fmt.Println("Data tidak ditemukan")
			break
		}

		if cari == datas[tengah] {

			// Print data ketika ditemukan
			fmt.Println("Data :", cari, "Berada pada Index ke :", tengah)
			break

		} else if cari > datas[tengah] {

			awal = tengah + 1

		} else if cari < datas[tengah] {

			akhir = tengah - 1

		}
	}
}
