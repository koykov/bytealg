package bytealg

import "github.com/koykov/any2bytes"

func init() {
	// Register chain buffer to bytes function.
	any2bytes.RegisterAnyToBytesFn(ChainBufToBytes)
}
