## How to merge main to feature branch

Lets say you have a local branch feature123 which was pulled
More changes have been added to Remote branch main
In order for the feature123 branch to be merged to main,
you need resolve all the conflicts if you find any.
Below are the cmds to perform merge on git.

open cmd 
navigate to project folder (cd FolderPath)
```bash
git status
git checkout main
git fetch --all
git pull
```
```bash
git checkout featureRamana
git fetch --all
git pull
git merge main
```
## Verify and change the user name and email for all git commits and pushes
```bash
git config --list --show-origin
git config --global user.name "Sandeep Bijili" 
git config --global user.email "sandeep.bijili@example.com" 
```