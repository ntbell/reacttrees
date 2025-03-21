package reacttrees

import (
	"github.com/ntbell/reacttrees/parse"
)

var EntryPoint = "../test-react-app/src/index.js"

var contentStruct = []struct {
	importLine    string // Whole line with the import statement
	importPath    string // Path extracted from the importLine
	filename      string // Filename extracted from the importPath (for creating tree-view)
	componentName string // For detecting component and matching importLine
}{}

func main() {
	parse.FileNameWithoutExtension()
	// 1. Collect entrypoint file
	// 2. Read contents
	// 3. Search for indicator characters like < />
	// 4. Find filename from import statements
	// 5. Go to file path
	// 6. Repeat
}
