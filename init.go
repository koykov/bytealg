package bytealg

// todo remove due to move to github.com/koykov/bytebuf

import (
	"github.com/koykov/x2bytes"
)

func init() {
	// Register chain buffer to bytes function.
	x2bytes.RegisterToBytesFn(ChainBufToBytes)
}
