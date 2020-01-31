package main

// Extended Ping

import (
    "bytes"
    "fmt"
    "log"
    "net/url"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

func ePing(pingCount, pingHost string) error {
    cmd := exec.Command("ping", pingCount, pingHost)
    stdout := bytes.Buffer{}
    cmd.Stdout = &stdout
    if err := cmd.Start(); err != nil {
        return err
    }
    // realtime echo
    if err := cmd.Wait(); err != nil {
        return err
    } else {
        fmt.Println(stdout.String())
    }
    return nil
}

func main() {
    // eping https://baidu.com:443/urlPath/index.html
    if len(os.Args) < 1 {
        fmt.Println("Please enter host.")
        os.Exit(0)
    }
    if len(os.Args) == 2 {
        var pingUrl string = os.Args[1]
        var pingCount string
        // get host: baidu.com
        if !strings.HasPrefix(os.Args[1], "http") {
            pingUrl = "http://" + os.Args[1]
        }
        pingUrlInfo, err := url.Parse(pingUrl)
        if err != nil {
            log.Fatal("url.Parse err:", err)
        }
        pingHost := strings.Split(pingUrlInfo.Host, ":")[0]
        // windows default ping 4 times
        if runtime.GOOS != "windows" {
            pingCount = "-c4"
        } else {
            pingCount = ""
        }
        // exec command
        err = ePing(pingCount, pingHost)
        if err != nil {
            log.Fatal("func ePing err:", err)
        }
    }
    if len(os.Args) > 2 {
        fmt.Println("Please enter one parameter.")
        os.Exit(0)
    }
}
