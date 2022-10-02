// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package govulncheck

import "golang.org/x/tools/go/packages"

const (
	// analysisBinary is used for binary analysis with vulncheck.Binary.
	analysisBinary = "binary"

	// analysisSource is used for source code analysis with vulncheck.Source.
	analysisSource = "source"
)

const (
	// outputText is the default output type for `govulncheck`.
	outputText = "text"

	//  outputVerbose is the output type for `govulncheck -v`.
	outputVerbose = "verbose"

	// outputJSON is the output type for `govulncheck -json`, which will print
	// the JSON-encoded vulncheck.Result.
	outputJSON = "json"

	// outputSummary is the output type for `govulncheck -summary-json`, which
	// will print the JSON-encoded govulncheck.Summary.
	//
	// This is only meant by use for experimental with gopls.
	outputSummary = "summary"
)

const (
	envGOVULNDB = "GOVULNDB"
	vulndbHost  = "https://vuln.go.dev"
)

// Config is the configuration for Main.
type Config struct {
	// AnalysisType specifies the vulncheck analysis type.
	AnalysisType string

	// OutputType specifies the output format type.
	OutputType string

	// Patterns are either the binary path for "binary" analysis mode, or
	// go package patterns for "source" analysis mode.
	Patterns []string

	// SourceLoadConfig specifies the package loading configuration.
	SourceLoadConfig packages.Config
}
