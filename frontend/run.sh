#!/bin/bash
setup() {
  if [[ -z $1 ]]; then
    echo 'Specify a port!'
    exit
  fi
  docker build -t shop-frontend . && \
  docker run -d -p $1:80 --name shop-frontend shop-frontend
  if [[ ! $? = 0 ]]; then
    exit
  fi
}

tear-down() {
  docker stop shop-frontend
  docker rm shop-frontend
  docker rmi shop-frontend
}

case $1 in
  'setup' )
  setup $2
    ;;
  'tear-down' )
  tear-down
    ;;
  * )
  echo 'Usage: ./run.sh setup <port>|teardown'
    ;;
esac
