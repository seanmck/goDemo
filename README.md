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

To network two containers together, they must be know each other's IP
address or be able to use a known hostname. To map container names to
hostnames, we can use a user-defined Docker network.

- Build the image `docker build -t go-web-app:latest .`
- Create a new bridge network `docker network create --driver bridge
  couchbase`
- Stop and remove the previously started couchbase container. Restart
  it, but add `--network couchbase` to the quick start command.
- Run  `docker run -it -p 12345:12345 --network couchbase go-web-app`.
- The Couchbase container has the name "db" which can be used as a
  hostname by the go-web-app pod. go-web-app only has a random ID as its
  hostname, but we don't need to reach it from inside the bridge. It's
  reachable at localhost:12345 on the host.


### How to get up and running on AKS
- Create your AKS Cluster
- https://docs.couchbase.com/operator/current/install-kubernetes.html
- Must follow these instructions if you create a new cluster ^ You have to run those commands wherever you downloaded those files
- `tar -xvzf couchbase-autonomous-operator-kubernetes_2.1.0-linux-x86_64.tar.gz` unzip and then cd into `cd couchbase-autonomous-operator-kubernetes_2.1.0-linux-x86_64`
- You can then follow https://docs.couchbase.com/operator/current/install-kubernetes.html
- Connect to your cluster
- Deploy the couchbase.yaml and the deployment.yaml
- Run `kubectl port-forward service/cb-example 2121:8091` to access the DB UI
- Create your Primary key on `cars`
- Have fun with postman :)

### How to get up and running with GH Actions 
- When you have a new cluster follow these steps https://docs.microsoft.com/en-us/azure/dev-spaces/how-to/github-actions put them into secretes and you are good to go
