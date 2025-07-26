package texts

import (
	"encoding"
	"encoding/hex"

	"github.com/pkg/errors"
)

var (
	_ encoding.TextMarshaler   = Hex("")
	_ encoding.TextUnmarshaler = (*Hex)(nil)
)

type Hex []byte

func (h Hex) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(h)), nil
}

func (h *Hex) UnmarshalText(data []byte) error {
	decoded, err := hex.DecodeString(string(data))
	if err != nil {
		return errors.WithStack(err)
	}

	*h = Hex(decoded)

	return nil
}
