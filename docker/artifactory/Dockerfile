# 1) grab apline linux base image
FROM alpine:3.14.0

# 2) Log in as Root
USER root 

# 3) install needed tools
RUN apk add wget \
      && apk add bash \
      && apk add rpm \
      && apk add net-tools \
      && apk add rsync \
      && apk add openjdk8 \
      && apk add openjdk8-jre-base 

# 4) update packages 
RUN apk update \
      && apk upgrade \
      && rm -rf /var/lib/apt/lists/*

# 4) expose port 8081 + 8082
EXPOSE 8081
EXPOSE 8082

# 6) Wget artifactory.zip file and unzip 
RUN wget https://repo.jfrog.org/artifactory/oss-releases-local/org/artifactory/oss/jfrog-artifactory-oss/6.10.0/jfrog-artifactory-oss-6.10.0.zip
RUN unzip jfrog-artifactory-oss-6.10.0.zip

# 7) Make needed jfrog directory under /opt and move artifactory
RUN mkdir /opt/jfrog
RUN mv artifactory-oss-6.10.0 artifactory
RUN mv artifactory /opt/jfrog
RUN export JFROG_HOME=/opt/jfrog
WORKDIR /opt/jfrog/artifactory/bin

# 8) start artifactory
ENTRYPOINT ["./artifactoryctl"]

# do `$ docker run -p 8081:8081 -p 8082:8082 image_name` when running the first time
# Log in with http://localhost:8081/artifactory/webapp/#/home
