package main

import (
	"fmt"
	"os"
)

// Konstanta
const appName = "Aplikasi CLI Sederhana"

// Global Variables
var userData = make(map[string]string) // Mengubah tipe data untuk email

// Fungsi utama untuk menampilkan menu
func main() {
	for {
		displayMainMenu()
	}
}

// Menampilkan menu utama
func displayMainMenu() {
	fmt.Println(appName)
	fmt.Println("\n--- Menu Utama ---")
	fmt.Println("1. Tampilkan Hello, World!")
	fmt.Println("2. Operasi Matematika")
	fmt.Println("3. Simpan dan Tampilkan Data Pengguna")
	fmt.Println("4. Hitung Faktorial")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih opsi (1-5): ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		printHelloWorld()
	case 2:
		menuOperasiMatematika()
	case 3:
		menuDataPengguna()
	case 4:
		hitungFaktorial()
	case 5:
		fmt.Println("Terima kasih telah menggunakan aplikasi!")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
	}
}

// 1. Fungsi untuk menampilkan Hello, World!
func printHelloWorld() {
	fmt.Println("\nHello, World!")
}

// 2. Menu untuk operasi matematika
func menuOperasiMatematika() {
	for {
		displayMathMenu()
		// Keluar dari menu setelah selesai
		return
	}
}

// Menampilkan menu operasi matematika
func displayMathMenu() {
	fmt.Println("\n--- Operasi Matematika ---")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Kembali ke menu utama")
	fmt.Print("Pilih operasi (1-5): ")

	var choice int
	fmt.Scan(&choice)

	if choice == 5 {
		// Mengembalikan ke menu utama
		return
	}

	// Meminta input angka dari pengguna
	var num1, num2 float64
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&num1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&num2)

	// Menggunakan switch untuk memilih operasi
	switch choice {
	case 1:
		fmt.Printf("Hasil penjumlahan: %.2f\n", penjumlahan(num1, num2))
	case 2:
		fmt.Printf("Hasil pengurangan: %.2f\n", pengurangan(num1, num2))
	case 3:
		fmt.Printf("Hasil perkalian: %.2f\n", perkalian(num1, num2))
	case 4:
		if num2 != 0 {
			fmt.Printf("Hasil pembagian: %.2f\n", pembagian(num1, num2))
		} else {
			fmt.Println("Error: Tidak bisa membagi dengan nol!")
		}
	default:
		fmt.Println("Operasi tidak valid, coba lagi.")
	}
}

// Fungsi untuk penjumlahan
func penjumlahan(a, b float64) float64 {
	return a + b
}

// Fungsi untuk pengurangan
func pengurangan(a, b float64) float64 {
	return a - b
}

// Fungsi untuk perkalian
func perkalian(a, b float64) float64 {
	return a * b
}

// Fungsi untuk pembagian
func pembagian(a, b float64) float64 {
	return a / b
}

// 3. Menu untuk menyimpan dan menampilkan data pengguna
func menuDataPengguna() {
	for {
		displayUserMenu()
		// Keluar dari menu setelah selesai
		return
	}
}

// Menampilkan menu data pengguna
func displayUserMenu() {
	fmt.Println("\n--- Menu Data Pengguna ---")
	fmt.Println("1. Tambahkan Data Pengguna")
	fmt.Println("2. Tampilkan Semua Data Pengguna")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih opsi (1-3): ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addUserData()
	case 2:
		displayUserData()
	case 3:
		// Mengembalikan kontrol ke menu utama
		return
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
	}
}

// Fungsi untuk menambahkan data pengguna
func addUserData() {
	var name, email string

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&name)

	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)

	userData[name] = email
	fmt.Println("Data berhasil ditambahkan!")

	displayUserMenu()
}

// Fungsi untuk menampilkan data pengguna
func displayUserData() {
	if len(userData) == 0 {
		fmt.Println("Tidak ada data pengguna.")
		return
	}

	fmt.Println("\nData Pengguna:")
	for name, email := range userData {
		fmt.Printf("Nama: %s\nEmail: %s\n", name, email)
	}
	displayUserMenu()
}

// 4. Fungsi untuk menghitung faktorial menggunakan rekursi
func hitungFaktorial() {
	var num int
	fmt.Print("Masukkan angka untuk menghitung faktorial: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Faktorial tidak bisa dihitung untuk bilangan negatif.")
	} else {
		fmt.Printf("Faktorial dari %d adalah: %d\n", num, faktorial(num))
	}
}

// Fungsi rekursif untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}