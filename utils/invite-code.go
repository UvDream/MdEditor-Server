package utils

//
//import (
//	"errors"
//	"math/rand"
//	"strings"
//	"time"
//)
//
//type code struct {
//	base    string // 进制的包含字符, string类型
//	decimal uint64 // 进制长度
//	pad     string // 补位字符,若生成的code小于最小长度,则补位+随机字符, 补位字符不能在进制字符中
//	len     int    // code最小长度
//}
//
//// id转code
//func (c *code) idToCode(id uint64) string {
//	mod := uint64(0)
//	res := ""
//	for id != 0 {
//		mod = id % c.decimal
//		id = id / c.decimal
//		res += string(c.base[mod])
//	}
//	resLen := len(res)
//	if resLen < c.len {
//		res += c.pad
//		for i := 0; i < c.len-resLen-1; i++ {
//			rand.Seed(time.Now().UnixNano())
//			res += string(c.base[rand.Intn(int(c.decimal))])
//		}
//	}
//	return res
//}
//
//// code转id
//func (c *code) codeToId(code string) uint64 {
//	res := uint64(0)
//	lenCode := len(code)
//
//	//var baseArr [] byte = []byte(c.base)
//	baseArr := []byte(c.base)     // 字符串进制转换为byte数组
//	baseRev := make(map[byte]int) // 进制数据键值转换为map
//	for k, v := range baseArr {
//		baseRev[v] = k
//	}
//
//	// 查找补位字符的位置
//	isPad := strings.Index(code, c.pad)
//	if isPad != -1 {
//		lenCode = isPad
//	}
//
//	r := 0
//	for i := 0; i < lenCode; i++ {
//		// 补充字符直接跳过
//		if string(code[i]) == c.pad {
//			continue
//		}
//		index := baseRev[code[i]]
//		b := uint64(1)
//		for j := 0; j < r; j++ {
//			b *= c.decimal
//		}
//		// pow 类型为 float64 , 类型转换太麻烦, 所以自己循环实现pow的功能
//		//res += float64(index) * math.Pow(float64(32), float64(2))
//		res += uint64(index) * b
//		r++
//	}
//	return res
//}
//
//// 初始化检查
//func (c *code) initCheck() (bool, error) {
//	lenBase := len(c.base)
//	// 检查进制字符
//	if c.base == "" {
//		return false, errors.New("base string is nil or empty")
//	}
//	// 检查长度是否符合
//	if uint64(lenBase) != c.decimal {
//		return false, errors.New("base length and len not match")
//	}
//	return true, errors.New("")
//}
