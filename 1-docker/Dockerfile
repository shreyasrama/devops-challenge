FROM node:lts-alpine

WORKDIR /app

COPY yarn.lock ./
COPY package.json ./
COPY ./ ./

RUN yarn install

CMD yarn start