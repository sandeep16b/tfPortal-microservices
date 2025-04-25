# Azure Documentation

## Add Firewall rules
### find the public external IP address of you local machine
#### Open Powershell and run the below command
 ```bash
 (Invoke-RestMethod -Uri "https://api.ipify.org")
 ```
 #### Open Azure Portal
 #### Search for Database
 #### Go to Security -> networking -> Add your IP address

 ## ipconfig and InvokeRestMethod

| Command                     | Type of IP           | Scope          | Use Case                                      |
|----------------------------|-----------------------|----------------|-----------------------------------------------|
| `ipconfig`                 | Internal (Private)    | Local network  | Debugging LAN, router issues                  |
| `Invoke-RestMethod ...`    | External (Public)     | Internet-wide  | Access to cloud services, whitelisting in Azure |
