# Hotel Reservation System

## Project Outline
- users -> book room from an hotel
- admins -> going to check reservation/bookings
- Authentication and Authorization -> JWT tokens
- hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management ->  seed data, migrations

### Mongodb driver
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing the mongodb driver
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber
Documentation
```
https://gofiber.io/
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```

