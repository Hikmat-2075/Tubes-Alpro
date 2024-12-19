package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Skripsi struct {
	aktifitas        string
	waktu, prioritas int
	progres          float64
}

const NMAX int = 15

type TabSkripsi [NMAX]Skripsi

var arrSkripsi TabSkripsi
var n int

func main() {
	menuAwal()
}

func menuAwal() {
	// I.S.: Program diinisialisasi dan menunggu input dari pengguna.
	// F.S.: Pengguna memasukkan aktivitas, atau keluar dari program.
	var pilih int
	bersihLayar()
	fmt.Println("==============================")
	fmt.Println("        Selamat Datang        ")
	fmt.Println("==============================")
	fmt.Println("  Silahkan masukan aktivitas   ")
	fmt.Println("==============================")
	fmt.Println("1. Masukan aktivitas")
	fmt.Println("2. Exit")
	fmt.Print("pilih(1/2): ")
	fmt.Scan(&pilih)
	for {
		if pilih == 1 {
			InputData(&arrSkripsi, &n)
			bersihLayar()
			menu()
			return
		} else if pilih == 2 {
			bersihLayar()
			selesai()
			return
		} else {
			fmt.Println("angka yang anda masukan tidak valid")
			fmt.Scan(&pilih)
		}
	}
}

