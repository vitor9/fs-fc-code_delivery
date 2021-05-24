FROM node:12.14.0-alpine3.10

RUN apk add --no-cache bash 

RUN npm install -g @nestjs/cli@7.5.6

RUN npm install mongoose @nestjs/mongoose

RUN npm install @nestjs/config

RUN npm install kafkajs @nestjs/microservices

USER node

WORKDIR /home/node/app