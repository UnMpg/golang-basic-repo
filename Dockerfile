#use official Golang Image
FROM golang:1.20.12-alpine3.19

#Working Directori
WORKDIR /go/src/app

#Copy the source code
COPY . .

#Download and Install the Dependencies
RUN go get -d -v ./...

#Build the Go App
RUN go build -o go-project .

#Expose The PORT
EXPOSE 8000

#Run the excutable
CMD [ "./go-project" ]
