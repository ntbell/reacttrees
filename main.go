package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ntbell/reacttrees/parse"
)

// https://www.w3schools.com/tags/
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element
var htmlElements = map[string]bool{
	"a":          true,
	"abbr":       true,
	"address":    true,
	"audio":      true,
	"b":          true,
	"bdi":        true,
	"bdo":        true,
	"blockquote": true,
	"body":       true,
	"br":         true,
	"button":     true,
	"canvas":     true,
	"caption":    true,
	"cite":       true,
	"code":       true,
	"col":        true,
	"colgroup":   true,
	"data":       true,
	"details":    true,
	"dfn":        true,
	"div":        true,
	"dl":         true,
	"dt":         true,
	"em":         true,
	"embed":      true,
	"fieldset":   true,
	"figcaption": true,
	"figure":     true,
	"footer":     true,
	"form":       true,
	"h1":         true,
	"h2":         true,
	"h3":         true,
	"h4":         true,
	"h5":         true,
	"h6":         true,
	"head":       true,
	"header":     true,
	"hr":         true,
	"html":       true,
	"iframe":     true,
	"img":        true,
	"input":      true,
	"ins":        true,
	"label":      true,
	"legend":     true,
	"li":         true,
	"link":       true,
	"main":       true,
	"mark":       true,
	"menu":       true,
	"menuitem":   true,
	"meta":       true,
	"meter":      true,
	"nav":        true,
	"ol":         true,
	"option":     true,
	"p":          true,
	"param":      true,
	"picture":    true,
	"pre":        true,
	"progress":   true,
	"q":          true,
	"rp":         true,
	"rt":         true,
	"ruby":       true,
	"s":          true,
	"script":     true,
	"section":    true,
	"select":     true,
	"slot":       true,
	"small":      true,
	"source":     true,
	"span":       true,
	"strong":     true,
	"style":      true,
	"sub":        true,
	"summary":    true,
	"sup":        true,
	"table":      true,
	"template":   true,
	"textarea":   true,
	"tfoot":      true,
	"th":         true,
	"thead":      true,
	"time":       true,
	"title":      true,
	"track":      true,
	"u":          true,
	"ul":         true,
	"var":        true,
	"video":      true,
	"wbr":        true,
}

// Define allowed extensions
var allowedExtensions = []string{".tsx", ".jsx", ".ts", ".js"}

type Node struct {
	Component string
	Children  []*Node
}

func NewNode(value string) *Node {
	return &Node{
		Component: value,
		Children:  []*Node{},
	}
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}

	// Print the current node's value with indentation based on its level
	fmt.Printf("%s%s\n", strings.Repeat(" ", level*2), node.Component)

	// Recursively print each child node with increased indentation
	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}

func ConstructTree(entrypoint string, node *Node) error {
	file, err := os.Open(entrypoint)
	if err != nil {
		// log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Lines with import statements
	var imports = []string{}

	// Lines containing JSX
	var jsx = []string{}

	// Read the file line-by-line
	for scanner.Scan() {
		line := scanner.Text()

		// ToDo: What if we return JSX from a function? (Unsupported for now?)
		line = strings.Trim(line, " ")

		// Skip empty lines
		if line == "" {
			continue
		}

		// Check if the line contains "import" or "require" and ends with one of the allowed extensions
		if strings.HasPrefix(line, "import") || strings.Contains(line, "require") {
			for _, ext := range allowedExtensions {
				if strings.Contains(line, ext) {
					// Append the import statement if the line has an allowed extension
					imports = append(imports, line)
					break // Exit the loop once a match is found
				}
			}
		}

		// Find lines with JSX (with strings like: "<SomeText")
		re := regexp.MustCompile(`<\w+`)
		if re.MatchString(line) {
			jsx = append(jsx, line)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get all JSX tags without "<" or "/>"
	var tags = []string{}
	for _, jsxLine := range jsx {
		// Regular expression to match tags like <TagName> or <TagName />
		// Ignores <!DOCTYPE and <!-- tags
		re := regexp.MustCompile(`<([a-zA-Z0-9\.]+)`)

		// Find matches in the input string
		// match[0] contains original string
		// match[1] contains tag name
		match := re.FindStringSubmatch(jsxLine)

		// There wasn't valid JSX in the line
		if match == nil {
			continue
		}

		tag := match[1]

		// Skip the HTML elements since they don't have a corresponding file
		if htmlElements[tag] {
			continue
		}

		tags = append(tags, tag)
	}

	// Find the corresponding import for each tag
	for _, tag := range tags {
		for _, importLine := range imports {
			// We found the import statement!
			if strings.Contains(importLine, tag) {
				// Extract the path from the import

				// Regular expression to match text between single quotes
				re := regexp.MustCompile(`['"](.*?)['"]`)

				// Find all matches in the input string
				matches := re.FindAllStringSubmatch(importLine, -1)

				// All matches found in single or double quotes
				for _, match := range matches {
					if len(match) > 1 {
						// Create a new node for the tag (component)
						// Only do that here, so that we don't draw for non-imported components
						child := NewNode(tag)
						node.AddChild(child)

						// If we found an import, navigate there and recurse
						routePrefix := "./test-react-app/src/"
						path := filepath.Join(routePrefix, match[1])
						ConstructTree(path, child)
					}
				}

			}
		}
	}

	return nil
}

// Get the React root file (ToDo: passed in CMD?)

func main() {
	// Get the root path from the arguments
	// Ex. go run . ./test-react-app/src/index.js
	rootPath := os.Args[1]

	// Create the root node
	rootFileName := parse.FileNameWithoutExtension(rootPath)
	rootNode := NewNode(rootFileName)

	// Create our N-ary tree
	err := ConstructTree(rootPath, rootNode)
	if err != nil {
		log.Fatal(err)
	}

	// Print the tree
	PrintTree(rootNode, 0)
}
