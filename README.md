# walk
walk is a CLI application that crawls into file system directories looking for specific files.

## Features
- Search for specific files
- Filter by size
- Listing the matched files
- Archiving the matched files
- Deleting the matched files

## Building from source
Ensure the GO SDK is installed
1. Clone the repository
   ```bash
   git https://github.com/hayohtee/walk.git
   ```
3. Change into the project directory
   ```bash
   cd walk
   ```
4. Compile
   ```bash
   go build ./...
   ```
## Usage
```bash
# Listing files that match a particular extension
./walk -root [directory name] -ext [file extension] -list

# Filter by size
./walk -root [directory name] -ext [file extension] -size [the maximum size in bytes] -list

# Deleting files that match a particular extension
./walk -root [directory name] -ext [file extension] -del

# Archiving files that match a particular extension
./walk -root [directory name] -ext [file extension] -archive [archive directory name]

# Writing the output of deleted files to the log
./walk -root [directory name] -ext [file extension] -del -log [file name for the log]

# Show all available commands
./walk -h
```
