package main;

import (
    "github.com/awslabs/aws-sdk-go/service/s3"
    "github.com/awslabs/aws-sdk-go/aws"
    "github.com/awslabs/aws-sdk-go/aws/credentials"
    "os"
    "fmt"
)
var auth *aws.Config;

func main(){
    if(getConf()){
        
        var cred *credentials.Credentials;
        
        if(len(config["profile"])>0){
    
            cred = credentials.NewChainCredentials(
        
            []credentials.Provider{
                &credentials.EnvProvider{},
                &credentials.SharedCredentialsProvider{Filename: "", Profile: config["profile"]},
                &credentials.EC2RoleProvider{},
            })
            auth = &aws.Config{Region: config["region"], Credentials: cred};
        }else{
            auth = &aws.Config{Region: config["region"]};
        }
        
        
        
        var exit bool = false;
        
        var dialog string;
        
        var key string;
        if(len(os.Args)==2){
            for(!exit){
        
                key = makeKey();
        
                if(exists(key)){
                    exit = false;
                }else{
                    exit = true;
                }
            }
            set(key, os.Args[1]);
            fmt.Println("\nhttp://"+config["domain"]+"/"+key+"\n");
        }else if(len(os.Args)==3){
            if(exists(os.Args[1])){
                fmt.Println("\nThe given key exists. Do you want to replace it? ( YES / NO )\n")
                fmt.Scanln(&dialog);
                if(dialog=="yes" || dialog=="YES" || dialog=="Y" || dialog=="y"){
                }else{
                    return
                }
            }
            set(os.Args[1], os.Args[2]);
            fmt.Println("\nhttp://"+config["domain"]+"/"+os.Args[1]+"\n")
        }
    }
}

func set(key string, url string){
    svc := s3.New(auth);
    
    params := &s3.PutObjectInput{
        Bucket:         aws.String(config["bucket"]),
        Key:            aws.String(key),
        ACL:            aws.String("public-read"),
        WebsiteRedirectLocation: aws.String(url),
    }
    _,err := svc.PutObject(params)

    if awserr := aws.Error(err); awserr != nil {
        // A service error occurred.
        fmt.Println("Error:", awserr.Code, awserr.Message)
    } else if err != nil {
        // A non-service error occurred.
        panic(err)
    }
}

func exists(key string)(bool){
    svc := s3.New(auth);
    params := &s3.HeadObjectInput{
        Bucket:               aws.String(config["bucket"]), // Required
        Key:                  aws.String(key),  // Required
    }
    _,err := svc.HeadObject(params)
    if(err!=nil){
        return false;
    }
    return true;
}
