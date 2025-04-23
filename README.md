GitDash 🖥️

# 🔨 Building 
1. Clone the repo
`git clone git@github.com:RickardAhlstedt/GitDash.git && cd GitDash`
2. If you have my [cicd-go](https://github.com/RickardAhlstedt/cicd-go) installed prior simply execute `cicd-go`
3. If you don't have it, simply execute the following commands:
```bash
go mod tidy
go build -o gitdash ./cmd
cp gitdash.yaml ~/.gitdash.yaml
# optional
cp gitdash /usr/local/bin/gitdash
```

# 🚀 Usage
Run `gitdash` and it will output a status for your configured path (from `~/.gitdash.yaml`)  
If you wish to use a different config-file, please use the `--config`-flag and provide a path to the file you wish to use

# 📝 Configuration
```yaml
paths:
  - ~/code
ignore:
  - "**/vendor"
  - "**/node_modules"
theme:
  accent_color: cyan #Not yet implemented
```
- You can setup mulitple paths that the tool will check
- Specify paths for ignore, for traversing projects
- In the future you will be able to specify a theme for the output, planned to be used when tview is implemented

## Output
```
╰─ gitdash
📂 Scanning paths for git repositories...
🔍 Found 2 repositories

📁 GitDash                        [main] ↑0 ↓0 ✴
📁 cicd-go                        [main] ↑0 ↓0 ✴
```
### Reading the output

```
[main] ↑0 ↓0 ✴
|      |  |   \ Marks if the repo is dirty (uncommited changes)
|      |   \ How many commits the repo is behind with   
|       \ How many commits the repo is ahead with
 \ The current branch of the repo
```
