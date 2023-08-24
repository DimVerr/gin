# Password Manager 

Password Manager allows you to create and store your users and credentials in Docker or your own databases with API 

## Installation

1. Download and install [Docker](https://www.docker.com/).

2. Download the repository on your device.

3. Set up your root to the project directory in Dockerfile, db.yaml, and db-test.yaml.

4. (Optional) Setup your own environment variables.

## Usage
1. Launch Docker 

2. There are 2 docker-compose files:
If you want to create both prod and test databases then run 
```golang
docker-compose build
```
If you want to create only prod or only test database then run
``` golang 
    docker-compose build -f db.yaml
```
or 
``` golang 
    docker-compose build -f db-test.yaml
```
3. Run command to run created containers

```golang
docker-compose up
```
The default URL is http://localhost:8080
## Swagger
http://localhost:8080/swagger/index.html 

## Postman

The postman collection file is in the repository, so you can add it to your Postman app and use it.

## CLI 
There is another mine repository with CLI for this project.  
You can find it here https://github.com/DimVerr/password_manager_cli.

## Upgrade 
You can upgrade the project with your own ideas using [GORM](https://gorm.io/) and [GIN-GONIC](https://gin-gonic.com/) libraries:
1. Start from "handlers" directory.

2. Create your_handler.go file with your functionality. 
All DB functions are placed in **config** directory, models in **models** directory, and token functions in **token** directory.

3. Add your function link to main.go file as in the repository.

### Congatulations! App is ready for use 

