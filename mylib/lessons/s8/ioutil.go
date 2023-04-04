package s8

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Main03() {
	//ioutilは現在非推奨
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	//↑のコードの場合contentもerrも同じ読み込んだファイルの文字列が入るだけなので、読み込んだファイルと同じものが新しいファイルとして出力される。
	fmt.Println(string(content))
	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
		log.Fatalln(err)
	}
}
