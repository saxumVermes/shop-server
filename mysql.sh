#!/bin/bash
source ~/.zshrc
if [[ -z $1 ]]; then
  echo 'Specify a port number'
  exit 1
fi
mkdir -p /var/lib/mysql
docker run -d \
  --name shop-database \
  -v $PWD/var/lib/mysql:/var/lib/mysql/ \
  -p $1:3306 \
  -e MYSQL_USER=user \
  -e MYSQL_PASSWORD=mysql \
  -e MYSQL_DATABASE=shop \
  -e MYSQL_ROOT_PASSWORD=mysql \
  mysql/mysql-server:5.7.23-1.1.7
