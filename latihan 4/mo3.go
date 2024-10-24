package main

import (
	"fmt"
	"math"
)

func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Println("Masukkan koordinat titik A (x y) :")
	fmt.Scan(&ax, &ay)
	fmt.Println("Masukkan koordinat titik B (x y) :")
	fmt.Scan(&bx, &by)
	fmt.Println("Masukkan koordinat titik C (x y) :")
	fmt.Scan(&cx, &cy)

	sisiAB := hitungJarak(ax, ay, bx, by)
	sisiBC := hitungJarak(bx, by, cx, cy)
	sisiCA := hitungJarak(cx, cy, ax, ay)

	sisiTerpanjang := sisiAB
	if sisiBC > sisiTerpanjang {
		sisiTerpanjang = sisiBC
	}
	if sisiCA > sisiTerpanjang {
		sisiTerpanjang = sisiCA
	}

	fmt.Printf("Panjang sisi terpanjang dari segitiga adalah : %.2f\n", sisiTerpanjang)
}
