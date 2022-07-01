# Go Mocker
A way to make a simple and flexible API to validate and test integrations.

## How to use

### Naked mode
When running the program without providing a configuration you will get an API that listens to port
and logs the requests

### Structure config
Setup the configuration inside a folder with the api path structure that you are testing (quite
literally you setup `/api/v1/content.yaml` to mock a `/api/v1/content` endpoint)


## TODO
- [] read from files
- [] method comes from the file name (`_FILE_._METHOD_.yaml`)
- [] file path becomes the endpoint prefix
- [] path parameters 
- [] use go template to format the response based on the input
- [] support for relational database (?)

