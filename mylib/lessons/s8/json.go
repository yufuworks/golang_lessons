package s8

import (
	"encoding/json"
	"fmt"
)

type EmptyStruct struct {
}

type Person struct {
	// 通常の表記だと json.Marshal でそのまま構造体の名前を使ってしまうのでバッククォートで出力する名前をつける
	// パースするときにも指定した名前で見るようになるが大文字小文字は区別しない。それ以外に変えるとこのコードの場合unmarshalが正しく動かなくなる。
	// 全く別の名前を指定もできる。
	//Name      string   `json:"nameXXXX"`
	// 型を変えることもできる
	//Age       int      `json:"age,string"`
	// jsonで取り扱わないようにもできる。
	//Name      string   `json:"-"`
	// 値が0や空の時は出力しない場合
	//Age  int    `json:"age,omitempty"`
	// 空の構造体を除外したい場合はポインタを渡さないとだめ
	//T         EmptyStruct `json:"T,omitempty"`
	Name      string       `json:"name"`
	Age       int          `json:"age"`
	Nicknames []string     `json:"nicknames"`
	T         *EmptyStruct `json:"T,omitempty"`
}

// この名前で関数を作成するとmarshal時の動作をカスタムできる
// この関数をコメントアウトすると元の動作に戻る
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

// この名前で関数を作成するとunmarshal時の動作をカスタムできる。
func (p *Person) UnmarshalJSON(b []byte) error {
	//一時的に使う構造体を定義、mashalでも使える
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func Main04() {
	b := []byte(`{"name":"mike","age":20,"nicknames":["a","b","c"]}`)
	var p Person
	// jsonのデータと一致する要素を構造体に自動で割り当ててくれる。大文字小文字を区別しない。
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// 構造体からjsonを作成してくれる
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
