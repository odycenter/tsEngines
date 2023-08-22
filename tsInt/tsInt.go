package tsint

import "strconv"

func IntToString(info int) string {
	tmp := strconv.Itoa(info)
	return tmp
}

func Int64ToString(info int64) string {
	tmp := strconv.FormatInt(info, 10)
	return tmp
}
