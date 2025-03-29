package utils

import "os"

func FileExits(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
