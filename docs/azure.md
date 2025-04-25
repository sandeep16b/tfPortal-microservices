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

| Command                     | Type of IP           | Scope                                                     | Use Case                                                     |
|------------------------------|-----------------------|------------------------------------------------------------|---------------------------------------------------------------|
| `ipconfig`                   | Internal (Private)    | Only visible inside your home or office network.           | Debugging your local network, identifying your machine’s LAN IP. |
| `Invoke-RestMethod ...`       | External (Public)     | This is the IP your ISP assigns to you, visible to the internet. | Needed when accessing cloud services like Azure SQL, VPNs, or APIs that use IP whitelisting. |

