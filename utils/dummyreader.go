package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/bxcodec/faker/v4"
)

const lo, hi int = 10, 100

func GetReader() (<-chan string, int) {
	line := make(chan string)
	randomString := faker.Name()
	n := generateRandomNumber(hi, lo)
	go func() {
		for i := 0; i < n; i++ {
			line <- fmt.Sprintf(`%s: Line no %d`, randomString, i)
		}
		close(line)
	}()
	return line, n
}

func generateRandomNumber(hi, lo int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(hi)))
	y := x.Int64()
	rn := lo + int(y)
	return rn
}
