package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

var (
	Init     int
	SuperKey string
)

const (
	NUmStr  = "128934567"
	CharStr = "ABCklmnopqrstuDENOPQRSTUVWXYFGHIJKLMZabcdefghijvwxyz"
	SpecStr = "+=@#~,.[]()!%^*"
)

//解析参数
func parseArgs() {
	//需要接受指针，就传递地址,&
	flag.IntVar(&Init, "i", 1, "初始化")
	flag.StringVar(&SuperKey, "sk", "superkey",
		`-f super key`)
	flag.Parse()
}

func InitSuperPasswd() string {
	//初始化密码切片
	var passwd []byte = make([]byte, 2048, 2048)
	//源字符串
	var sourceStr string
	//判断字符类型,如果是数字
	sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
	fmt.Println("source:", sourceStr)

	//遍历，生成一个随机index索引,
	for i := 0; i < 2048; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}

func main() {
	parseArgs()
	if Init == 0 {
		t2 := time.Now()
		rand.Seed(t2.UnixNano())
		passwd := InitSuperPasswd()
		fmt.Println(passwd)
		ioutil.WriteFile("super_key", []byte(passwd), 777)
	}

	fmt.Println(GetNewPassword())
}

func GetNewPassword() string {
	t1 := time.Now()
	time1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 1, 19, 15, 12, t1.Location())
	time2 := time.Date(t1.Year(), t1.Month(), t1.Day(), 5, 25, 21, 14, t1.Location())
	time3 := time.Date(t1.Year(), t1.Month(), t1.Day(), 2, 14, 11, 17, t1.Location())
	time4 := time.Date(t1.Year(), t1.Month(), t1.Day(), 12, 22, 56, 33, t1.Location())
	time5 := time.Date(t1.Year(), t1.Month(), t1.Day(), 22, 6, 15, 56, t1.Location())
	time6 := time.Date(t1.Year(), t1.Month(), t1.Day(), 4, 14, 15, 10, t1.Location())
	time7 := time.Date(t1.Year(), t1.Month(), t1.Day(), 8, 1, 3, 6, t1.Location())
	time8 := time.Date(t1.Year(), t1.Month(), t1.Day(), 11, 7, 3, 330, t1.Location())
	time9 := time.Date(t1.Year(), t1.Month(), t1.Day(), 15, 33, 12, 552, t1.Location())
	time10 := time.Date(t1.Year(), t1.Month(), t1.Day(), 15, 33, 1, 143, t1.Location())
	time11 := time.Date(t1.Year(), t1.Month(), t1.Day(), 1, 23, 44, 654, t1.Location())
	time12 := time.Date(t1.Year(), t1.Month(), t1.Day(), 4, 12, 21, 8, t1.Location())
	time13 := time.Date(t1.Year(), t1.Month(), t1.Day(), 6, 33, 42, 324, t1.Location())
	time14 := time.Date(t1.Year(), t1.Month(), t1.Day(), 14, 0, 18, 237, t1.Location())
	time15 := time.Date(t1.Year(), t1.Month(), t1.Day(), 9, 1, 34, 890, t1.Location())
	time16 := time.Date(t1.Year(), t1.Month(), t1.Day(), 11, 15, 51, 998, t1.Location())
	time17 := time.Date(t1.Year(), t1.Month(), t1.Day(), 21, 17, 50, 128, t1.Location())
	time18 := time.Date(t1.Year(), t1.Month(), t1.Day(), 23, 19, 29, 338, t1.Location())
	time19 := time.Date(t1.Year(), t1.Month(), t1.Day(), 17, 33, 24, 438, t1.Location())
	time20 := time.Date(t1.Year(), t1.Month(), t1.Day(), 1, 12, 26, 658, t1.Location())
	time21 := time.Date(t1.Year(), t1.Month(), t1.Day(), 2, 15, 12, 128, t1.Location())
	time22 := time.Date(t1.Year(), t1.Month(), t1.Day(), 5, 41, 42, 458, t1.Location())
	time23 := time.Date(t1.Year(), t1.Month(), t1.Day(), 8, 51, 59, 28, t1.Location())
	time24 := time.Date(t1.Year(), t1.Month(), t1.Day(), 9, 46, 12, 218, t1.Location())
	time25 := time.Date(t1.Year(), t1.Month(), t1.Day(), 18, 39, 41, 411, t1.Location())
	time26 := time.Date(t1.Year(), t1.Month(), t1.Day(), 11, 27, 55, 789, t1.Location())
	time27 := time.Date(t1.Year(), t1.Month(), t1.Day(), 15, 54, 56, 432, t1.Location())
	time28 := time.Date(t1.Year(), t1.Month(), t1.Day(), 17, 31, 39, 568, t1.Location())
	time29 := time.Date(t1.Year(), t1.Month(), t1.Day(), 22, 32, 51, 645, t1.Location())
	time30 := time.Date(t1.Year(), t1.Month(), t1.Day(), 19, 52, 42, 125, t1.Location())
	time31 := time.Date(t1.Year(), t1.Month(), t1.Day(), 21, 41, 41, 345, t1.Location())
	time32 := time.Date(t1.Year(), t1.Month(), t1.Day(), 11, 12, 32, 234, t1.Location())

	skey := []int64{time1.UnixNano(), time3.UnixNano(), time2.UnixNano(), time4.Unix(), time6.Unix(), time7.Unix(), time8.Unix(), time9.Unix(), time5.Unix(), time10.Unix(), time11.Unix(), time12.Unix(), time13.Unix(), time14.Unix(), time15.Unix(), time16.Unix(), time17.Unix(), time18.Unix(), time19.Unix(), time20.Unix(), time21.Unix(), time22.Unix(), time23.Unix(), time24.Unix(), time25.Unix(), time26.Unix(), time27.Unix(), time28.Unix(), time29.Unix(), time30.Unix(), time31.Unix(), time32.Unix()}
	superKey, _ := ioutil.ReadFile("super_key")
	data := fmt.Sprintf("%s", string(superKey))
	nPass := ""
	//添加随机数种子
	lenpass := len(skey)
	cup := t1.Day()
	for i := 0; i < lenpass; i++ {
		rand.Seed(skey[i])
		if i == cup {
			nPass = fmt.Sprintf("%s|%s", nPass, string(data[rand.Intn(2048)]))
		} else {
			nPass = fmt.Sprintf("%s%s", nPass, string(data[rand.Intn(2048)]))
		}
	}
	return nPass
}
