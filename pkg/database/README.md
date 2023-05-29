## Database in Docker

### Setup

Before running service, please create network bridge and volume container.

Create volume container (e.g. books):

```
docker volume create [container name]
```

Create bridge network (e.g. booknet):

```
docker network create -d bridge [network name]
```

### Configuration

To use database, you need to configure dabatase. Otherwise you may observe errors such as "Authentication error" or "ERROR: relation '[Table name]' does not exist".

Log into the database:

docker exec -it [volume container name] ./cockroach sql --insecure

and create database and user:

```
CREATE DATABASE [db name];
CREATE USER [user name];
GRANT ALL ON DATABASE [db name] TO [user name];
quit
```
