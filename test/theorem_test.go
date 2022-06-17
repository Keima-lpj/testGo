package test

import "testing"

func TestTheorem(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
	}

	for k, v := range tests {
		if ans := getTheorem(v.a, v.b); ans != v.c {
			t.Errorf("第%d个测试数据错误!  期望%d，得到%d", k, v.c, ans)
		}
	}
}

func BenchmarkTheorem(t *testing.B) {
	for i := 0; i < t.N; i++ {
		if ans := getTheorem(5, 12); ans != 13 {
			t.Errorf("111")
		}
	}
}
