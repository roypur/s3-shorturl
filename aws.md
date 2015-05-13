##Overview

In outline, all of the AWS SDKs now use a standard approach for how to manage credentials. 
This includes the AWS command-line interface (CLI) and the AWS Tools for Windows PowerShell. To configure credentials, you can do this:

Keep them in a set of environment variables that have standard names for all SDKs. The .NET SDK is an exception here, as I'll explain shortly.
Keep them in a central credential file that is shared by all the SDKs. And in that credentials file, now you can
use profiles. You can keep multiple sets of credentials in the same credentials files using different profile names.
When you initialize a connection to AWS in code, you can specify the set of credentials you want to use.

Some of these options were available previously, but differed between SDKs. For example, different SDKs used different environment variable names, 
and the location and format of credentials files differed between SDKs. 
Now, with only one exception, you can choose any of these approaches, and whatever you do will work for all the SDKs. 
Once you've configured one SDK or the CLI, you can use the same credentials for any other.
Environment Variables

All the SDKs except the .NET SDK now can automatically look for credentials in the same environment variables: AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY. 
###Credentials File and Profiles

Instead of keeping credentials in environment variables, you can now put credentials into a single file that's in a central location. The default location is this:

    ~/.aws/credentials (Linux/Mac)
    %USERPROFILE%\.aws\credentials  (Windows)

An important point is that the default location for the credentials file is a user directory. It's no longer part of a project file structure, 
such as an app.config file (.NET) or .properties file (Java). 
This can enhance security by allowing you to keep the credentials in a location that's accessible only to you, 
and it makes it less likely that you'll inadvertently upload credentials if you upload a project to a developer sharing site like GitHub.

The format for the credentials is the same for all the SDKs and the AWS CLI:

    [default]
    aws_access_key_id = ACCESS_KEY
    aws_secret_access_key = SECRET_KEY

As noted, you can keep multiple sets of credentials in the same file, identifying each set using a profile name, like the following example. I'll discuss how to use these profiles shortly.

    [default]
    aws_access_key_id = ACCESS_KEY
    aws_secret_access_key = SECRET_KEY
 
    [Alice]
    aws_access_key_id = Alice_access_key_ID
    aws_secret_access_key = Alice_secret_access_key
 
    [Bob]
    aws_access_key_id = Bob_access_key_ID
    aws_secret_access_key = Bob_secret_access_key

For additional security of the credentials file, you can set the file's permissions to make sure that only the owner is allowed to access the file. 
In Linux, use chmod 600 to set owner-only permissions. In Windows, use the Properties window or use the icacls command.
