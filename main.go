package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
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

// 読み込んだtxtファイル（ページ）をキャッシングしておく
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// キャッシュに用意したtemplatesに対してExecuteTemplateすることでファイルを毎回読み込まなくて済む
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /view/test のうちtest部分のみを切り出して、それにマッチするファイルを読み込む
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	// rを使ってないと注意されたのでトリあずログに出してみた
	path := r.URL.Path
	log.Printf("path: %v\n", path)
	// エラーがあった場合でもeditパスに流す？？
	if err != nil {
		// もし存在しないページにアクセスしてもeditページに進むことでその後保存すれば新しいページとして保存される。
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// saveできた場合元のviewに戻る
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// パスのバリデーション、実装したページ三つのパスを判定する
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, string string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエストパスをバリデーションにかける、/区切りのスライスになる。
		m := validPath.FindStringSubmatch(r.URL.Path)
		log.Printf("m type:%T value:%v m2:%v\n", m, m, m[2])
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// mの2つめの要素が各Handler関数の第三引数、title strin になる想定
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	// サーバー起動、問題があればログ出力するのでこのように書ける。この状態で動かすとlocalhostでwebサーバーが立ち上がる。
	log.Fatal(http.ListenAndServe(":8080", nil))
}
