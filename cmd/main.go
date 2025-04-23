package main

import (
	"dontpanic/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.GetAnalyser())
}
