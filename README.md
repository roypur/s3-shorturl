#S3 URL shortener

##Config file
$HOME/.s3-shorturl/config.json

    {
        "region":"amazon-region",
        "bucket":"amazon-bucket",
        "domain":"domain_to_use_in_links",
        "profile":"aws-credentials-profile"
    }
    
[Configure Amazon Credentials](http://github.com/roypur/s3-shorturl/blob/master/aws.md)
    
##Command line

    s3-shorturl (optional_keyword) url_to_shorten
