package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

/*
rand.Seed(time.Now().UnixNano()) sử dụng thời gian hiện tại tính bằng nanosecond

	như một giá trị seed để đảm bảo rằng mỗi lần bạn chạy chương trình, chuỗi số ngẫu nhiên mà bạn nhận được sẽ không trùng lặp.
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0 -> max - min
}

// RandomString generates a random a string of length n
func RandomString(n int) string {
	// Bui;der type help creat string effectively
	var sb strings.Builder
	k := len(alphabet)

	// In each loop, creating a random int 0 -> k and then write it to Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMomey generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%@email.com", RandomString(6))
}
