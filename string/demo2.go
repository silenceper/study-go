package main

import(
    "encoding/hex"
    "crypto/rand"
    "fmt"
    "io"
    "encoding/base64"
)

func main(){
    b:=make([]byte,1)
    n,err := rand.Read(b)
    if err!=nil{
        return 
    }
    fmt.Println(n)
    fmt.Println(b)
    fmt.Println(hex.EncodeToString(b))

    fmt.Println(sessionId())
}

func sessionId() string{
    b := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return base64.URLEncoding.EncodeToString(b)
} 
