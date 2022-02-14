# gotests-template

My template for cweill/gotests

## Usage (VSCode)

Add gotests settings to ./vscode/settings.json
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
