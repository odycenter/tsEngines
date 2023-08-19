package tsSlice

import "strconv"

// https://play.golang.org/p/Qg_uv_inCek
// IntegerContains contains checks if a integer is present in a slice
func IntegerContains(elems []int64, v int64) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func ToInt64Slice(numStrings []string) (rtnNums []int64) {
	for _, v := range numStrings {
		id, _ := strconv.ParseInt(v, 10, 64)
		rtnNums = append(rtnNums, id)
	}
	return
}

func ToStringNumSlice(int64s []int64) (rtnStrings []string) {
	for _, v := range int64s {
		number := strconv.FormatInt(v, 10)
		rtnStrings = append(rtnStrings, number)
	}
	return rtnStrings
}

func StringContains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// RemoveRepInt64ByLoop 通过两重循环过滤重复元素
func RemoveRepInt64ByLoop(slc []int64) []int64 {
	var result []int64 // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// RemoveRepInt64ByMap 通过map主键唯一的特性过滤重复元素
func RemoveRepInt64ByMap(slc []int64) []int64 {
	var result []int64
	tempMap := map[int64]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// RemoveRepInt64 slice int64 切片去重
func RemoveRepInt64(slc []int64) []int64 {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return RemoveRepInt64ByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return RemoveRepInt64ByMap(slc)
	}
}

// RemoveRepStringByLoop 通过两重循环过滤重复元素
func RemoveRepStringByLoop(slc []string) []string {
	var result []string // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// RemoveRepStringByMap 通过map主键唯一的特性过滤重复元素
func RemoveRepStringByMap(slc []string) []string {
	var result []string
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// RemoveRepString slice String 切片去重
func RemoveRepString(slc []string) []string {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return RemoveRepStringByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return RemoveRepStringByMap(slc)
	}
}
