package messages

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

//Start with Test then next letter must be in Capital letter.
// Immediate failure -t.FailNow , Fatal and Fatalf. Fatal will stop the execution of tests
// Non immediate - t.Fail , Error and Errorf
func TestGreet(t *testing.T) {
	got := Greet("Ankit")
	expect := "Hello, Ankit!"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q got: %q", expect, got)
	}

}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Ankit", expect: "Hello, Ankit!"},
		{input: "", expect: "Hello, !"},
	}

	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not got expected result for input: '%v'. Expected %q and got %q", s.input, got, s.expect)
		}
	}

}

func TestDapart(t *testing.T) {
	t.Log("Running Depart Test")
	got := depart("Ankit")
	expect := "Goodbye, Ankit"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q got: %q", expect, got)
	}
}

func BenchmarkSHA1(b *testing.B) {
	data := []byte("My Benchmark testing..")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("My Benchmark testing..")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("My Benchmark testing..")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}

func BenchmarkSHA512Alloc(b *testing.B) {
	data := []byte("My Benchmark testing..")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h := sha512.New()
		h.Sum(data)
	}
}
