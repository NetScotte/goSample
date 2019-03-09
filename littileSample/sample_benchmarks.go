package main 

func BenchmarkIndex(b *testing.B) {
	const s = "some_text=somevalue"
	for i := 0, i < b.N; i++ {
		strings.Index(s, "v")
	}
}

//run: go test -test.bench=Index