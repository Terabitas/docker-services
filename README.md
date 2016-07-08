# docker-services

`docker-services` is a tool which allows you to define and run dockerized applications. `docker-compose` allows you to 
define containers required to run fully functional service and `docker-services` allows to define what services you want to run.

# Why ?

If you are adopting microservices architecture it means you have a certain number services. 
Sum of these services builds up your platform. Setting up development, testing environments only with `docker-compose` is a tedious task. 
If you need to setup tens of services you will end up with some scripts that will help you to launch them. `docker-services` makes this task trivial.

# How it works ?

1) Install `docker`
2) Install `docker-compose`
3) Install `docker-services`

4) Create `docker-services.yml` file and define what services you want to run (use these sample ones to test this tool).

```
dir: ~/dcs
host: github.com
services:
    "nildev/auth":
        branch: master
    "nginxinc/docker-nginx":
        branch: master
```

Now setup and start services:

```
docker-services setup
docker-services start
```

`docker-services` will download `docker-compose.yml` from defined repositories and will use `docker-compose` to start it.

# How to make my services compatible with `docker-services`

There are only two things required for that. First of all in root directory of your repository create `docker-compose.yml` file.
Secondly make sure that all services defined in `docker-compose` has `container_name` key defined. 

## How should I name containers 

Your `docker-compose.yml` will contain service that should be publicly accessible by other services. This is what we call `main` service. 
It could be some REST API app for example. This service most likely will require some internal services in order to function, for example database.
These additional services we call `internal` services. 

`main` service `container_name` key should have value of your git repository name. For example:
  
```
github.com/nildev/auth.git - nildev-auth
github.com/nginxinc/docker-nginx.git - nginxinc-docker-nginx
```

And all `internal` services should append `container_name` key with `main` service name:

```
mysql.nildev-auth
mongodb.nginxinc-docker-nginx
```

For example this is how `docker-compose.yml` could look for a service:
```
version: '2'
services:
    nildevAuth:
        image: nildev/auth:$DOCKER_SERVICES_TAG
        container_name: "nildev-auth"
    nildevMysql:
        image: nildev/mysql:$DOCKER_SERVICES_TAG
        container_name: "mysql.nildev-auth"
```