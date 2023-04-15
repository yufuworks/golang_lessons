package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test のうちtest部分のみを切り出して、それにマッチするファイルを読み込む
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	_, err := fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	// サーバー起動、問題があればログ出力するのでこのように書ける。この状態で動かすとlocalhostでwebサーバーが立ち上がる。
	log.Fatal(http.ListenAndServe(":8080", nil))
}
