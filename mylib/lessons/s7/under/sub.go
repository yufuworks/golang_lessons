package under

import "fmt"

// Public 変数も大文字でパブリックになりパッケージ外から呼び出し可能。
var Public string = "Public"
var private string = "private"

// Hello javaのクラス名のようなファイル名による制限は特になさそう？
func Hello() {
	fmt.Println("Hello!")
}
