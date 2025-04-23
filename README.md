## REFERENCES

[Unitchecker](https://pkg.go.dev/golang.org/x/tools/go/analysis/unitchecker) VS [Singlechecker](https://pkg.go.dev/golang.org/x/tools/go/analysis/singlechecker)-

    In context of Go static analysis tools, unitchecker and singlechecker are both packages within the golang.org/x/tools/go/analysis library,
    but they serve slightly different purposes. unitchecker is designed for analyzing a single compilation unit (a package or file) during a build,
    often used by build systems like go vet. singlechecker, on the other hand, is for creating standalone tools that run a single analysis.

    unitchecker:
    1. Analyzes a single compilation unit during a build process.
    2. Invoked by build systems like go vet.
    3. Supports a command-line protocol to interact with the build system.
    4. Suitable for integration with build pipelines and tools that require analysis during compilation.

    singlechecker:
    1. Provides a way to create standalone tools for running a single analysis.
    2. Makes it easy for analysis providers to offer a standalone tool for their analysis.
    3. Can be used to create custom linters or analysis tools.
    4. Ideal for creating tools that can be run independently of a build system.

[Abstract Syntax Tree(AST)](https://pkg.go.dev/go/ast) - representation of Golang components such as modules, packages, literals, structures, declarations, comments etc. [example for basic AST traversal](https://www.zupzup.org/go-ast-traversal/index.html)

[Inspect](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/inspect) - Defines an analyser for AST inspector.

[Pass](https://pkg.go.dev/golang.org/x/tools/go/analysis#hdr-Pass) - A Pass describes a single unit of work: the application of a particular Analyzer to a particular package of Go code. The Pass provides information to the Analyzer’s Run function about the package being analyzed. It provides operations to the Run function for reporting diagnostics and other information back to the driver. The Fset, Files, Pkg, and TypesInfo fields provide source positions, the syntax trees, type information for a single package of Go code.

``` bash
type Pass struct {
	Fset   		*token.FileSet
	Files		[]*ast.File
	OtherFiles	[]string
	Pkg		*types.Package
	TypesInfo	*types.Info
	ResultOf	map[*Analyzer]interface{}
	Report		func(Diagnostic)
	...
}
```

[Guide](https://disaev.me/p/writing-useful-go-analysis-linter/) - Golang linter developer guide