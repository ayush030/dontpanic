package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ayush030/dontpanic/analyzer"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPanic(t *testing.T) {
	// gives you pwd
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	var expectedPanicCounts = 5 // in testData/src/test.go

	testData := filepath.Join(filepath.Dir(filepath.Dir(wd)), "dontpanic/analyzer/testData/src")

	results := analysistest.Run(t, testData, analyzer.GetAnalyser())
	assert.EqualValues(t, expectedPanicCounts, len((*results[0]).Diagnostics))
}
