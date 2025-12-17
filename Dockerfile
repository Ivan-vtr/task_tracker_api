FROM ubuntu:latest
LABEL authors="flora"

ENTRYPOINT ["top", "-b"]