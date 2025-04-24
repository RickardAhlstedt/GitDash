# ⚙️ GitDash

## 🔨 Building 
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

## 🚀 Usage
Run `gitdash` and it will output a status for your configured path (from `~/.gitdash.yaml`)  
If you wish to use a different config-file, please use the `--config`-flag and provide a path to the file you wish to use

## 📝 Configuration
```yaml
paths:
  - ~/code
ignore:
  - "**/vendor"
  - "**/node_modules"
theme:
  name: lolcat
```
- You can setup mulitple paths that the tool will check
- Specify paths for ignore, for traversing projects

## 👓 Output
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

## 🎨 Themes
Current themes that are built in are:
- lolcat
- nord
- monochrome
- dracula
- solarized-dark
- solarized-light

ℹ️ **Please note** that if an invalid theme or no theme/colors are defined, the TUI will default to white.

### Theming
You can supply your own theme by defining it in `~/.gitdash.yaml`
```yaml
theme:
  colors:
    branch: "#ffaa00"
    ahead: "#00ff88"
    behind: "#ff5588"
    dirty: "#ff3333"
    clean: "#00dd00"
    path: "#888888"
    header: "#44ccff"
```

## CLI Flags for GitDash

You can use the following flags when running `gitdash` to customize its behavior:

- `--config <path>`: Specify the path to the GitDash configuration file. Default is `~/.gitdash.yaml`.
- `--sort <name|branch|ahead|behind>`: Sort repositories by `name`, `branch`, `ahead`, or `behind`.
- `--tui`: Use Text User Interface (TUI) to display repositories in an interactive view. Default is `false`.
- `--theme <theme_name>`: Set a theme for the interface (e.g., `lolcat`, `dracula`, `nord`, or custom theme from config).
- `--fetch <true|false>`: Fetch the latest status from the origin for each repository. Default is `false`.
- `--quiet`: Suppress log messages and display only essential information. Default is `false`.