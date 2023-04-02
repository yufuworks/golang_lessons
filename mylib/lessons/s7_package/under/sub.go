package under

import "fmt"

/*
Person 頭文字が大文字だとパブリックになり他のパッケージから参照できる。
小文字のフィールドは外部から参照されない。
*/
type Person struct {
	Name   string
	Age    int
	height int
}

// SetHeight ポインター型*Personで作ればプライベートフィールドに対するSetterが作れる
func (p *Person) SetHeight(height int) {
	p.height = height
}

// Public 変数も大文字でパブリックになりパッケージ外から呼び出し可能。
var Public string = "Public"
var private string = "private"

// Hello javaのクラス名のようなファイル名による制限は特になさそう？
func Hello() {
	fmt.Println("Hello!")
}
