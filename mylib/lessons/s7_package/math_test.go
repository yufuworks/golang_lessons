package s7_package

import "testing"

//var Debug bool = true

/*
デフォルトでは基本的なテストしかできないため、サードパーティのテストパッケージを使うと良い
Ginkgo : rubyライク
Gomega : javascriptライク
*/
func TestAverage(t *testing.T) {
	// 特定条件でテストをスキップさせる
	//if Debug {
	//	t.Skip("Skip Reason")
	//}
	v := Average([]int{1, 2, 3, 4, 5})
	//v := Average([]int{1, 2, 3, 4, 5, 6, 7})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}

// goland で自動作成したテスト
//func TestAverage(t *testing.T) {
//	type args struct {
//		s []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Average(tt.args.s); got != tt.want {
//				t.Errorf("Average() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
