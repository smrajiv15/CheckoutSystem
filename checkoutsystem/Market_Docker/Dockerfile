#Container for running farmer Market Problem
FROM centos:7
MAINTAINER smrajiv15@gmail.com

#installing needed utilities to create Golang environment
RUN yum install wget -y
RUN yum install git -y

#Golang development setup
RUN wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
RUN mkdir -p /home/Go_Project/src
RUN mkdir -p /home/Go_Project/bin
RUN mkdir -p /home/Go_Project/pkg

ENV GOPATH=/home/Go_Project
ENV GOBIN=/home/Go_Project/bin
ENV PATH="${PATH}:/usr/local/go/bin"


#store the project in appropriate location
RUN git clone https://github.com/rajiv15/MarketProblem.git /home/Go_Project/src

#install Mongo driver for Golang
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson

#Make aliases to run it user friendly
RUN echo 'alias market="go run /home/Go_Project/src/checkoutsystem/main.go"' >> ~/.bashrc
RUN echo 'alias marketTest="go test checkoutsystem/checkout -v"' >> ~/.bashrc
