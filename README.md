# golang-mongodb-restful-starter-kit


golang-mongodb-restful-starter-kit is a golang restful api starter kit using mongoDB and JWT. It is very useful for beginners.

### Features!

  - Used mux for routing
  - Used salt to hash the password
  - Used JWT for authentication
  - Follows repository, service and model structure
  - Used interface to hide implementation of repositories and services
  - Error handling
  
### Project structure
```bash
.
├── README.md
├── app
│   ├── handlers                  // API controllers / handlers
│   │   ├── auth.handler.go
│   │   ├── model.handler.go
│   │   └── user.handler.go
│   ├── middleware                // API middleware
│   │   └── cors.go
│   ├── models                    // DB models
│   │   └── user.model.go
│   ├── repositories              // DB repository
│   │   └── user
│   │       └── user.repository.go  // UseRepository
│   └── services                    // Services
│       ├── auth
│       │   └── auth.service.go
│       ├── jwt
│       │   └── jwt.go
│       └── user
│           └── user.service.go
├── config
│   └── config.go                 // Environment configurations, read from .env file
├── db
│   └── mongo.go                  // Mongodb Connection and Session
├── docs                          // Swagger docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod                      // Go modules
├── go.sum
├── main.go
├── routes      // API routes
│   └── api.go
└── utility   //  Contains helpers, basic operations, commons operations, errors, validations                     
    ├── common.go
    ├── errors.go
    ├── handler.go
    ├── role.go
    └── string.go
```
  
### Installation

golang-mongodb-restful-starter-kit requires [Go](https://golang.org/) 1.10+ to run.

Install the dependencies and devDependencies and start the server.

```sh
$ git clone https://github.com/ypankaj007/golang-mongodb-restful-starter-kit.git
$ cd golang-mongodb-restful-starter-kit
$ go run main.go
```
#### Building for source
For production release:
```sh
$ go build
```
  ### Todos

  - Email verification
  - Containerized - Docker + Kubernetes
  - Write Tests
  - Write scripts
  
### Tech

Project uses a number of open source projects to work properly:

* [Go] - Awesome programing language by Google
* [mux] - Implements a request router and dispatcher in Go
* [MongoDB] - document-based, big community, database
* [Redis] - in-memory database using key-value pairs
* [Docker] - Build, Share, and Run Any App, Anywhere
* [Kubernetes] - Automating deployment, scaling, and management of containerized applications


API endpoint - http://localhost:8080
Swagger endpoint - http://localhost:8080/swagger/index.html


   [mux]: <https://www.gorillatoolkit.org/pkg/mux>
   [Go]: <https://golang.org/>
   [MongoDB]: <https://www.mongodb.com/>
   [Docker]: <https://www.docker.com/>
   [Kubernetes]: <https://kubernetes.io/>
   [Redis]: <https://redis.io/>
