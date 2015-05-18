package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "encoding/json"
    "runtime"
)

var config map[string]string;

func getConf()(bool){
    var home string;
    
    if(runtime.GOOS=="linux"){
        home=os.Getenv("HOME")
    }else if(runtime.GOOS=="windows"){
        home=os.Getenv("HOMEPATH");
    }
    
    
    file, err := ioutil.ReadFile(home + "/.s3-shorturl/config.json");
    if(err==nil){
        config = make(map[string]string);
        err = json.Unmarshal(file, &config);
            if(err!=nil){
                fmt.Println(err);
                return false;       
            }else{
                return true;
            }
    }else{
        fmt.Println(err);
    }
    return false;
}
