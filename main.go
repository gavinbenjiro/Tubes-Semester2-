package main

import (
	"fmt"
)

const NMAX int = 15

type lapangan struct {
	namaLapangan, noTelp string
	hargaTotal int
}

type olahraga [NMAX]lapangan

func main() {
	var badminton olahraga

	LapBadminton(&badminton)
	Menu(badminton)
}

func Pemesanan(namaUser, noTelpUser, lapanganDipesan *string, badminton olahraga) {
	var noTelpAdmin string
	var n int

	n = NMAX
	fmt.Println()
	fmt.Println("Isi nama & no telepon anda: ")
	fmt.Print("Nama: ")
	fmt.Scan(namaUser)
	fmt.Print("No Telp: ")
	fmt.Scan(noTelpUser)
	fmt.Print("Lapangan yang ingin dipesan: ")
	fmt.Scan(lapanganDipesan)
	for !CekLapangan(*lapanganDipesan, badminton) {
		fmt.Println("Lapangan tidak tersedia, silahkan input kembali")
		fmt.Print("Lapangan yang ingin dipesan: ")
		fmt.Scan(lapanganDipesan)
	}
	fmt.Printf("User atas nama %s dapat menghubungi admin %s dengan nomor berikut: \n", *namaUser, *lapanganDipesan)
	for i := 0; i < n; i++ {
		if *lapanganDipesan == badminton[i].namaLapangan {
			noTelpAdmin = badminton[i].noTelp
		}
	}
	fmt.Print(noTelpAdmin)
}

func CekLapangan(namaLap string, badminton olahraga) bool {
	var status bool
	var n int

	status = false
	n = NMAX
	for i := 0; i < n; i++ {
		if namaLap == badminton[i].namaLapangan {
			status = true
		}
	}
	return status
}

func Menu(badminton olahraga) {
	var Yellow = "\033[33m" 
	var Reset = "\033[0m" 
	var inputUser int
	var namaUser, noTelpUser, lapanganDipesan string

	fmt.Println(Yellow + "----------------------------------")
	fmt.Println("Aplikasi Sortir Lapangan Badminton")
	fmt.Println("----------------------------------" + Reset)
	fmt.Print("Berikut daftar lapangan badminton yang dapat anda lihat: ")
	fmt.Println()
	CetakLapBadminton(badminton)
	fmt.Println("1. Lakukan filter lapangan")
	fmt.Println("2. Lakukan pemesanan lapangan")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih: ")
	for inputUser < 1 || inputUser > 3 {
		fmt.Scan(&inputUser)
		switch inputUser {
		case 1:
			FilterLapangan(badminton)
		case 2:
			Pemesanan(&namaUser, &noTelpUser, &lapanganDipesan, badminton)
		case 3:
			Keluar()
		}
		if inputUser < 1 || inputUser > 3 {
			fmt.Print("Input invalid, silahkan input kembali: ")
		}
	}
}

func Keluar() {
	fmt.Println()
	fmt.Println("Terimakasih Telah Mengunjungi Aplikasi Kami.")
	fmt.Println("Semoga Harimu Menyenangkan!!!")
}

