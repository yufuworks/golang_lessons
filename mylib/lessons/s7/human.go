// Package s7 is made by golang lesson
package s7

import "fmt"

/*
Person 頭文字が大文字だとパブリックになり他のパッケージから参照できる。
小文字のフィールドは外部から参照されない。
*/
type Person struct {
	// Name hogehoge 説明 go doc ... でターミナルから参照される。
	Name   string
	Age    int
	height int
}

// SetHeight ポインター型*Personで作ればプライベートフィールドに対するSetterが作れる
func (p *Person) SetHeight(height int) {
	p.height = height
}

func Say() {
	fmt.Println("Human!")
}
