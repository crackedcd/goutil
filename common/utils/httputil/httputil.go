package httputil

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

func Get(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        log.Println(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
    }
    defer resp.Body.Close()
    return string(body)
}

func Post(url string, form string) string {
    resp, err := http.Post(url,
        "application/x-www-form-urlencoded",
        strings.NewReader(form))
    if err != nil {
        log.Println(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
    }
    defer resp.Body.Close()
    return string(body)
}

func Do(url string, method string, form string) string {
    client := &http.Client{}
    req, err := http.NewRequest(method, url, strings.NewReader(form))
    if err != nil {
        log.Println(err)
    }
    // header
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("User-Agent",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) " +
            "Chrome/72.0.3626.121 Safari/537.36")
    resp, err := client.Do(req)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
    }
    defer resp.Body.Close()
    return string(body)
}

