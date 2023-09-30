package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTableTest(b *testing.B) {
	benctTest := []struct {
		name, result string
	}{
		{name: "farid", result: "farid"},
		{name: "anang", result: "anang"},
	}
	for _, value := range benctTest {
		b.Run(value.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(value.result)
			}
		})
	}

}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("farid")
	}
}
func BenchmarkHelloWorld2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("anang")
	}
}

// tableTest
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		Name, Request, Expected string
	}{
		{
			Name:     "farid",
			Request:  "farid",
			Expected: "hello farid",
		},
		{
			Name:     "anang",
			Request:  "anang",
			Expected: "hello anang",
		},
		{
			Name:     "samudra",
			Request:  "samudra",
			Expected: "hello samudra",
		},
	}
	for _, test := range tests {
		result := HelloWorld(test.Request)
		t.Run(test.Name, func(t *testing.T) {
			require.Equal(t, test.Expected, result)
		})
	}
}

// subtTest
// kita gunakan subtest untuk membuat func Test di dalam funct Test dengan menggunakan t.Run()
// dan cara menjalankannya dengan mengetik go mod -v -run=TestSubtest atau bisa juga dengan menambah /farid atau anang
func TestSubtest(t *testing.T) {
	t.Run("farid", func(t *testing.T) {
		result := HelloWorld("farid")
		require.Equal(t, "helo farid", result, "Ini bukan hello farid")
	})
	t.Run("anang", func(t *testing.T) {
		result := HelloWorld("anang")
		require.Equal(t, "hello anang", result, "ini buhan hello anang")
	})
}

// Main
func TestMain(m *testing.M) {
	// Testing Unit Before
	fmt.Println("INI UNIT SEBELUM")

	m.Run()

	fmt.Println("INI UNIT SESUDAH")
	// Testing Unit After
}

// Skip(args)
// jika dia menemukan kondisi tertentu dan kita menjalankan t.Skip(arg) maka dia akan langsung
// berhenti atau batal
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("ini tidak bis di jalankan di windows")
	}
	result := HelloWorld("farid")
	require.Equal(t, "hello farid", result)
}

// // require
// kita gunakan require untuk menghemat kode program di mana sebelumnya kita
// menggunakan if else sekarang kita gunakan require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("farid")
	require.Equal(t, "hello farid", result)
}

// // assert
// kita gunakan assert untuk menghemat kode program di mana sebelumnya kita
// menggunakan if else sekarang kita gunakan assert
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("farid")
	assert.Equal(t, "hello farid", result)
	fmt.Println("Hello farid done")
}

// kita gunakan if else dan Fail() atau Error(args)
// jika menemukan error maka kodenya akan tetap berlanjut sampa selesai dan
// tetap menjalankan errornya
func TestHelloWorldFarid(t *testing.T) {
	result := HelloWorld("anang")
	if result != "hello anang" {
		// t.Fail()
		t.Error("Ini bukan hello farid")
	}
	fmt.Println("Hello anang done")
}

// // kita gunakan if else dan FailNow() atau Fatal(args)
// // jika menemukan error maka kodenya akan langsung berhenti
func TestHelloWorldAnang(t *testing.T) {
	result := HelloWorld("anang")
	if result != "hello anang" {
		// t.FailNow()
		t.Fatal("Ini bukan hello anang")
	}
	fmt.Println("Hello anang done")
}
