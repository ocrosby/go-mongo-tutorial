# MongoDB Tutorial

## Running Mongo Locally in a Docker Container

Pulling the image from Docker Hub

```sh
docker pull mongo:latest
```

Running the image

```sh
docker run -d -p 27017:27017 --name=mongo-example mongo:latest
```

Listing running containers

```sh
docker ps
```

Accessing the MongoDB shell in the running container

```sh
docker exec -it mongo-example mongosh
```

Stopping a container

```sh
docker stop <container-id>
```


## References

- [Official Docker Image for MongoDB](https://hub.docker.com/_/mongo)
- [How to Use GO with MongoDB](https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver) 
- [MongoDB Go Driver Tutorial](https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial)
- [MongoDB Go Tutorial](https://github.com/tfogo/mongodb-go-tutorial)