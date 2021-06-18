package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

/* 你会看到下面的代码和写测试差不多。
testing.B 可使你访问隐性命名（cryptically named）b.N。
基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
代码运行的次数不会对你产生影响，测试框架会选择一个它所认为的最佳值，以便让你获得更合理的结果。
用 go test -bench=. 来运行基准测试。 (如果在 Windows Powershell 环境下使用 go test -bench=".")
*/
//注意：基准测试默认是顺序运行的。
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

/*练习
修改测试代码，以便调用者可以指定字符重复的次数，然后修复代码
写一个 ExampleRepeat 来完善你的函数文档
看一下 strings(https://golang.org/pkg/strings) 包。找到你认为可能有用的函数，并对它们编写一些测试。投入时间学习标准库会慢慢得到回报。
*/
