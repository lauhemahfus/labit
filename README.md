# Labit
Labit is a simple version control system that allows you to track changes, stage files, commit updates, and view history. 

## Current Features

- **Initialize a repository**: `labit init`  
- **Stage files**: `labit add <file>`  
- **Commit changes**: `labit commit -m "message"`  
- **View commit history**: `labit log`  
- **Check repository status**: `labit status`  

## Future Plans

- **Show diffs between versions**: `labit diff <commit1> <commit2>`  
- **Checkout / Restore to previous versions**: `labit checkout <commit>`  
- **Branching and merging**:  
  - Create branches: `labit branch <branch-name>`  
  - Merge branches: `labit merge <branch-name>`  
- **Tag important commits**: `labit tag <tag-name>`  
- **Undo commits or changes** :`labit reset`, `labit revert`  
- **Interactive history visualization**  

---

## Clone This Repository

You can clone the Labit repository using Git:

```bash
# Using HTTPS
git clone https://github.com/lauhemahfus/labit.git

# Using SSH 
git clone git@github.com:lauhemahfus/labit.git

```
Then move into the project folder:

```bash
cd labit
```



## Usages
### Linux / macOS

```bash
# Build Labit
go build -o bin/labit cmd/labit/main.go

# Copy binary to your project folder 
cp bin/labit /your/project/folder/
cd /your/project/folder/

# Make sure it's executable
chmod +x labit

# Initialize a repository
./labit init

# Stage files
./labit add file.txt

# Commit changes
./labit commit -m "Initial commit"

# View commit history
./labit log

# View repository status
./labit status
```

### Windows 

```bash
# Build Labit
go build -o bin\labit.exe cmd\labit\main.go

# Copy binary to your project folder
copy bin\labit.exe \your\project\folder\
cd \your\project\folder\

# Initialize a repository
labit.exe init

# Stage files
labit.exe add file.txt

# Commit changes
labit.exe commit -m "Initial commit"

# View commit history
labit.exe log

# View repository status
labit.exe status
```