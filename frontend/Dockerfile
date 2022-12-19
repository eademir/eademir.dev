FROM node:18.12.1-alpine3.16 AS build

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

FROM node:18.12.1-alpine3.16

WORKDIR /app

COPY --from=build /app/dist/ ./dist
COPY --from=build /app/package*.json .
RUN npm i --production

EXPOSE 3100

CMD npm run preview

# docker build . -t eademir/blog:frontend 
# docker run -dp 3100:3100 --name frontend eademir/blog:frontend