func LapBadminton(badminton *olahraga) {
	badminton[0].namaLapangan = "GOR_Iniro"
	badminton[1].namaLapangan = "Indraprasta_Hall"
	badminton[2].namaLapangan = "GOR_Sawojajar"
	badminton[3].namaLapangan = "Bintang_Badminton"
	badminton[4].namaLapangan = "Kaypi_GOR"
	badminton[5].namaLapangan = "Wins_Arena"
	badminton[6].namaLapangan = "GOR_Moskow"
	badminton[7].namaLapangan = "GOR_Ong"
	badminton[8].namaLapangan = "Heptaseas_Badminton"
	badminton[9].namaLapangan = "Microsoft_Badminton"
	badminton[10].namaLapangan = "Badminton_Republic"
	badminton[11].namaLapangan = "GOR_Combi"
	badminton[12].namaLapangan = "Choppit_Arena"
	badminton[13].namaLapangan = "Moris_Arena"
	badminton[14].namaLapangan = "Tel-U_Hall"

	badminton[0].noTelp = "0821-4567-3435"
	badminton[1].noTelp = "0821-7459-3923"
	badminton[2].noTelp = "0812-2374-2332"
	badminton[3].noTelp = "0212-1135-7587"
	badminton[4].noTelp = "0821-1467-4866"
	badminton[5].noTelp = "0877-6496-0730"
	badminton[6].noTelp = "0891-7834-2847"
	badminton[7].noTelp = "0712-4394-1834"
	badminton[8].noTelp = "0832-5335-3434"
	badminton[9].noTelp = "0822-2505-2024"
	badminton[10].noTelp = "0179-4569-2838"
	badminton[11].noTelp = "7484-2947-2923"
	badminton[12].noTelp = "0834-1476-2135"
	badminton[13].noTelp = "8275-2940-2927"
	badminton[14].noTelp = "0483-5738-3858"
	
	badminton[0].hargaTotal = 60000
	badminton[1].hargaTotal = 75000
	badminton[2].hargaTotal = 70000
	badminton[3].hargaTotal = 100000
	badminton[4].hargaTotal = 80000 
	badminton[5].hargaTotal = 85000
	badminton[6].hargaTotal = 50000
	badminton[7].hargaTotal = 90000
	badminton[8].hargaTotal = 190000
	badminton[9].hargaTotal = 130000
	badminton[10].hargaTotal = 95000
	badminton[11].hargaTotal = 55000
	badminton[12].hargaTotal = 65000
	badminton[13].hargaTotal = 81000
	badminton[14].hargaTotal = 120000
}

func CetakLapBadminton(badminton olahraga) {
	fmt.Println()
	fmt.Println("Daftar lapangan yang tersedia: ")
	for i := 0; i < NMAX; i++ {
		fmt.Printf("%s | Rp.%d \n", badminton[i].namaLapangan, badminton[i].hargaTotal)
	}
	fmt.Println("-----------------------------")
}

func FilterLapangan(badminton olahraga) {
	var inputUser int
	var inputSetelahFilter, startLoop int
	var berhenti bool
	var namaUser, noTelpUser, lapanganDipesan string
	fmt.Println()
	fmt.Println("Bagaimana anda ingin melakukan filter: ")
	fmt.Println("1. Harga terendah hingga tertinggi")
	fmt.Println("2. Harga tertinggi hingga terendah")
	fmt.Println("3. Kisaran harga")
	fmt.Println("4. Harga spesifik")
	fmt.Println("5. Kembali")
	fmt.Println("6. Keluar")
	fmt.Print("pilih: ")
	for inputUser < 1 || inputUser > 6 {
		fmt.Scan(&inputUser)
		switch inputUser {
		case 1: 
			SelectionAscending(&badminton)
			CetakLapBadminton(badminton)
		case 2:
			InsertionDescending(&badminton)
			CetakLapBadminton(badminton)
		case 3:
			SequentialSearch(&badminton)
		case 4:
			BinarySearch(&badminton)
		case 5: 
			Menu(badminton)
			berhenti = true
		case 6: 
			Keluar()
			berhenti = true
		}
		if inputUser < 1 || inputUser > 6 {
			fmt.Print("Input invalid, silahkan input kembali: ")
		}
	}
	if !berhenti {
		for inputSetelahFilter < 1 || inputSetelahFilter > 2 {
			if inputSetelahFilter == 1 || inputSetelahFilter == 2 || startLoop == 0 {
				fmt.Println("1. Lakukan pemesanan lapangan")
				fmt.Println("2. Kembali")
				fmt.Print("Pilih: ")
				startLoop = 1
			}
			fmt.Scan(&inputSetelahFilter)
			switch inputSetelahFilter {
			case 1:
				Pemesanan(&namaUser, &noTelpUser, &lapanganDipesan, badminton)
			case 2:
				FilterLapangan(badminton)
			}
			if inputSetelahFilter < 1 || inputSetelahFilter > 2 {
				fmt.Print("Input invalid, silahkan input kembali: ")
			}
		}
	}
}

