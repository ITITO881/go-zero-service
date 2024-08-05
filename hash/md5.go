package hash

import (
	"crypto/md5"
	"fmt"
	"io"
)

func StringsMd5Hash(args ...string) string {
	h5 := md5.New()
	for i := 0; i < len(args); i++ {
		_, err := io.WriteString(h5, args[i])
		if err != nil {
			return "Hash error"
		}
	}
	return fmt.Sprintf("%x", h5.Sum(nil))
}
