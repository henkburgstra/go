// Flow Based Programming
package fbp

import (
//	"code.google.com/p/gcfg"
)

func Connect(from *Port, to *Port) {
	pipe := NewPipe()
	pipe.Connect(from, to)
}
