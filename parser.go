package out2json

import (
	"errors"
	"regexp"
	"strings"
)
import lls "github.com/daichi-m/go18ds/stacks/linkedliststack"

var startSpacesRx = regexp.MustCompile(`^( )*`)

const spacesPerTab = 4

// Parse parses the outline format, returning a root Node or nil if the tree was empty or not valid
func Parse(input string) (*Node, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	normalized, err := normalizeLines(lines)
	if err != nil {
		return nil, err
	}
	return createNodeTree(normalized)
}

// normalizeLines parses lines and normalizes line depth and line text (making sure lines are always at most one more level
// deeper than the previous line
func normalizeLines(lines []string) ([]Line, error) {
	var result []Line
	prevLineDepth := 0

	for _, line := range lines {
		lineDepth := getLineDepth(line)
		normalizedDepth := normalizeLineDepth(lineDepth, prevLineDepth)
		result = append(result, Line{
			Text:  normalizeLineText(line),
			Depth: normalizedDepth,
		})
		prevLineDepth = normalizedDepth
	}

	return result, nil
}

// normalizeLineDepth ensures that lineDepth only indents at most one level each time from parent
func normalizeLineDepth(lineDepth int, prevLineDepth int) int {
	if lineDepth > prevLineDepth {
		return prevLineDepth + 1
	} else {
		return lineDepth
	}
}

// normalizeLineText extracts the text of the line (minus the - and any extra spacing)
func normalizeLineText(line string) string {
	text := strings.TrimSpace(line)
	if strings.HasPrefix(text, "-") {
		text = strings.TrimSpace(text[1:len(text)])
	}
	return text
}

// getLineDepth calculates the depth of a given line based on the following rules:
// - If the line has no preceding tabs ('\t') and doesn't tart with '-' it is considered level 0
// - Otherwise any line starting with zero or more tabs followed by a dash has the depth of number of tabs+1
func getLineDepth(line string) int {
	runes := []rune(line)
	if len(runes) < 1 {
		return 0
	}
	match := startSpacesRx.Find([]byte(line))
	matchStr := string(match)
	spaceCount := len(matchStr)
	tabCount := spaceCount / spacesPerTab
	if tabCount == 0 && runes[0] != rune('-') {
		// no dash '-' char means it is root
		return 0
	}
	return tabCount + 1
}

// createNodeTree returns the root node in a node tree (or nil if there are no lines or an error)
func createNodeTree(lines []Line) (*Node, error) {
	// validate input
	if len(lines) == 0 {
		return nil, nil
	}
	rootLine := lines[0]
	if rootLine.Depth != 0 {
		return nil, errors.New("first line is not at level 0 (root)")
	}

	// setup up stack
	rootNode := NewNode(rootLine)
	prevNode := rootNode
	parentStack := lls.New[*Node]()

	// add lines to correct parents
	for _, line := range lines[1:] {
		node := NewNode(line)

		// push the prev node as parent
		if node.Depth > prevNode.Depth {
			parentStack.Push(prevNode)
		} else if node.Depth < prevNode.Depth {
			// pop stack until we get to the correct parent
			for !parentStack.Empty() {
				p, ok := parentStack.Pop()
				if ok && p.Depth == node.Depth {
					break
				}
			}
		}
		if parent, ok := parentStack.Peek(); ok {
			parent.Children = append(parent.Children, node)
		}
		prevNode = node
	}

	return rootNode, nil
}
