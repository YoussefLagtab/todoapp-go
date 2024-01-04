package helpers

import (
	"strconv"
)

func GetIdFromString(id string) (uint, error) {
	uid, err := strconv.ParseUint(id, 10, 64)
	return uint(uid), err
}
