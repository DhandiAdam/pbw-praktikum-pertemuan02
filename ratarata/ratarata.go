package ratarata

func HitungRataRata(nilaiSiswa []int) float64 {
	total := 0
	for _, nilai := range nilaiSiswa {
		total += nilai
	}

	rataRata := float64(total) / float64(len(nilaiSiswa))
	return rataRata
}
