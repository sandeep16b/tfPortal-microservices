## Azure Devops Pipeline to AWS EC2
### cd to Downloads on your local
```bash
chmod 400 tfPortal-microservices-dev.pem
ssh -i "tfPortal-microservices-dev.pem" ubuntu@54.242.162.37

```


### NGINX 

#### Global NGINX settings
```bash
sudo nano /etc/nginx/nginx.conf
```
### Virtual host config for a specific domain/site
#### Open NGINX config in a text editor (like nano):
```bash
sudo nano /etc/nginx/sites-available/default
```
#### Then reload NGINX:
```bash
sudo nginx -t
sudo systemctl reload nginx
```
#### Test NGINX config:
```bash
sudo nginx -t
```
#### Reload NGINX to apply changes:
```bash
sudo systemctl reload nginx
```

#### 
```bash
ls -l /etc/nginx/sites-enabled
```

### ✅ Step 1: Confirm App Is Running on EC2
On your EC2 instance, run:

```bash
ps aux | grep '[m]ain.go'
```

```bash
curl http://localhost:8093/user/
```

### 🔍 Check the logs:
Run this:

```bash
cat user_app.log
```

```bash
cat user_app.err
```