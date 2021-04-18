FROM node:lts-alpine

WORKDIR /app

COPY yarn.lock ./
COPY package.json ./
COPY ./ ./

RUN addgroup testgroup && adduser -DH -G testgroup shreyas
RUN chown -R shreyas:testgroup /app
USER shreyas

RUN yarn install

CMD yarn start