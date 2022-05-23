FROM golang
ENV GO111MODULE=on
WORKDIR /app

# Add the rest of the files
COPY ./api/. /app/.

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
EXPOSE 5002

RUN go build ./main.go
ENTRYPOINT /app/main

#docker build -t lib-api . --load 
#docker run -d --rm -p 5002:5002 --name lib-api-1 lib-api