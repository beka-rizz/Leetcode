package easy

func IsAnagram(s string, t string) bool {
    mp := make(map[rune]int)
	// saving occurences of runes in s to the map
    for _, b := range s {
        mp[b]++
    }

	// checking that all runes of t exists in s
    for _, b := range t {
		if _, ok := mp[b]; ok && mp[b] != 0 {
			mp[b]--
		} else {
			return false
		}
    }
	// checking that all runes of s used in t
	for r := range mp {
		if (mp[r] != 0) {
			return false
		}
	}
	return true
}