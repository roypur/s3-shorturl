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
    var filepath string;
    
    if(runtime.GOOS=="linux"){
        filepath=os.Getenv("HOME")+"/.s3-shorturl/config.json";
    }else if(runtime.GOOS=="windows"){
        filepath=os.Getenv("HOMEPATH")+"\\s3-shorturl\\config.json";
    }
    
    
    file, err := ioutil.ReadFile(filepath);
    
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
