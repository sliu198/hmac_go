package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        log.Fatalln("must contain exactly 2 arguments")
    }
    key,_ := base64.RawURLEncoding.DecodeString(os.Args[1])
    hash := hmac.New(sha256.New, key)
    hash.Write([]byte(os.Args[2]))
    sig := hash.Sum(nil)
    sigB64 := base64.RawURLEncoding.EncodeToString(sig)
    fmt.Println(sigB64)
}
