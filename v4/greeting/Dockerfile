FROM golang:latest
RUN mkdir /app
WORKDIR /app 
COPY greeting.bin /app

RUN mkdir /app/certs
COPY certs/server.crt /app/certs/server.crt
COPY certs/server.key /app/certs/server.key

# We have to run following commands as we're using self-sign ca
RUN cp /app/certs/server.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

EXPOSE 8080:8080
EXPOSE 8000:8000
CMD ["/app/greeting.bin"]
