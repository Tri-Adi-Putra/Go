package go_say_hello

import "testing"

// Perhatikan aturan penamaan function dan parameternya
func TestSayHello(t *testing.T) {
	result := SayHello("Tri")

	// Logika pengecekan manual (nanti kita akan pakai cara lebih canggih)
	if result != "Hello Tri" {
		// Beri tahu Go kalau tes gagal
		t.Error("Harusnya 'Hello Tri', tapi malah dapat:", result)
	}
}
