### Silently remove or delete files and folders for the computer (Windows)
```bash
Remove-Item "$env:TEMP\*" -Recurse -Force -ErrorAction SilentlyContinue
```