FROM nginx:1.15-alpine

COPY . /usr/share/nginx/html/
RUN sed -i "21i\\\n    error_page 405 =200 \$uri;" /etc/nginx/conf.d/default.conf
