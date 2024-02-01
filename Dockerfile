#use official Golang Image
FROM golang:latest

#Working Directori
RUN mkdir /src/
WORKDIR /src

#Copy the source code
COPY . /src

#Download and Install the Dependencies
RUN go get -d -v ./...
RUN go install -v ./...

#Build the Go App
RUN go build -o go-project

#Expose The PORT
EXPOSE 8000

#Run the excutable
CMD [ "./go-project" ]
