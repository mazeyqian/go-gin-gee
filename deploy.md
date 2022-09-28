Tokyo-CentOS-9-x64

ssh-keygen -t rsa -C "mazey@mazey.net"

cd /

mkdir web

cd web

git clone git@github.com:mazeyqian/go-gin-gee.git

yum -y update

yum install nginx -y

nginx

chkconfig nginx on

curl http://127.0.0.1

cd /etc/nginx

git clone git@github.com:mazeyqian/feperf.com.conf.d.git

cd ~

yum install golang -y

go version

cd /web/go-gin-gee

go run scripts/init/main.go

go run cmd/api/main.go

GOOS=linux GOARCH=amd64 go build -o dist/api cmd/api/main.go

cd ~

yum install epel-release -y

yum install supervisor -y

systemctl start supervisord

systemctl status supervisord

systemctl enable supervisord

ps -ef|grep supervisord

cd /etc/supervisord.d

nano api.ini

[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api
autostart=true
autorestart=true
stderr_logfile=/web/go-gin-gee/log/api.err
stdout_logfile=/web/go-gin-gee/log/api.log

supervisorctl reload

supervisorctl status

cd ~

vim /etc/nginx/nginx.conf

listen       80;
# listen       [::]:80;

include /etc/nginx/conf.d/*.conf;
include /etc/nginx/feperf.com.conf.d/*.conf;

kill -9 $(lsof -i tcp:80 -t)

systemctl restart nginx

systemctl status nginx

curl http://127.0.0.1:3000/api/ping

setsebool -P httpd_can_network_connect 1

setsebool -P httpd_can_network_relay 1
