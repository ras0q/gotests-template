# gotests-template

My template for cweill/gotests

## Usage (VSCode)

1. Add this repo as a submodule to your project
```sh
$ git submodule add git@github.com:Ras96/gotests-template.git .vscode/gotests-template
```

2. Add gotests settings to ./vscode/settings.json
```txt
{
  ...,
  "go.generateTestsFlags": [
    "-i",
    "-parallel",
    "-template_dir",
    "./vscode/gotests-template/templates"
  ],
  ...
}
```
