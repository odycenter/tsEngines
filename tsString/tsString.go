package tsString

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"github.com/odycenter/tsEngines/tsFuzzy"
	"github.com/odycenter/tsEngines/tsMap"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/transform"

	"github.com/axgle/mahonia"
	"github.com/shopspring/decimal"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IfString(b bool, t, f string) string {
	if b {
		return t
	}
	return f
}

func And(src int64, data int64) int64 {
	return src & data
}
func Or(src int64, data int64) int64 {
	return src | data
}

/*
*
src:="hello! 你好！"
16 = len(src)
10 = utf8.RuneCountInString(src)
*/
func Len(src string) int {
	return utf8.RuneCountInString(src)
}

/*
*
Hidden("张三丰", 1, 1)
输出“张*丰”
第三个参数，默认为：'*'
*/
func Hidden(src string, showBegin int, showEnd int, hidden ...rune) string {
	placeholder := '*'
	if len(hidden) > 0 {
		placeholder = hidden[0]
	}

	srcRune := []rune(src)
	tempRune := srcRune[:showBegin]
	count := len(srcRune)
	if showBegin+showEnd > count {
		return src
	}
	startLen := count - showEnd

	for {
		if len(tempRune) < startLen {
			tempRune = append(tempRune, placeholder)
		} else {
			break
		}
	}

	return string(tempRune) + string(srcRune[startLen:])
}
func Hidden2(src string, showBegin int, showEnd int, hiddenCount int, hidden ...rune) string {
	placeholder := '*'
	if len(hidden) > 0 {
		placeholder = hidden[0]
	}

	srcRune := []rune(src)
	tempRune := srcRune[:showBegin]
	count := len(srcRune)
	if showBegin+showEnd > count {
		return src
	}

	startLen := count - showEnd
	startLenNew := count - showEnd
	if startLen > showBegin+hiddenCount {
		startLen = showBegin + hiddenCount
	}
	for {
		if len(tempRune) < startLen {
			tempRune = append(tempRune, placeholder)
		} else {
			break
		}
	}

	return string(tempRune) + string(srcRune[startLenNew:])
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// 轉換 文字轉浮點數
func StringToFloat64(info string) float64 {
	tmp, _ := strconv.ParseFloat(info, 64)
	return tmp
}

// 轉換 文字轉數字 32位元
func StringToInt(info string) int {
	tmp, _ := strconv.Atoi(info)
	return tmp
}

// 轉換 文字轉數字 64位元
func StringToInt64(info string) int64 {
	tmp, _ := strconv.ParseInt(info, 10, 64)
	return tmp
}

// 轉換 文字轉小數點
func StringToDecimal(info string) decimal.Decimal {
	tmp, _ := decimal.NewFromString(info)
	return tmp
}

func StructToJsonString(structt interface{}) (jsonString string, err error) {
	data, err := json.Marshal(structt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func InterfaceToString(x interface{}) string {
	tmp := fmt.Sprintf("%v", x)
	return tmp
}

func ToBool(str string) bool {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	}
	return false
}

func FromInt(v int) string {
	return fmt.Sprintf("%d", v)
}

func FromInt64(v int64) string {
	return fmt.Sprintf("%d", v)
}

func FromUInt64(v uint64) string {
	return fmt.Sprintf("%d", v)
}

func FromByte(v byte) string {
	return fmt.Sprintf("%d", v)
}

func Split(s, sep string, needSpace ...bool) []string {
	var ns bool
	if len(needSpace) != 0 {
		ns = needSpace[0]
	}
	var arr []string
	for _, v := range strings.Split(s, sep) {
		if v == "" && !ns {
			continue
		}
		arr = append(arr, v)
	}
	return arr
}

func JoinIntArr(elems []int, sep string) string {
	buff := strings.Builder{}
	for k, v := range elems {
		if k != 0 {
			buff.WriteString(sep)
		}
		buff.WriteString(fmt.Sprintf("%d", v))
	}
	return buff.String()
}

// 删除 byte为32的（空格），和左右的（空格）
func TrimSpace(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Replace(str, " ", "", -1)
	return str
}

// 删除左右空格
func TrimLrSpace(str string) string {
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func GetInvitationCode(l int) string {
	str := "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = str[rand.Intn(len(str))]
	}
	return string(bytes)
}

func GetRandomInt(l int) string {
	str := "0123456789"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = str[rand.Intn(len(str))]
	}
	return string(bytes)
}

func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = str[rand.Intn(len(str))]
	}
	return string(bytes)
}

// 将数字类型的数组转换为字符串
func ImplodeIntToString(arr []int64, seq string) (s string) {
	for _, v := range arr {
		s += FromInt64(v) + seq
	}

	return strings.TrimRight(s, seq)
}

