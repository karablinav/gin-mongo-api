REST API with Golang and MongoDB

## Dependencies

Install all required dependencies:

```bash
go get -u github.com/gin-gonic/gin go.mongodb.org/mongo-driver/mongo github.com/joho/godotenv github.com/go-playground/validator/v10
```

## API:

Create user:
```bash
curl --location --request POST 'http://localhost:8080/users/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Marco Reus",
    "location": "Sweden",
    "title": "Software Engineer"
}'
```

Get user by id: 
```bash
curl --location --request GET 'http://localhost:8080/users/63c03f81cfba3a189540572b'
```

Update user info:
```bash
curl --location --request PUT 'http://localhost:8080/users/63c03f81cfba3a189540572b' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Marco Reus!",
    "location": "Sweden",
    "title": "Go Software Engineer"
}'
```

Delete user:
```bash
curl --location --request DELETE 'http://localhost:8080/users/63c03f81cfba3a189540572b'
```



`github.com/gin-gonic/gin` is a framework for building web applications.

`go.mongodb.org/mongo-driver/mongo` is a driver for connecting to MongoDB.

`github.com/joho/godotenv` is a library for managing environment variables.

`github.com/go-playground/validator/v10` is a library for validating structs and fields.

## You may find these resources helpful:

- [Gin-gonic](https://github.com/gin-gonic/gin) - framework for building web applications.
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/) - driver for connecting to MongoDB.
- [Go Validator](https://github.com/go-playground/validator) - library for managing environment variables.
- [Go Environment Loader](https://github.com/joho/godotenv) -  library for validating structs and fields.
- [Golang dto-mapper](https://github.com/dranikpg/dto-mapper) - l