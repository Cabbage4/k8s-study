FROM centos

RUN mkdir /app
COPY main /app

WORKDIR /app
CMD ["/app/main"]