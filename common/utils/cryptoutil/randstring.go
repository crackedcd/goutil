package cryptoutil

import (
    "math/rand"
    "strings"
    "text/scanner"
    "time"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
    letterIdxBits = 6
    letterIdxMask = 1<<letterIdxBits - 1
    letterIdxMax  = 63 / letterIdxBits
)

func RandStr(n int, letter string) string {
    b := make([]byte, n)
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letter) {
            b[i] = letter[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }
    return string(b)
}

func RangeInt(min int, max int) int {
    //rand.Seed(time.Now().UnixNano())
    randNum := rand.Intn(max-min+1) + min
    return randNum
}

func Shuffle(s string) string {
    tokens := []rune(s)
    var src []string
    for _, char := range tokens {
        src = append(src, scanner.TokenString(char))
    }
    letters := make([]string, len(src))
    //rand.Seed(time.Now().UTC().UnixNano())
    perm := rand.Perm(len(src))
    for i, v := range perm {
        letters[v] = src[i]
    }
    return strings.Join(letters, "")
}

