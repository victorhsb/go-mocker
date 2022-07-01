# Go Mocker
A way to make a simple and flexible API to validate and test integrations.

## How to use
Setup the configuration inside a folder with the api path structure that you are testing (quite
literally you setup `/api/v1/content.get.yaml` to mock a `GET /api/v1/content` endpoint)

## TO-DO
- [ ] naked mode (for running without configuration)
- [x] read from files
- [x] method comes from the file name (`_FILE_._METHOD_.yaml`)
- [ ] file path becomes the endpoint prefix
- [ ] path parameters 
- [ ] use go template to format the response based on the input
- [ ] support for relational database (?)

