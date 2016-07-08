# docker-services

`docker-services` is a tool which allows you to define and run dockerized applications. `docker-compose` allows you to define containers required to run fully functional service and `docker-services` allows to define what services you want to run.

# Why ?

If you are adopting microservices architecture it means you have a certain number services. Sum of these services builds up your platform. Setting up development, testing environments only with `docker-compose` is a tedious task. If you need to setup tens of services you will end up with some scripts that will help you to launch them. `docker-services` makes this task trival.

# How it works ?

