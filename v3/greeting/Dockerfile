FROM golang:latest
RUN mkdir /app
WORKDIR /app 
COPY greeting.bin /app
EXPOSE 8080:8080
EXPOSE 8000:8000
CMD ["/app/greeting.bin"]

