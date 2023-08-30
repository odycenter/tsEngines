package tsArray

// 兩個陣列的交集
func IntersectArray(firstArr []string, secondArr []string) []string {
	var inter []string
	mp := make(map[string]bool)

	for _, s := range firstArr {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range secondArr {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}
