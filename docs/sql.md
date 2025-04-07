# Package for go
Run the below cmd
go get github.com/denisenkom/go-mssqldb

# Check if SQL Server port is open
```ps
Test-NetConnection -ComputerName localhost -Port 1433
```

# Allow SQL Server port 1433 through firewall
```ps
New-NetFirewallRule -DisplayName "Allow SQL Port 1433" -Direction Inbound -Protocol TCP -LocalPort 1433 -Action Allow
```

# Restart SQL Server service
```bash
Restart-Service -Name 'MSSQL$SQLEXPRESS'
```

# Start SQL Server Browser service (for named instances)
```bash
Set-Service -Name 'SQLBrowser' -StartupType Automatic
Start-Service -Name 'SQLBrowser'
```