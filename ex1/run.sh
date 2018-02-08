#!/bin/sh

NORMAL="\\033[0;39m"
#RED="\\033[1;31m"
BLUE="\\033[1;34m"

IMAGE_NAME="docker-test1"
CONTAINER_NAME="test1"

log() {
  echo "$BLUE \nYou are inside the container: \"$1\""
  echo "$NORMAL\n"
}

bash() {
  log $CONTAINER_NAME 
  docker run -it --rm --name $CONTAINER_NAME $IMAGE_NAME /bin/bash
}

error() {
  echo ""
  echo "$RED >>> ERROR - $1$NORMAL"
}

build() {
  docker build -t $IMAGE_NAME .
  [ $? != 0 ] && \
    error "Docker image build failed !" && exit 100
}

attach() {
  docker attach $CONTAINER_NAME
}

stop() {
  docker stop $CONTAINER_NAME
}

start() {
  docker start $CONTAINER_NAME
}

remove() {
  echo "$BLUE\nRemoving the container: \"$CONTAINER_NAME\"" 
  echo "$NORMAL"
  docker rm -f $CONTAINER_NAME &> /dev/null || true 
}

run() {
  log $CONTAINER_NAME
  docker run -it --name $CONTAINER_NAME $IMAGE_NAME 
}

help() {
  echo "-----------------------------------------------------------------------"
  echo "                      Available commands                              -"
  echo "-----------------------------------------------------------------------"
  echo "$BLUE"
  echo "   > build  - To build the Docker image"
  echo "   > stop   - To stop container"
  echo "   > start  - To start container"
  echo "   > run    - Log you into container"
  echo "   > bash   - Log you into bash shell inside container"
  echo "   > remove - Remove container"
  echo "   > help   - Display this help"
  echo "$NORMAL"
  echo "-----------------------------------------------------------------------"
}

$*

