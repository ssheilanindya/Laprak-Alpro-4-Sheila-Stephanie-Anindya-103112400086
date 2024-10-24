package main

import "fmt"

func main() {
	var diskon, totalBelanja, totalAkhir, diskonPersen int

	fmt.Print("Masukkan total belanja awal : ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Masukkan besaran diskon (dalam persen): ")
	fmt.Scan(&diskonPersen)

	diskon = (totalBelanja * diskonPersen) / 100
	totalAkhir = totalBelanja - diskon

	fmt.Printf("Total belanja setelah diskon: %d\n", totalAkhir)
}
