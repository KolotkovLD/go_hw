package tdd_practice

import (
	"testing"
)

func TestExecutePipeline(t *testing.T) {
	ExecutePipeline(Build, Lint, Deploy)
}

func Build() {}

func Lint() {}

func Deploy() {}
