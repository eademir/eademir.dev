FROM --platform=linux/amd64 nginx:alpine
RUN rm /etc/nginx/nginx.conf 

COPY nginx.conf /etc/nginx

#docker build . -t nginx 
#docker run --name nginx -p 80:80 -d --network web nginx 