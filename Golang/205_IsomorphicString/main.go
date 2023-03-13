package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if s == "" || t == "" || len(s) != len(t) {
		return false
	}

	var s_slice []int = make([]int, 256)
	var t_slice []int = make([]int, 256)

	for idx, _ := range t {
		if s_slice[int(s[idx])] != t_slice[int(t[idx])] {
			return false
		}

		s_slice[int(s[idx])] = idx + 1
		t_slice[int(t[idx])] = idx + 1
	}

	return true
}

func isIsomorphic2(s string, t string) bool {
	ms := make(map[byte]byte, len(s))
	mt := make(map[byte]byte, len(t))

	for i := 0; i < len(s); i++ {
		sByte := s[i]
		tByte := t[i]

		valS, okS := ms[sByte]
		valT, okT := mt[tByte]

		if !okS && !okT {
			ms[sByte] = tByte
			mt[tByte] = sByte
			continue
		}

		if valS != tByte || valT != sByte {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}
