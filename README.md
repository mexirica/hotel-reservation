# Hotel reservation backend

This is a complete Web API project in Go (Golang) using the Fiber framework to simulate a hotel reservation service. The main goal of this project is to demonstrate how to create a robust and efficient web application for managing hotel reservations, including functionalities such as creating, querying, updating, and canceling reservations.

The project is structured as follows:

Environment Setup: Before starting development, it is necessary to set up the development environment with Go and install the Fiber framework. This can be done by following the official installation steps available in the Go and Fiber documentation.

Project Structure: The project follows an organized structure, separating responsibilities into different packages, such as handlers for handling HTTP requests, models for defining data structures, and services for implementing business logic.

Data Modeling: For this hotel reservation service, the main models involved are Reservation, which contains information about the reservation, such as check-in date, check-out date, number of guests, and Hotel, which stores details about the hotel, like name, location, and capacity.

Business Logic Implementation: The business logic is implemented in the services, where rules such as validation of reservation dates, checking hotel availability, and calculating prices based on the stay period are handled.

Testing: The project includes unit tests to ensure the correctness of functions and integration of components. The tests cover both the business logic and the API handlers.

Deployment: The project can be easily deployed in production environments, supporting Docker containers to simplify the deployment process and scalability.

This project serves as a practical example of how to develop a complete web application in Go using Fiber, covering everything from initial setup to deployment, through data modeling, endpoint implementation, business logic, and testing.

# Project environment variables
```
HTTP_LISTEN_ADDRESS=:3000
JWT_SECRET=somethingsupersecretthatNOBODYKNOWS
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

## Project outline
- users -> book room from an hotel 
- admins -> going to check reservation/bookings 
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
### Mongodb driver 
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber 
Documentation
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v3
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -p 27017:27017 -d mongo:latest 
```
