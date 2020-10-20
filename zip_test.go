package zip

import (
	"encoding/hex"
	"testing"
)

const (
	source   = "The only thing new in the world is the history you don't know."
	expected = "1f8b08000000000002ff1cc7c10980300c05d055fecd9bd33842830996ffa18d846c2f787c971bc4d9480fdea01582483794d61c88fdc363a756a3f5628847e2a1eafc020000ffffa4b3debf3e000000"
)

func TestOne(t *testing.T) {
	compressed := Compress([]byte(source))
	encoded := hex.EncodeToString(compressed)

	if encoded != expected {
		t.Fatalf("failed to compress")
	}

	decompressed, err := Decompress(compressed)
	if err != nil {
		t.Fatalf("error during decompression (%v)", err)
	}

	if string(decompressed) != source {
		t.Fatalf("roundtrip failed")
	}
}
