FROM golang:1.20

RUN apt-get update && apt-get -y install vim && apt-get install libgeos-dev -y


#set workdir
WORKDIR /app

#copy code
COPY . .

# build
WORKDIR /app/cmd/main
RUN go build -o echo

# run
CMD ["/app/cmd/main/echo"]