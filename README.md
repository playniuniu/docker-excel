# Docker for Excel

Use excel to create docker container

### Run

```shell
docker run --name docker-excel --restart always -v /var/run/docker.sock:/var/run/docker.sock -d -p 9001:9001 playniuniu/docker-excel
```

### Excel

demo.csv

```csv
name,image,version,port,env,remove
web,nginx,alpine,8080:80,,true
database,mysql,5.7,3306:3306,MYSQL_ROOT_PASSWORD=123456,true
```

**Note 1:** Only support one port pair

**Note 2:** env value is split by `&`

### Develope

Prepare go package

```shell
go get -u github.com/gin-gonic/gin
go get -u github.com/gabriel-vasile/mimetype
go get -u github.com/docker/go-connections/nat
go get -u golang.org/x/sys/unix
```

Fix go client bug

```shell
rm -rf ${GOPATH}/src/github.com/docker/docker/vendor/github.com/docker/go-connections/nat
```

Make executable file

```shell
make
```

Make docker

```shell
make docker
```

### Reference

1. [docker client bug](https://github.com/moby/moby/issues/28269)
2. [how to expose port](https://medium.com/backendarmy/controlling-the-docker-engine-in-go-d25fc0fe2c45)
