package main
import "fmt"

type Akun struct {
	NamaLayanan   string
	Email         string
	KataSandi     string
	TanggalUpdate string
}

var DataAkun [100]Akun
var totalAkun int = 0

func main() {
	var pilihan int
	for {
		fmt.Println("\n=== Aplikasi Pengelola Kata Sandi Pribadi (SecurePass) ===")
		fmt.Println("1. Tambah Akun")
		fmt.Println("2. Tampilkan Semua Akun ")
		fmt.Println("3. Cari Akun ")
		fmt.Println("4. Urutkan Akun Alfabetis ")
		fmt.Println("5. Hapus Akun")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahAkun()
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func TambahAkun() {
	if totalAkun >= 100 {
		fmt.Println("Kapasitas penyimpanan penuh!")
		return
	}

	var layanan, email, sandi string
	fmt.Print("Masukkan Nama Layanan (contoh: Google): ")
	fmt.Scan(&layanan)
	fmt.Print("Masukkan Email/Username: ")
	fmt.Scan(&email)
	fmt.Print("Masukkan Kata Sandi: ")
	fmt.Scan(&sandi)
}
