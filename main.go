package main

import (
    "fmt"
    "strings"
    "crypto/ed25519"
    "encoding/base32"
)

func main() {
    pub, priv, err := ed25519.GenerateKey(nil)
    if err != nil {
        fmt.Println(err)
    } else {
        pub_str := base32.StdEncoding.EncodeToString([]byte(pub))
        pub_str = strings.ToLower(pub_str)
        priv_str := base32.StdEncoding.EncodeToString(priv)
        priv_str = strings.ToLower(priv_str)
        fmt.Printf("%v\n%v\n",pub_str,priv_str)
    }
}