func BinarySearch(badminton *olahraga) {
	var left, right, mid, n, harga int
	var tersedia bool
	
	n = NMAX
	tersedia = false
	SelectionAscending(&*badminton)
	fmt.Println()
	for !tersedia {
		fmt.Print("Masukkan harga: ")
		fmt.Scan(&harga)
		fmt.Println("---------------")
		left = 0
		right = n-1
		mid = (left + right) / 2
		for left <= right && badminton[mid].hargaTotal != harga {
			if harga < badminton[mid].hargaTotal {
				right = mid - 1
			} else {
				left = mid + 1
			}
			mid = (left + right) / 2
		}
		tersedia = mid >= 0 && badminton[mid].hargaTotal == harga
		if !tersedia {
			fmt.Printf("Lapangan dengan harga %d tidak tersedia, input kembali. \n", harga)
		} else {
			fmt.Println("Daftar lapangan yang tersedia: ")
			fmt.Printf("%s | Rp.%d \n", badminton[mid].namaLapangan, badminton[mid].hargaTotal)
			fmt.Println("---------------")
		}
	}
} 

func SequentialSearch(badminton *olahraga) {
	var n, hargaMinimal, hargaMaksimal int
	var tersedia bool
	
	n = NMAX
	tersedia = false
	fmt.Println()
	for !tersedia {
		fmt.Print("Masukkan harga minimal: ")
		fmt.Scan(&hargaMinimal)
		fmt.Print("Masukkan harga maksimal: ")
		fmt.Scan(&hargaMaksimal)
		fmt.Println("------------------------")
		fmt.Println("Daftar lapangan yang tersedia: ")
		for i := 0; i < n; i++ {
			if badminton[i].hargaTotal >= hargaMinimal && badminton[i].hargaTotal <= hargaMaksimal {
				fmt.Printf("%s | Rp.%d \n", badminton[i].namaLapangan, badminton[i].hargaTotal)
				tersedia = true
			}
		}
		if !tersedia {
			fmt.Printf("Lapangan dengan harga %d hingga %d tidak tersedia, input kembali. \n", hargaMinimal, hargaMaksimal)
		} else {
			fmt.Println("------------------------")
		}
	}
}

func InsertionDescending(badminton *olahraga) {
	var pass, i, tempTotal, n int
	var tempNama,tempNomor string

	n = NMAX
	pass = 1
	for pass <= n-1 {
		i = pass
		tempTotal = badminton[pass].hargaTotal
		tempNama = badminton[pass].namaLapangan
		tempNomor = badminton[pass].noTelp
		for i > 0 && tempTotal > badminton[i-1].hargaTotal {
			badminton[i].hargaTotal = badminton[i-1].hargaTotal
			badminton[i].namaLapangan = badminton[i-1].namaLapangan
			badminton[i].noTelp = badminton[i-1].noTelp
			i--
		}
		badminton[i].hargaTotal = tempTotal
		badminton[i].namaLapangan = tempNama
		badminton[i].noTelp = tempNomor
		pass++
	}
}

func SelectionAscending(badminton *olahraga) {
	var index, n int

	n = NMAX
	for n > 0 {
		index = FindMax(*badminton, n)
		Swap(&*badminton, index, n)
		n--
	}
}


func FindMax(badminton olahraga, n int) int {
	var max, index int

	max = badminton[0].hargaTotal
	for i := 1; i < n; i++ {
		if max < badminton[i].hargaTotal {
			max = badminton[i].hargaTotal
			index = i
		}
	}
	return index
}

func Swap(badminton *olahraga, idx, n int) {
	var tempTotal int
	var tempNama, tempNomor string

	tempNama = badminton[n-1].namaLapangan
	badminton[n-1].namaLapangan = badminton[idx].namaLapangan
	badminton[idx].namaLapangan = tempNama

	tempTotal = badminton[n-1].hargaTotal
	badminton[n-1].hargaTotal = badminton[idx].hargaTotal
	badminton[idx].hargaTotal = tempTotal

	tempNomor = badminton[n-1].noTelp
	badminton[n-1].noTelp = badminton[idx].noTelp
	badminton[idx].noTelp = tempNomor
}