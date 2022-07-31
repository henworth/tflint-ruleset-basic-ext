# TFLint Ruleset for terraform-provider-basic-ext

TFLint ruleset extension plugin for common terraform code syntax check

## Requirements

- TFLint v0.35+
- Go v1.18

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

Note that if you install the plugin with make install, you must omit the `version` and `source` attributes in `.tflint.hcl`:

```hcl
plugin "basic-ext" {
    enabled = true
}
```

Follow the instructions to edit the generated files and open a new pull request.
