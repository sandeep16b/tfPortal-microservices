## To install jenkins EC2 instance from git bash
```bash
sudo yum update -y
sudo amazon-linux-extras enable java-openjdk11
sudo yum install java-11-openjdk-devel -y
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat/jenkins.repo
sudo rpm --import https://pkg.jenkins.io/redhat/jenkins.io.key

sudo yum install jenkins -y
```