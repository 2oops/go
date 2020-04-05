package demo

import (
	"testing"
	"fmt"
)

// // 单元测试 go test -v
// func TestGetArea(t *testing.T) {
// 	area := GetArea(40, 100)
// 	if area != 4000 {
// 		t.Error("Failure")
// 	}
// }

// func TestAdd(t *testing.T) {
// 	area := Add(40, 100)
// 	// t.FailNow() // 标记单元测试结果
	
// 	if area != 140 { // demo_test.go:16: Failure
// 		t.Error("Failure")
// 	}
// }

// 性能测试  go test -bench="."
// func BenchmarkGetArea(t *testing.B) {
// 	for i := 0; i < t.N; i ++ {
// 		GetArea(400, 100)
// 	}
// }

// func BenchmarkAdd(t *testing.B) {
// 	for i := 0; i < t.N; i ++ {
// 		Add(400, 100)
// 	}
// 	t.Log("end")
// }


// 内存测试 go test -v -bench=Alloc -benchmen demo_test.go
func Benchmark_Alloc(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
	b.StopTimer()
}
// 覆盖率测试 go test -cover
// 知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码）
// 代码即为性能测试和单元测试函数
