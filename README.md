# Prerequisites
1. Add **example.env** file for set environment variable for run test api on local.
    - PORT: use for assign the running port of api.
    - SIGN: use for sign/verify JWT.
    - DB_CONN: use for connect to database(MySQL)
2. **MySQL** for use as database of api.
3. **Docker** for run service image.

# Command
build-app: ```go build -ldflags "-X main.buildcommit=`git rev-parse --short HEAD` -X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" -o app```

run-maria: ```docker run --detach -p 127.0.0.1:3306:3306 --name some-mariadb --env MARIADB_ROOT_PASSWORD=my-secret-pw --env MARIADB_DATABASE=my-app mariadb:latest```

build-docker-image: ```docker build -t todo:test -f Dockerfile .```

run-container: ```docker run -p 8081:8081 --env-file ./example.env --link some-mariadb:db --name myapp todo:test```