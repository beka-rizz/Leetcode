package easy

var mp = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func RomanToInt(s string) int {
	var result int = 0
	for id := 0; id < len(s); id++ {
		if id+1 < len(s) {
			currentCh := mp[string(s[id])]
			nextCh := mp[string(s[id+1])]
			if nextCh > currentCh {
				result += mp[string(s[id])+string(s[id+1])]
				id++
				continue
			}
		}
		result += mp[string(s[id])]
	}
	return result
}
