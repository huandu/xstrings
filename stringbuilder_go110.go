//+build !go1.10

package xstrings

import "bytes"

type bufferString struct {
	bytes.Buffer
}
