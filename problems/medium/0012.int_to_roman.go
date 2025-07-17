package medium

func IntToRoman(num int) string {
    mp := map[int]string{
        1:    "I",
        4:    "IV",
        5:    "V",
        9:    "IX",
        10:   "X",
        40:   "XL",
        50:   "L",
        90:   "XC",
        100:  "C",
        400:  "CD",
        500:  "D",
        900:  "CM",
        1000: "M",
    }
    if val, ok := mp[num]; ok {
        return val
    }
    
    res := ""
    cnt := 1
    for num > 0 {
        rem := num % 10
		rom := findRoman(mp, rem, cnt)
        res = rom + res
        num /= 10
        cnt *= 10
    }
    return res
}

func findRoman(mp map[int]string, rem int, cnt int) string {
	if val, ok := mp[rem * cnt]; ok {
		return val
	}
	var res string
	if rem < 4 {
		for range rem {
			res += mp[cnt]
		}
	} else {
		res += mp[5 * cnt]
		for range rem - 5 {
			res += mp[cnt]
		}
	}
	return res
}