# Quote API
## Gin API with CURD 🍵
### Development tools version & prerequisites
    Docker version 19.03.13, build 4484c46d9d
    docker-compose version 1.27.3, build 4092ae5d
    GNU Make 4.1
    go version go1.14.8 linux/amd64
### Postman collections
- Import ```postman/QuoteApi.postman_collection.json``` in postman for request presets
- For request documentation  

### ENVs 
Create ```.env``` with 
    
    GIN_MODE=debug # also try=> release
    SERVER_HOST=0.0.0.0 # No perticular address (access anywhere)
    PORT=7000 # Server port
    AUTH_USERNAME=admin # for basic auth
    AUTH_PASSWORD=password # for basic auth
    MONGO_HOST=mongo # Docker container host or ip address of mongodb server
    MONGO_PORT=27017 # db port
    MONGO_DATABASE=quotesdb # DB Name
    MONGO_COLLECTION=api_quotes # Collection name
    
    ## For mongodb credentials If configured
    # MONGO_USER=username
    # MONGO_PASSWORD=password

### Docker start
    $ make start_docker
### Docker-Compose start
    $ make start_docker_compose
### Start local
    $ make start_local
