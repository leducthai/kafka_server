package hash

import (
	"fmt"
	"hash/fnv"
)

func HashToInt(value interface{}) int {
	strValue, ok := value.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", value)
	}

	h := fnv.New64()
	h.Write([]byte(strValue))
	return int(h.Sum64())
}
