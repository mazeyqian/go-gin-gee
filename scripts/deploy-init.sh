# Linux-CentOS-9-x64

# Install
cd ~;
yum -y update;
yum install nginx -y;
yum install golang -y;
yum install epel-release -y;
yum install supervisor -y;
yum install git -y;
go version;

# Git
cd ~;
ssh-keygen -t rsa -C "mazey@mazey.net";
cat ~/.ssh/id_rsa.pub;

# Web
cd /;
mkdir web;
cd web;
git clone git@github.com:mazeyqian/go-gin-gee.git;
cd /web/go-gin-gee;
go run scripts/init/main.go;
GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go
# GOOS=linux GOARCH=amd64 go build -o dist/api cmd/api/main.go;
# go run cmd/api/main.go;

# Supervisor
cd /etc/supervisord.d;
nano api.ini;
`
[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api
autostart=true
autorestart=true
stderr_logfile=/web/go-gin-gee/log/api.err
stdout_logfile=/web/go-gin-gee/log/api.log
`
systemctl start supervisord;
systemctl status supervisord;
systemctl enable supervisord;
ps -ef|grep supervisord;
curl http://127.0.0.1:3000/api/ping;
# supervisorctl reload;
# supervisorctl status;

# Nginx
cd /etc/nginx;
kill -9 $(lsof -i tcp:80 -t);
setsebool -P httpd_can_network_connect 1;
setsebool -P httpd_can_network_relay 1;
firewall-cmd --zone=public --add-port=80/tcp --permanent;
firewall-cmd --reload;
nginx;
chkconfig nginx on;
curl http://127.0.0.1;
git clone git@github.com:mazeyqian/feperf.com.conf.d.git;
vim /etc/nginx/nginx.conf;
`
listen       80;
# listen       [::]:80;
include /etc/nginx/conf.d/*.conf;
include /etc/nginx/feperf.com.conf.d/*.conf;
`
systemctl restart nginx;
systemctl status nginx;
