# 2) grab apline linux base image
FROM alpine:3.14

# 3) install needed tools
RUN apk add wget \
      && apk add sudo \
      && apk add openrc \
      && apk add fontconfig \
      && apk add freetype \
      && apk add openjdk8 \
      && apk add openjdk8-jre \
      && apk add openjdk8-jre-base \
      && apk add ttf-dejavu \
      && apk add git \
      && apk add xvfb \
      && apk add go

# 3a) update all the packages 
RUN apk update \
      && apk upgrade \
      && rm -rf /var/lib/apt/lists/*

# 4) expose port 8080
EXPOSE 8080

# 5) add and change to new user
ARG USER=myuser
ENV HOME /home/$USER

RUN adduser -D $USER
USER $USER
WORKDIR $HOME

# 6) add an entrypoint instruction to run jenkins
RUN wget https://get.jenkins.io/war-stable/2.303.1/jenkins.war
ENTRYPOINT ["java", "-jar", "jenkins.war"]

# do `$ docker run -p 8080:8080 jenkins:tag` when running the first time
# then do `$ docker start $container_id` to reopen it.