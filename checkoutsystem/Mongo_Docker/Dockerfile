#Container for running farmer Market Problem
FROM centos:7
MAINTAINER smrajiv15@gmail.com

#install MongoDB
RUN echo -e "[mongodb]\n\
name=MongoDB Repository\n\
baseurl=http://downloads-distro.mongodb.org/repo/redhat/os/x86_64/\n\
gpgcheck=0\n\
enabled=1" >> /etc/yum.repos.d/mongodb-org.repo

RUN yum -y update
RUN yum -y install mongodb-org mongodb-org-server
EXPOSE 27017
RUN mkdir -p /data/db
ENTRYPOINT ["/usr/bin/mongod"]
