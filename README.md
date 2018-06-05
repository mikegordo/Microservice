# Microservice demo
Just playing with microservices in Golang.

Three containers share the same network. They work together to produce the result.

## Generator
Generator can generate some response based on request HTTP POST parameters in JSON format:
- `"length"` (integer) 
- `"caps"` (boolean)

### Building container
`docker build -t mikegordo/generator generator/`

### Starting container
`docker run -p 8080:8080 --rm --name generator mikegordo/generator`

### Testing the app
`curl -i -X POST -d "{\"length\":10, \"caps\":false}" localhost:8080`

### Response example
    {
        "Status":true,
        "ErrorCode":200,
        "Value":"idcgkvepqs"
    }

## Fetcher
Fetches the response from the Generator every `FETCHER_FREQ` seconds (default 5) and stores it to Redis.
Uses environment variables `FETCHER_LENGTH` (default 16) and `FETCHER_CAPS` (default false).

### Building container
`docker build -t mikegordo/fetcher fetcher/`

### Starting container
*First*, start Generator, then Redis container.

`docker run --name redis --net container:generator --rm redis`

Start the app

`docker run --rm --name fetcher --net container:generator mikegordo/fetcher`

Example with variables

`docker run --rm --name fetcher --net container:generator -e FETCHER_LENGTH=8 -e FETCHER_CAPS=true mikegordo/fetcher`
