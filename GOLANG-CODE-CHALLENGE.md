# Code Challenge
The automotive industry wants to build an API Service. 

They want to register their cars, list and view them.

---

## Car Specifications

| Specifications | Type | Example
| :---: | :---: | :---: |
| Type | One of (Sedan, Van, Suv, motor-bike) | Suv 
| Name | any possible name | Mercedes benz X460 
| Color | One of (red, green, blue) | red
| Speed Range (km) | between 0 - 240 | 25
| Features | Many of (sunroof, panorama, auto-parking, surround-system) |  sunroof, auto-parking

## Write your code
1. Define your own Data structure "schema" and should be saved in Postgresql DB. please feel free to split your data structure into multiple tables if this sounds better for you.
2. API should use JSON or ProtoBuf (GRPC) serialization
3. Create an API endpoint that can register a car
4. Create an API endpoint to view a car details
5. Create an API endpoint to filter cars by color and/or type (or any other filter)

## Guidelines
* Write your code in Golang, no frontend or HTML needed, just backend code.
* Use Postman or BloomRPC  to try your API locally
* Design the API and responses correctly and take care of the response messages, errors, HTTP response status codes.
* Design the database properly and use PostgreSQL functions to keep a timestamp of each record, when it was created and updated (create_time, last_updated).
* Use the correct data types and Data structures and validate data when needed.
* Build the API to support REST (mandatory) and gRPC (if possible, but not mandatory).
* Cover your code with tests, unit tests, functional / integration tests, show us the best way to implement tests for this kind of application.
* Improvise: make your own assumptions about any missing information from the assignment according to your understanding of programming and software development best practices.
* Document your solution and how to build and run the application locally with a simple Readme file.

---
## Finally
* Once you have completed the challenge. Please create a repository on Github and send us a link to your repository.
* Tell us how long it took you to implement this codebase.
Sed the details (link) of your repository to :
anna@area99.com, zaid@area99.com, almasry@area99.com
---
Good luck

Happy coding ;)