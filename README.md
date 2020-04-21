# HTTP Test Server

Simple HTTP test server for debugging. Prints raw request to stdout, and returns OK.

## Usage

```sh
docker run -p 9000:80 jdharmon/http-test
```

```sh
curl -v -X POST -H 'Content-Type: application/json' --data '{ \"text\": \"Hello!\" }' http://localhost:9000/api/method
```