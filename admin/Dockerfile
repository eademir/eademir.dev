FROM node:lts AS build

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build

FROM nginx:stable-alpine
COPY --from=build /app/public /usr/share/nginx/html
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf