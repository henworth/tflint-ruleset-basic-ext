package project

import "fmt"

// Version is ruleset version
const Version string = "0.0.5"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/terraform-linters/tflint-ruleset-basic-ext/blob/%s/docs/rules/%s.md", Version, name)
}
