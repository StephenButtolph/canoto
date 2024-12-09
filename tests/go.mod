module github.com/StephenButtolph/canoto/tests

go 1.22.9

replace github.com/StephenButtolph/canoto => ./..

require (
	github.com/StephenButtolph/canoto v0.0.0-20241209012112-9ba099fee1b3
	github.com/stretchr/testify v1.10.0
	github.com/thepudds/fzgen v0.4.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sanity-io/litter v1.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
