#S3 URL shortener

##Config file
$HOME/.s3-shorturl/config.json

    {
        "region":"amazon-region",
        "bucket":"amazon-bucket",
        "domain":"domain_to_use_in_links",
        "profile":"aws-credentials-profile"
    }
    
[http://9f.no/ee0a](Configure amazon credentials)
    
##Command line

    s3-shorturl (optional_keyword) url_to_shorten
