package s8

import (
	"fmt"
	"regexp"
)

func Main01() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	//r2から検索対象と結果のスライスを作る
	fs2 := r2.FindStringSubmatch("/view/test")
	fmt.Println(fs2, fs2[0], fs2[1], fs2[2])
	//マッチしない文字列を与えるとエラーになる？？
	//fs2 = r2.FindStringSubmatch("/data/test")
	fmt.Println(fs2, fs2[0], fs2[1], fs2[2])
}
