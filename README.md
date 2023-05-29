## Go Book Service

### Database setup in Docker

Before running service, please create network bridge and volume container.

Create volume container (e.g. books):

```
docker volume create [container name]
```

Create bridge network (e.g. booknet):

```
docker network create -d bridge [network name]
```
