package fileutil

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func PrintFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Split(bufio.ScanRunes)
    /*
       ScanLines (默认)
       ScanWords
       ScanRunes (遍历UTF-8字符非常有用)
       ScanBytes
    */

    //是否有下一行
    for scanner.Scan() {
        fmt.Println("read string:", scanner.Text())
    }
}

