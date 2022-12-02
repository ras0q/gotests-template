# gotests-template

My template for cweill/gotests

## Usage (VSCode)

Add gotests settings to ./vscode/settings.json

```txt
{
  ...,
  "go.generateTestsFlags": [
    "-template_dir",
    "./path/to/Ras96/gotests-template/templates" // Or templates2
  ],
  ...
}
```

## Contents

### templates (deprecated)

[example_templates_test.go](./example_templates_test.go)

### templates2 (deprecated)

[example_templates2_test.go](./example_templates2_test.go)

- TDD
- Parallel test
- Use `map[string]struct`(map key is testname)
- Use [google/go-cmp](https://github.com/google/go-cmp)

### map-gomock-gocmp-errors-setupFields

[example_map-errors-gomock-gocmp_test.go](./example_map-errors-gomock-gocmp_test.go)

- manage subtest names as map keys
- initialize receiver fields with `*gomock.Controller`
- can change to parallel test with a `-parallel` flag
- use following packages
  - standard `errors` package
  - `gomock`: <https://github.com/golang/mock>
  - `go-cmp`: <https://github.com/google/go-cmp>
