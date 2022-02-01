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

## Run With Makefile:
- Run The Project Without Building:
```sh
$ make run
```
This Is Also The Default Target Which Will If No Argument Is Provided To Make

- Build The Project:
```sh
$ make build
```
This Would Create A Directory `bin` In The Project Root And Build The Resulting Executable There

- Run The Project In A Docker Container:
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

## Run With docker-compose :
- Running The Project :
```sh
$ docker-compose up
```
This Will Build The Necessary Image On The First Run

- Bringing Down The Project :
```sh
$ docker-compose down --volumes && docker image rm -f area99_web
```


## Running Manually With Go:
- Running The Project :
```sh
$ go build . && ./assessment
```