func bersihLayar() {
	// I.S.: Konsol berisi outputan dari code sebelumnya dan perlu dibersihkan.
	// F.S.: Layar konsol telah dibersihkan.
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func menu() {
	var pilih int
	for {
	// I.S.: Pengguna telah memasukkan aktivitas dan berada di menu utama menunggu tindakan lebih lanjut.
	// F.S.: Pengguna telah memilih untuk melihat, menambah, mengedit, menghapus aktivitas, atau keluar ke menu awal.
		bersihLayar()
		fmt.Println("==========================")
		fmt.Println("      PILIHAN MENU        ")
		fmt.Println("==========================")
		fmt.Println("1. MENAMPILKAN AKTIFITAS")
		fmt.Println("2. Menambahkan Aktifitas ")
		fmt.Println("3. Mengedit Aktifitas")
		fmt.Println("4. Menghapus Aktifitas")
		fmt.Println("5. Progres Skripsi")
		fmt.Println("6. Exit")
		fmt.Println("========================= ")
		fmt.Print("Pilih (1/2/3/4/5/6): ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			bersihLayar()
			menuTampilan()
		case 2:
			bersihLayar()
			tambahData(&arrSkripsi, &n)
		case 3:
			bersihLayar()
			mengeditData(&arrSkripsi, n)
		case 4:
			bersihLayar()
			var p string
			fmt.Print("Aktivitas yang ingin di hapus: ")
			fmt.Scan(&p)
			hapusData(&arrSkripsi, &n, p)
		case 5:
			bersihLayar()
			progresSkripsi(arrSkripsi, n)
		case 6:
			bersihLayar()
			selesai()
			return
		default:
			fmt.Println("angka yang anda masukan tidak valid")
		}
	}
}

func InputData(A *TabSkripsi, n *int) {
	// I.S.: Array TabSkripsi mungkin kosong atau sebagian terisi, dan n menunjukkan jumlah aktivitas saat ini.
	// F.S.: Array TabSkripsi diperbarui dengan aktivitas baru yang dimasukkan oleh pengguna, dan n diperbarui untuk mencerminkan jumlah aktivitas baru.
	*n = 0
	var x, p int
	var s string
	i := 0
	fmt.Println("=======================")
	fmt.Println(" Menambahkan Data")
	fmt.Println("=======================")
	fmt.Print("Aktivitas ke-1: ")
	fmt.Scan(&s)

	for s != "none" && i < NMAX {
		fmt.Print("Perkiraan waktu pengerjaan (dalam hari): ")
		fmt.Scan(&x)
		fmt.Print("Prioritas: ")
		fmt.Scan(&p)
		fmt.Println("=======================")
		A[*n].aktifitas = s
		A[*n].waktu = x
		A[*n].prioritas = p
		*n++
		fmt.Printf("Aktivitas ke-%d: ", *n+1)
		fmt.Scan(&s)
		i++
	}
}

func TampilData(A TabSkripsi, n int) {
	// I.S.: Array TabSkripsi berisi aktivitas, dan n menunjukkan jumlah aktivitas.
	// F.S.: Aktivitas dalam TabSkripsi ditampilkan di konsol.
	fmt.Println("======================================")
	fmt.Println("            Tampilan Data             ")
	fmt.Println("======================================")
	fmt.Printf("%-25s %-15s %s\n", "Aktivitas", "Waktu Pengerjaan", "Prioritas")
	for i := 0; i < n; i++ {
		fmt.Printf("%-25s %-15d %d\n", A[i].aktifitas, A[i].waktu, A[i].prioritas)
	}
	fmt.Println("======================================")
	totalWaktu(arrSkripsi, n)
	fmt.Println("======================================")
	var pilih int
	fmt.Print("<1.kembali ke menu> : ")
	fmt.Scan(&pilih)
	for pilih != 1 {
		fmt.Println("Angka yang anda masukan tidak valid")
		fmt.Scan(&pilih)
	}
}

func tambahData(A *TabSkripsi, n *int) {
	// I.S.: Array TabSkripsi sebagian terisi, dan n menunjukkan jumlah aktivitas.
	// F.S.: Aktivitas baru ditambahkan ke TabSkripsi, dan n bertambah.
	if *n < NMAX {
		var s string
		var x, p int
		fmt.Printf("Aktivitas ke-%d: ", *n+1)
		fmt.Scan(&s)
		fmt.Print("Perkiraan waktu pengerjaan (dalam hari): ")
		fmt.Scan(&x)
		fmt.Print("Prioritas: ")
		fmt.Scan(&p)
		A[*n].aktifitas = s
		A[*n].waktu = x
		A[*n].prioritas = p
		*n++
		bersihLayar()
	} else {
		fmt.Println("Tidak bisa menambah aktivitas lagi.")
	}
}

func mengeditData(T *TabSkripsi, n int) {
	// I.S.: Array TabSkripsi berisi aktivitas, dan n menunjukkan jumlah aktivitas.
	// F.S.: Aktivitas yang ditentukan dalam TabSkripsi diperbarui berdasarkan input pengguna.
	var Aktivitas, opsi, newAktifitas, bagian string
	var pilih int
	for {
		bersihLayar()
		fmt.Println("=======================================")
		fmt.Println("            MENGUBAH DATA              ")
		fmt.Println("=======================================")
		fmt.Print("Aktivitas yang ingin diubah: ")
		fmt.Scan(&Aktivitas)
		idx := sequential_search(*T, n, Aktivitas)
		if idx != -1 {
			fmt.Printf("Aktivitas '%s' ditemukan.\n", Aktivitas)
			fmt.Println("1. Ubah semua data || 2. ubah sebagian data ")
			fmt.Print("Pilihan: ")
			fmt.Scan(&opsi)
			if opsi == "1" {
				fmt.Print("Inputkan aktifitas baru: ")
				fmt.Scan(&newAktifitas)
				T[idx].aktifitas = newAktifitas
				fmt.Print("Perkiraan waktu pengerjaan (dalam hari): ")
				fmt.Scan(&T[idx].waktu)
				fmt.Print("Prioritas: ")
				fmt.Scan(&T[idx].prioritas)
			} else if opsi == "2" {
				fmt.Print("Bagian yang mau diubah (Aktifitas/waktu/prioritas): ")
				fmt.Scan(&bagian)
				if bagian == "Aktifitas" {
					fmt.Print("Aktivitas baru: ")
					fmt.Scan(&T[idx].aktifitas)
				} else if bagian == "waktu" {
					fmt.Print("Waktu pengerjaan baru (dalam hari): ")
					fmt.Scan(&T[idx].waktu)
				} else if bagian == "prioritas" {
					fmt.Print("Prioritas baru: ")
					fmt.Scan(&T[idx].prioritas)
				} else {
					fmt.Println("Bagian yang anda masukkan tidak valid")
				}
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else {
			fmt.Println("Aktivitas tidak ditemukan.")
		}
		fmt.Println("<<0.Ubah data lain>>", "\t", "<<1.Kembali ke Menu>>")
		fmt.Scan(&pilih)
		if pilih == 1 {
			bersihLayar()
			return
		}
	}
}

func hapusData(A *TabSkripsi, n *int, aktifitas string) {
	// I.S.: terdefinisi Array TabSkripsi berisi aktivitas, dan n menunjukkan jumlah aktivitas.
	// F.S.: Aktivitas yang ditentukan dihapus dari TabSkripsi, dan n berkurang.
	idx := sequential_search(*A, *n, aktifitas)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func menuTampilan() {
	// I.S.: Pengguna berada di menu utama dan memilih untuk melihat aktivitas.
	// F.S.: Pengguna dapat melihat aktivitas dalam berbagai urutan yang diinginkan atau kembali ke menu utama.
	var pilih int
	for {
		bersihLayar()
		fmt.Println("======================")
		fmt.Println(" MENU MENAMPILKAN AKTIVITAS")
		fmt.Println("======================")
		fmt.Println("1. Menampilkan data")
		fmt.Println("2. Menampilkan berdasarkan prioritas")
		fmt.Println("3. Menampilkan berdasarkan waktu")
		fmt.Println("4. Kembali ke menu")
		fmt.Println("======================")
		fmt.Print("Pilih(1/2/3/4): ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			bersihLayar()
			TampilData(arrSkripsi, n)
		case 2:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("<1. Descending> || <2. Ascending>")
			fmt.Scan(&pilih)
			if pilih == 1{
				bersihLayar()
				urutMenurunPrioritas(&arrSkripsi, n)
				TampilData(arrSkripsi, n)
			}else if pilih == 2{
				bersihLayar()
				urutMenaikPrioritas(&arrSkripsi, n)
				TampilData(arrSkripsi, n)
			}
		case 3:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("<1. Descending> || <2. Ascending>")
			fmt.Scan(&pilih)
			if pilih == 1{
				bersihLayar()
				urutWaktuMenurun(&arrSkripsi, n)
				TampilData(arrSkripsi, n)
			}else if pilih == 2{
				bersihLayar()
				urutWaktuMenaik(&arrSkripsi, n)
				TampilData(arrSkripsi, n)
			}
		case 4:
			bersihLayar()
			return
		default:
			fmt.Println("angka yang anda masukan tidak valid")
		}
	}
}

func urutMenaikPrioritas(A *TabSkripsi, n int) {
	// I.S.: Array TabSkripsi berisi aktivitas dalam urutan acak, dan n menunjukkan jumlah aktivitas.
	// F.S.: Array TabSkripsi diurutkan dalam urutan menaik berdasarkan prioritas.
	var i, pass int
	var temp Skripsi
	for pass = 1; pass < n; pass++ {
		i = pass
		temp = A[i]
		for i > 0 && (A[i-1].prioritas > temp.prioritas || A[i-1].prioritas == temp.prioritas && A[i-1].waktu > temp.waktu) {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func urutMenurunPrioritas(A *TabSkripsi, n int) {
	// I.S.: Array TabSkripsi berisi aktivitas dalam urutan acak, dan n menunjukkan jumlah aktivitas.
	// F.S.: Array TabSkripsi diurutkan dalam urutan menurun berdasarkan prioritas.
	var pass, idx, j int
	var temp Skripsi
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for j = pass; j < n; j++ {
			if A[j].prioritas > A[idx].prioritas || A[j].prioritas == A[idx].prioritas && A[j].waktu > A[idx].waktu {
				idx = j
			}
		}
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
	}
}

func urutWaktuMenurun(A *TabSkripsi, n int){
	// I.S.: Array TabSkripsi berisi aktivitas dalam urutan acak, dan n menunjukkan jumlah aktivitas.
	// F.S.: Array TabSkripsi diurutkan dalam urutan menurun berdasarkan waktu.
	var pass, idx, j int
	var temp Skripsi
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for j = pass; j < n; j++ {
			if A[j].waktu > A[idx].waktu || A[j].prioritas == A[idx].prioritas && A[j].waktu > A[idx].waktu {
				idx = j
			}
		}
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
	}
}

func urutWaktuMenaik(A *TabSkripsi, n int){
	// I.S.: Array TabSkripsi berisi aktivitas dalam urutan acak, dan n menunjukkan jumlah aktivitas.
	// F.S.: Array TabSkripsi diurutkan dalam urutan menaik berdasarkan waktu.
	var i, pass int
	var temp Skripsi
	for pass = 1; pass < n; pass++ {
		i = pass
		temp = A[i]
		for i > 0 && (A[i-1].waktu > temp.waktu || A[i-1].prioritas == temp.prioritas && A[i-1].waktu > temp.waktu) {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func selesai() {
	// I.S.: Pengguna memilih untuk keluar dari program.
	// F.S.: Pesan terima kasih ditampilkan, dan program berakhir.
	fmt.Println("======================================")
	fmt.Println("          TERIMA KASIH                ")
	fmt.Println("======================================")
}

func sequential_search(T TabSkripsi, n int, k string) int {
	// I.S.: terdefinisi Array TabSkripsi berisi aktivitas, dan n menunjukkan jumlah aktivitas. Fungsi mencari aktivitas tertentu x.
	// F.S.: Fungsi menampilkan indeks dari aktivitas x jika ditemukan, atau -1 jika tidak ditemukan.
	for i := 0; i < n; i++ {
		if T[i].aktifitas == k {
			return i
		}
	}
	return -1
}

func progresSkripsi(A TabSkripsi, n int) {
	bersihLayar()
	fmt.Println("======================================")
	fmt.Println("         Progres Skripsi              ")
	fmt.Println("======================================")
	if n == 0 {
		fmt.Println("Tidak ada aktivitas yang dimasukkan.")
	} else {
		progres := (float64(n) / float64(NMAX)) * 100
		fmt.Printf("Anda telah menyelesaikan %.2f%% dari skripsi Anda.\n", progres)
	}
	fmt.Println("======================================")
	var pilih int
	fmt.Print("<1. Kembali ke menu> : ")
	fmt.Scan(&pilih)
	for pilih != 1 {
		fmt.Println("Angka yang anda masukan tidak valid")
		fmt.Scan(&pilih)
	}
}

func totalWaktu(A TabSkripsi, n int) {
	// I.S.: Array TabSkripsi berisi aktivitas, dan n menunjukkan jumlah aktivitas.
	// F.S.: Total waktu yang diperlukan untuk semua aktivitas dihitung dan dikembalikan.
	var bulan, hari int
	for i := 0; i < n; i++ {
		hari = hari + A[i].waktu
	}
	if hari >= 30 {
		bulan = hari / 30
		hari = hari % 30
		fmt.Println("deadline Waktu :", bulan, "Bulan ", hari, "Hari")
	} else {
		fmt.Println("deadline Waktu :", hari, "Hari")
	}
}

