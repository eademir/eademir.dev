FROM --platform=linux/amd64 node:18.12.1-alpine3.16 AS build

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

FROM --platform=linux/amd64 node:18.12.1-alpine3.16

WORKDIR /app

COPY --from=build /app/dist/ ./dist
COPY --from=build /app/package*.json .
RUN npm i --production

EXPOSE 3000

CMD npm run preview

# docker build . -t eademir/blog:panel 
# docker run -dp 3000:3000 --name panel eademir/blog:panel