// 将字符串类型的数组转换为字符串
func ImplodeStringToString(arr []string, seq string) (s string) {
	for _, v := range arr {
		s += v + seq
	}

	return strings.TrimRight(s, seq)
}

func CoverStringToArray(str string, sep string, needSpace bool) (arr []string) {
	a := strings.Split(str, ",")
	for _, v := range a {
		if v == "" && !needSpace {
			continue
		}
		arr = append(arr, v)
	}
	return
}
func CoverStringToInt64Array(str string, sep string, needSpace bool) (arr []int64) {
	a := strings.Split(str, ",")
	for _, v := range a {
		if v == "" && !needSpace {
			continue
		}
		arr = append(arr, StringToInt64(v))
	}
	return
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func CoverCamelToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CoverSnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 首字母小写
func LowerBegin(s string) string {
	if len(s) > 0 {
		b := []byte(s)
		c := b[0]
		if c >= 'A' && c <= 'Z' {
			b[0] = c + 'a' - 'A'
		}
		return string(b)
	}
	return ""
}

// 首字母大写
func UpperBegin(s string) string {
	if len(s) > 0 {
		b := []byte(s)
		c := b[0]
		if c >= 'a' && c <= 'z' {
			b[0] = c + 'A' - 'a'
		}
		return string(b)
	}
	return ""
}

/*
*
- 转换当前乱码内容，服务http-get调用返回中文乱码问题
*/
func ConvertByte2StringCorrect(byte []byte, charset string) string {
	var str string

	switch charset {
	case "GB18030":
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case "UTF8":
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func CoverInt64Arr2String(arr []int64) (d []string) {
	if len(arr) == 0 {
		return
	}
	for _, a := range arr {
		d = append(d, FromInt64(a))
	}
	return
}

func CoverIntArr2String(arr []int) (d []string) {
	if len(arr) == 0 {
		return
	}
	for _, a := range arr {
		d = append(d, FromInt(a))
	}
	return
}

func CoverBytes2String(arr []byte) (d string) {
	if len(arr) == 0 {
		return
	}
	for _, a := range arr {
		d += FromByte(a)
	}
	return
}

func ConvertString2Int32Arr(str, sep string) (d []int32) {
	if len(str) == 0 {
		return
	}
	arr := strings.Split(str, sep)
	for _, a := range arr {
		if a == "" {
			continue
		}
		d = append(d, int32(StringToInt(a)))
	}
	return
}

func ConvertString2Int64Arr(str, sep string) (d []int64) {
	if len(str) == 0 {
		return
	}
	arr := strings.Split(str, sep)
	for _, a := range arr {
		if a == "" {
			continue
		}
		d = append(d, StringToInt64(a))
	}
	return
}

func ConvertString2IntArr(str, sep string) (d []int) {
	if len(str) == 0 {
		return
	}
	arr := strings.Split(str, sep)
	for _, a := range arr {
		if a == "" {
			continue
		}
		d = append(d, StringToInt(a))
	}
	return
}

// 去除数组字符串的特殊字符
func TrimArrSep(arr []string, sep ...string) []string {
	var sepStr = "'" //默认去除单引号
	if len(sep) > 0 {
		sepStr = sep[0]
	}
	for num, row := range arr {
		if !strings.Contains(row, sepStr) { //判断是否存在单引号或者特殊字符，如果不存在，则直接跳出wei
			break
		}
		arr[num] = strings.ReplaceAll(row, sepStr, "")
	}
	return arr
}

//int to string

func Int64ToString(num int64) string {
	return fmt.Sprintf("%d", num)
}

func Int32ToString(num int32) string {
	return fmt.Sprintf("%d", num)
}
func IntToString(num int) string {
	return fmt.Sprintf("%d", num)
}
func Cleanse(s string, forceASCII bool) string {
	//只保留英文字母或數字, 其餘符號轉為 ''
	if forceASCII {
		s = ASCIIOnly(s)
	}
	s = strings.TrimSpace(s)
	rs := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			r = ' '
		}
		rs = append(rs, r)
	}
	return strings.ToLower(string(rs))
}
func ASCIIOnly(s string) string {
	b := make([]byte, 0, len(s))
	for _, r := range s {
		if r <= unicode.MaxASCII {
			b = append(b, byte(r))
		}
	}
	return string(b)
}

type MatchPair struct {
	Match string
	Score int
}

type MatchPairs []*MatchPair

func GetSimilarStrings(newChat string, oldChats []string, ratio float64) (MatchPairs, []string) {
	results := MatchPairs{}

	if len(oldChats) == 0 {
		return results, nil
	}

	var logStrArr []string

	//1. 將新的發話 與 歷史發話 分拆為 中文 與 英文
	//且將文字以外的字符替換成 " "
	oldChineseChats := make([]string, 0, len(oldChats))
	oldEnglishChats := make([]string, 0, len(oldChats))
	for index, v := range oldChats {
		logStrArr = append(logStrArr, fmt.Sprintf("oldChats[%d] is utf8 : %t, v : %s", index, Isutf8(v), v))
		oldChineseChat := ""
		oldEnglishChat := ""
		oldChineseChat, oldEnglishChat = seperateChineseAndEnglish(v)

		oldChineseChats = append(oldChineseChats, oldChineseChat)
		oldEnglishChats = append(oldEnglishChats, oldEnglishChat)
		logStrArr = append(logStrArr, fmt.Sprintf("舊中文 : '%s'", oldChineseChats))
		logStrArr = append(logStrArr, fmt.Sprintf("舊英文 : '%s'\n", oldEnglishChats))
	}

	newChineseChat := ""
	newEngChat := ""
	newChineseChat, newEngChat = seperateChineseAndEnglish(newChat)

	logStrArr = append(logStrArr, fmt.Sprintf("\n新中文 : = '%s'", newChineseChat))
	logStrArr = append(logStrArr, fmt.Sprintf(", 新英文 : = '%s'", newEngChat))

	decimal.DivisionPrecision = 2
	for index, v := range oldChats {
		chineseRatio := decimal.NewFromFloat(0.0)
		engRatio := decimal.NewFromFloat(0.0)
		oldChineseStr := oldChineseChats[index]
		oldEngStr := oldEnglishChats[index]
		if len(newChineseChat) != 0 && len(oldChineseStr) != 0 {
			//2. 比對簡繁體中文
			cRatio, log := ChineseMatchRatio(newChineseChat, oldChineseStr)
			chineseRatio = cRatio.Mul(decimal.NewFromInt(100))

			logStrArr = append(logStrArr, fmt.Sprintf("\noldChats[%d],  %s, cRatio : %s", index, log, cRatio.String()))
		} else {
			logStrArr = append(logStrArr, fmt.Sprintf("\noldChats[%d],無需比對 新 : %d,  舊 : %d", index, len(newChineseChat), len(oldChineseStr)))
		}

		if len(newEngChat) != 0 && len(oldEngStr) != 0 {
			//3. 比對英文
			engRatio = EnglishMatchRation(newEngChat, oldEngStr)
		}

		var engCharCount []string
		if len(oldEngStr) > 0 {
			engCharCount = strings.Split(oldEngStr, " ")
		}

		totalLength := Len(oldChineseStr) + len(engCharCount)
		dTotalLength := decimal.NewFromInt(int64(totalLength))
		dOldChineseStrLength := decimal.NewFromInt(int64(Len(oldChineseStr)))
		deEngCharCountLength := decimal.NewFromInt(int64(len(engCharCount)))

		resRatio := chineseRatio.Mul(dOldChineseStrLength).Div(dTotalLength).
			Add(engRatio.Mul(deEngCharCountLength).Div(dTotalLength))

		//resRatio := chineseRatio*float64(Len(oldChineseStr))/float64(totalLength) +
		//	engRatio*float64(len(engCharCount))/float64(totalLength)

		if resRatio.Round(2).GreaterThanOrEqual(decimal.NewFromFloat(ratio)) {
			iRatio, _ := strconv.Atoi(resRatio.String())
			pair := &MatchPair{Match: v, Score: iRatio}
			results = append(results, pair)
		}
	}

	return results, logStrArr
}

func removeLastSpaceChar(r []rune) []rune {
	lastR := string(r[len(r)-1:])
	if lastR == " " {
		r = r[0 : len(r)-1]
	}

	return r
}

func seperateChineseAndEnglish(oldChat string) (chineseStr string, englishStr string) {
	isChinese := regexp.MustCompile("\\p{Han}")

	strArr := []rune(oldChat)

	chineseArr := make([]rune, 0, Len(string(strArr)))
	englishArr := make([]rune, 0, Len(string(strArr)))

	for _, v := range strArr {
		tmpStr := string(v)
		if isChinese.MatchString(tmpStr) {
			chineseArr = append(chineseArr, v)
		} else {
			if !unicode.IsLetter(v) && !unicode.IsNumber(v) {
				if len(englishArr) > 0 {
					lastR := string(englishArr[len(englishArr)-1:])
					if lastR != " " {
						englishArr = append(englishArr, ' ')
					}
				} else {
					englishArr = append(englishArr, ' ')
				}
			} else {
				englishArr = append(englishArr, v)
			}
		}
	}

	if len(chineseArr) > 0 {
		chineseArr = removeLastSpaceChar(chineseArr)
		chineseStr = string(chineseArr)
	} else {
		chineseStr = ""
	}

	if len(englishArr) > 0 {
		englishArr = removeLastSpaceChar(englishArr)
		englishStr = string(englishArr)
	} else {
		englishStr = ""
	}

	return chineseStr, englishStr
}

func EnglishMatchRation(query, matchWords string) decimal.Decimal {
	scorer := func(s1, s2 string) int {
		return tsFuzzy.WRatio(s1, s2)
	}

	score := scorer(query, matchWords)
	return decimal.NewFromInt(int64(score))
}

func ChineseMatchRatio(str1, matchWords string) (decimal.Decimal, string) {
	m := tsMap.NewMapStrMgr()

	var logsStr string

	charsA := []rune(str1)
	charsB := []rune(matchWords)

	low, high := charsA, charsB
	if len(low) > len(high) {
		low = charsB
		high = charsA
	}

	logsStr += fmt.Sprintf("low : '%c', hight : '%c'", low, high)

	//1. 聯集字串1到 Set Map
	for i := 0; i < len(high); i++ {
		ok := m.Have(string(high[i]))
		if !ok {
			m.Add(string(high[i]), m.ObjMgr)
		}
	}

	//2. 聯集字串2到 Set Map
	for i := 0; i < len(low); i++ {
		ok := m.Have(string(low[i]))
		if !ok {
			m.Add(string(low[i]), m.ObjMgr)
		}
	}

	logsStr += fmt.Sprintf(", m.Count : %d", m.Count())

	arrA := make([]int, m.Count())
	arrB := make([]int, m.Count())

	//3. 計算字串1與聯集Set相同字數
	count := 0
	for k := range m.ObjMgr {
		eqCount := 0
		for j := 0; j < len(charsA); j++ {
			if k == string(charsA[j]) {
				eqCount = eqCount + 1
			}
		}
		arrA[count] = eqCount
		count++
	}
	logsStr += fmt.Sprintf(", 計算字串1與聯集Set相同字數 : %d", count)

	//3. 計算字串2與聯集Set相同字數
	count = 0
	for k := range m.ObjMgr {
		eqCount := 0
		for j := 0; j < len(charsB); j++ {
			if k == string(charsB[j]) {
				eqCount = eqCount + 1
			}
		}
		arrB[count] = eqCount
		count++
	}
	logsStr += fmt.Sprintf(", 計算字串2與聯集Set相同字數 : %d", count)

	//計算比例
	num := decimal.NewFromFloat(0.0)
	numA := decimal.NewFromFloat(0.0)
	numB := decimal.NewFromFloat(0.0)

	for i := 0; i < m.Count(); i++ {
		num = num.Add(decimal.NewFromFloat(float64(arrA[i] * arrB[i])))
		numA = numA.Add(decimal.NewFromFloat(math.Pow(float64(arrA[i]), 2)))
		numB = numB.Add(decimal.NewFromFloat(math.Pow(float64(arrB[i]), 2)))
	}

	rat := decimal.NewFromFloat(0.0)
	aF, _ := numA.Float64()
	bF, _ := numB.Float64()
	c := decimal.NewFromFloat(math.Sqrt(aF) * math.Sqrt(bF))
	if c.GreaterThan(decimal.NewFromFloat(0.0)) {
		dd := decimal.NewFromFloat(math.Sqrt(aF) * math.Sqrt(bF))
		rat = num.Div(dd)
	}

	logsStr += fmt.Sprintf(", num : %s, numA : %s, numB : %s ", num.String(), numA.String(), numB.String())

	return rat, logsStr
}

// 是否utf8
func Isutf8(s string) bool {
	return utf8.ValidString(s)
}

// 是否簡體
func Isgbk(s string) bool {
	if Isutf8(s) {
		return false
	}
	data := []byte(s)
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 字符串去重
func RemoveDuplicationString(str string) string {
	if str == "" {
		return ""
	}
	arr := Split(str, ",")
	newArr := RemoveDuplicationArray(arr)
	return ImplodeStringToString(newArr, ",")
}

// 数组去重
func RemoveDuplicationArray(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		if v == "" {
			continue
		}
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}

// RepeatAndJoin
// input: ?, 6
// output: ?,?,?,?,?,?
func RepeatAndJoin(str string, len int) string {
	var arr []string
	for i := 0; i < len; i++ {
		arr = append(arr, str)
	}
	return strings.Join(arr, ",")
}

func RemoveRepeatedInt(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
