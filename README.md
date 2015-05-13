#S3 URL shortener

##Config file
$HOME/.s3-shorturl/config.json

    {
        "region":"amazon-region",
        "bucket":"amazon-bucket",
        "key-id":"aws_access_key_id",
        "secret-key":"aws_secret_access_key",
        "domain":"domain_to_use_in_links"
    }
    
##Command line

    s3-shorturl (optional_keyword) url_to_shorten
