# Learning Redis with Golang

### Uses :-
- Redis package being used - **go-redis**
- REST Framework - **Fiber**
- Database - **MySQL**
- ORM - **GORM**

### How to run :-

- Run **Redis** service, Standalone installation OR the container way (`docker run -d -p 6379:6379 redis/redis-stack`)
- You can connect to Redis CLI using command `redis-cli`, where you can check the data being cached.
- Install the project dependancies.
- From the project root folder, Run `go run main.go`
- Start hitting the REST endpoints through some REST client like Postman.
