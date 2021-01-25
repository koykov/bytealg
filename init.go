package bytealg

import (
	"github.com/koykov/x2bytes"
)

func init() {
	// Register chain buffer to bytes function.
	x2bytes.RegisterToBytesFn(ChainBufToBytes)
}
