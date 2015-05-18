package main

import (
    "encoding/hex"
    "crypto/rand"
    "fmt"
)

func makeKey()(key string){

    b := make([]byte, 2);
    _, err := rand.Read(b);
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    
    key = hex.EncodeToString(b);
    
    return key;
}
