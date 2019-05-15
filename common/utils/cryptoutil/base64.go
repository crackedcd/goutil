package cryptoutil

import (
    "encoding/base64"
    "log"
)

func Base64Encode(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
    decodeBytes, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        log.Fatalln(err)
    }
    return string(decodeBytes)
}

func Base64UrlEncode(str string) string {
    return base64.URLEncoding.EncodeToString([]byte(str))
}

func Base64UrlDecode(str string) string {
    uDec, err := base64.URLEncoding.DecodeString(str)
    if err != nil {
        log.Fatalln(err)
    }
    return string(uDec)
}