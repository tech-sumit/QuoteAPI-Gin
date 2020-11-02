# Quote API 📜
## Gin API with CURD 🍵
### Development tools version & prerequisites
    Docker version 19.03.13, build 4484c46d9d
    docker-compose version 1.27.3, build 4092ae5d
    GNU Make 4.1
    go version go1.14.8 linux/amd64
### Postman collections
- Import ```postman/QuoteApi.postman_collection.json``` in postman for request presets
- For request documentation =>  https://documenter.getpostman.com/view/7867258/TVYM4vWQ

### ENVs 
Create ```.env``` with 

    # Also try=> release
    GIN_MODE=debug
    
    # No perticular address (access anywhere)
    SERVER_HOST=0.0.0.0
    # Server port
    PORT=7000
    
    # for basic auth
    AUTH_USERNAME=admin
    AUTH_PASSWORD=password
    
    # Docker container host or ip address of mongodb server
    MONGO_HOST=mongo 
    # db port
    MONGO_PORT=27017
    # DB Name
    MONGO_DATABASE=quotesdb
    # Collection name
    MONGO_COLLECTION=api_quotes
    
    ## For mongodb credentials If configured
    # MONGO_USER=username
    # MONGO_PASSWORD=password

### Docker start
    $ make start_docker
### Docker-Compose start
    $ make start_docker_compose
### Start local
    $ make start_local
