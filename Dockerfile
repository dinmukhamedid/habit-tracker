FROM ubuntu:latest
LABEL authors="dinmu"

ENTRYPOINT ["top", "-b"]