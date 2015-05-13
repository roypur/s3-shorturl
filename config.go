package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "encoding/json"
)

var config map[string]string;

func getConf()(bool){
    file, err := ioutil.ReadFile(os.Getenv("HOME") + "/.s3-shorturl/config.json");
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
