# goDemo

### How to get up and running on local host
- Head over to https://docs.couchbase.com/server/current/getting-started/do-a-quick-install.html and start up your couchbase server using the docker command. Once set up, delete the default travel bucket and make a new one named `cars` then create a primary index key by going to query and running: create primary index on `cars` using gsi;
- clone the repo and run the command `go run main.go`
- On postman, test out the API by sending a PUT to `localhost:12345/car` with the JSON body 
 `
 {
    "name": "Accord",
    "manufacturer": "Toyota",
    "year": "1999"
}
`
- Go back to your couchbaseDB and see the response in the documents section

### How to get up and running with Docker

- Build the image `docker build -t go-web-app:latest .`
- Running  `docker run -it -p 1234:1234 go-web-app` seems to create the container correctly
- As of right now, it cannot talk to the couchbaseDB container
