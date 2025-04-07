## To install jenkins EC2 instance from git bash
```bash
sudo cd root
sudo yum update -y
sudo amazon-linux-extras install java-openjdk11
```
```bash
## Install Java
sudo amazon-linux-extras install epel
# Download the Amazon Corretto 17 RPM package
sudo wget https://corretto.aws/downloads/latest/amazon-corretto-17-x64-linux-jdk.rpm
# Install the downloaded RPM package
sudo rpm -ivh amazon-corretto-17-x64-linux-jdk.rpm
sudo alternatives --config java
```
```bash
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat/jenkins.repo
sudo rpm --import https://pkg.jenkins.io/redhat/jenkins.io.key
sudo yum install jenkins -y

sudo wget -O /etc/yum.repos.d/jenkins.repo http://pkg.jenkins-ci.org/redhat/jenkins.repo
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo
sudo yum install jenkins -y
```
## To start jenkins
```bash
sudo systemctl start jenkins
sudo systemctl status jenkins
sudo systemctl enable jenkins
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
new-admin-password
```

## Install Git
```bash
sudo yum install git
```

## Install Golang
```bash
sudo yum remove golang
wget https://go.dev/dl/go1.24.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz
nano ~/.bash_profile
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
source ~/.bash_profile
go version
```

## Check jenkins logs
sudo journalctl -u jenkins.service --since "10 minutes ago"

## Check more detailed jenkins logs
sudo journalctl -xe | grep jenkins
