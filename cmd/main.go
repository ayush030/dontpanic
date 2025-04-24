package main

import (
	"github.com/ayush030/dontpanic/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.GetAnalyser())
}
