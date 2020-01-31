package main

// Extended Ping

import (
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
    //err := cmd.Run()
    stdout, err := cmd.StdoutPipe()
    cmd.Stderr = cmd.Stdout
    if err != nil {
        return err
    }
    if err = cmd.Start(); err != nil {
        return err
    }
    // realtime echo
    for {
        tmp := make([]byte, 1024)
        _, err := stdout.Read(tmp)
        fmt.Print(string(tmp))
        if err != nil {
            break
        }
    }
    if err = cmd.Wait(); err != nil {
        return err
    }
    return nil
}

func main() {
    // eping https://baidu.com:443/urlPath/index.html
    var pingUrl string = os.Args[1]
    var pingCount string
    if len(os.Args) > 2 {
        fmt.Println("Please enter a parameter.")
        os.Exit(0)
    }
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
