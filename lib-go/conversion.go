package adventofcode

import (
	"log"
	"strconv"
)

func StringToInt(line string) int {
	val, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal("problem converting string to int", err)
	}
	return val
}
