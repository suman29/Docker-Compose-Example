FROM nginx
EXPOSE 80

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx.conf /etc/nginx/nginx.conf
COPY app.conf /etc/nginx/conf.d/app.conf
