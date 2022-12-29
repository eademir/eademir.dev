FROM --platform=linux/amd64 node:18.12.1-alpine3.16 AS build

WORKDIR /frontend

COPY ./frontend .

RUN npm install

RUN npm run build

WORKDIR /admin

COPY ./admin .

RUN npm install

RUN npm run build

FROM --platform=linux/amd64 nginx:alpine

RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build /frontend/dist /usr/share/nginx/html/user
COPY --from=build /admin/dist /usr/share/nginx/html/admin
COPY nginx/nginx.conf /etc/nginx/

# docker build . -t eademir/blog:website