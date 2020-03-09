package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(data string, upper bool) string {

	bytes := md5.Sum([]byte(data))

	if upper {
		return fmt.Sprintf("%X", bytes)
	}
	return fmt.Sprintf("%x", bytes)
}
