FROM busybox

COPY ./bin/docker-services /docker-services

ENTRYPOINT ["./docker-services"]
