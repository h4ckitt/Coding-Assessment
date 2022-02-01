# What Is This??
This Is My Submission For The Area99 Golang Backend Developer Role To Build An API Service For An
Automotive Car Industry

# Decisions Made Outside The Instruction Scope
* Added Docker Support And A Makefile To Immensely Ease Getting The Project Up And Running.

# Let's Run It

## Pre-Requisites
- go 1.16+
- docker
- docker-compose

## Things To Note
- The Application Can Work In Either Of Two Modes (GRPC/REST) Which Can Be Set In The .env File
- Before Running In Docker, The `DATABASE_HOST` Entry in the .env Has To Be Changed To `db` And To `localhost` When Running Directly From Terminal. 

## Run With Makefile
- Run The Project Without Building
```sh
$ make run
```
This Is Also The Default Target Which Will If No Argument Is Provided To Make

- Build The Project
```sh
$ make build
```
This Would Create A Directory `bin` In The Project Root And Build The Resulting Executable There

- Run The Project In A Docker Container
```sh
$ make docker
```
This Approach Uses `docker-compose` To Build And Run Two Docker Containers With The Application And The
Database In Them Respectively.

- Run Tests
```sh
$ make test
```
This Runs All The Tests In All The Packages With Coverages Where Necessary

- Clean Up Created Docker Images And Volumes Pertaining To This Project:
```sh
$ make clean
```

## Run With docker-compose
- Running The Project
```sh
$ docker-compose up
```
This Will Build The Necessary Image On The First Run

- Bringing Down The Project
```sh
$ docker-compose down --volumes && docker image rm -f area99_web
```


## Running Manually With Go
- Running The Project
```sh
$ go build . && ./assessment
```
# Making Requests
## REST
The REST API URL For Making Requests Is : `http://localhost:8080/v1/car`

- To Register A Car
```sh
$ curl -X POST http://localhost:8080/v1/cars -d \
'{
"name": "Nissan Leaf",
"type": "sedan",
"color": "green",
"speed_range": 160,
"features": ["sunroof", "surround-system"]
}'
```

- To View A Car's Details
The Car's ID Needs To Be Known
```sh
$ curl http://localhost:8080/v1/cars/1
```

- To Get A Car By Type
```sh
$ curl http://localhost:8080/v1/cars?type=sedan
```

- To Get A Car By Color
```sh
$ curl http://localhost:8080/v1/cars?color=blue
```

## GRPC
Although I Couldn't Find A Way To Test My Server With Insomnia Nor Postman, I Found [grpcox](https://github.com/gusaul/grpcox) Helpful.

The .proto File Is Located In `adapter/grpc/grpc_proto`

The Methods On The GRPC Server Are:
- ViewCarDetails: This Takes An int32 Integer
- Register: 
```
{
  "Name": "",
  "Type": "",
  "Color": "",
  "SpeedRange": 0,
  "features": [
    ""
  ]
}
```
- GetCarsByColorOrType:
```
{
 "Color": "",
 "Type": "" 
}
```
This Is A OneOf Struct And Only One Of Color And Type Can Be Set At The Same Time.
Setting Both Sends Only The Type Variable