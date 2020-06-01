package main

import (
        "net/http"
        "log"
        "io/ioutil"
        "time"
        "flag"
        "net/url"
        "golang.org/x/net/proxy"
)

var INTERVAL int
var TARGET string
var SOCKS5_PROXY_URI string


func main() {
        flag.IntVar(&INTERVAL, "INTERVAL", 1, "text description")
        flag.StringVar(&TARGET, "TARGET", "https://check.torproject.org", "text description")
        flag.StringVar(&SOCKS5_PROXY_URI, "SOCKS5_PROXY_URI", "127.0.0.1:9050", "text description")
        flag.Parse()
        doEvery((time.Duration(INTERVAL) * time.Second), MakeRequest)
}

func doEvery(d time.Duration, callback func()) {
        for range time.Tick(d) {
                callback()
        }
}

func MakeRequest() {

        proxyURL, err := url.Parse("socks5://"+SOCKS5_PROXY_URI)
        dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
        if err != nil {
                        panic(err)
        }
        transport := &http.Transport{Dial: dialer.Dial}
        client := http.Client{Transport: transport}


        resp, err := client.Get(TARGET)
        if err != nil {
                log.Fatalln(err)
        }

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Fatalln(err)
        }
        log.Println(string(body))
}
