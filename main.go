package main

import (
	"fmt"
	"html/template"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	err := t.Execute(w, p)
	if err != nil {
		fmt.Println(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test のうちtest部分のみを切り出して、それにマッチするファイルを読み込む
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	// エラーがあった場合でもeditパスに流す？？
	if err != nil {
		fmt.Println(err)
		//p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// サーバー起動、問題があればログ出力するのでこのように書ける。この状態で動かすとlocalhostでwebサーバーが立ち上がる。
	log.Fatal(http.ListenAndServe(":8080", nil))
}
