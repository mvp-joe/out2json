package out2json

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	root, err := Parse(`
root
- child 1
    - grandchild 1
    - grandchild 2
- child 2
    - grandchild 1
`)
	require.NoError(t, err)

	// validate root
	require.NotNil(t, root, "root should exist")
	assert.Equal(t, "root", root.Text)
	assert.Equal(t, 0, root.Depth)
	require.Len(t, root.Children, 2)

	// validate Child 1
	child1 := root.Children[0]
	assert.Equal(t, "child 1", child1.Text)
	assert.Equal(t, 1, child1.Depth)
	require.Len(t, child1.Children, 2)
	c1Grandchild1 := child1.Children[0]
	assert.Equal(t, "grandchild 1", c1Grandchild1.Text)
	assert.Equal(t, 2, c1Grandchild1.Depth)
	c2Grandchild2 := child1.Children[1]
	assert.Equal(t, "grandchild 2", c2Grandchild2.Text)
	assert.Equal(t, 2, c2Grandchild2.Depth)

	// validate Child 2
	child2 := root.Children[1]
	assert.Equal(t, "child 2", child2.Text)
	assert.Equal(t, 1, child2.Depth)
	c2Grandchild1 := child1.Children[0]
	assert.Equal(t, "grandchild 1", c2Grandchild1.Text)
	assert.Equal(t, 2, c2Grandchild1.Depth)
}

func Test_normalizeLines(t *testing.T) {
	lines := strings.Split(`Root
- Child 1
    - Grandchild 1
    - Grandchild 2
        - Great-Grandchild 1
- Child 2
    - Grandchild 1`, "\n")

	normalized, err := normalizeLines(lines)
	require.NoError(t, err)
	require.Len(t, normalized, 7, "should parse seven lines")

	l1 := normalized[0]
	assert.Equal(t, "Root", l1.Text)
	assert.Equal(t, 0, l1.Depth)

	l2 := normalized[1]
	assert.Equal(t, "Child 1", l2.Text)
	assert.Equal(t, 1, l2.Depth)

	l3 := normalized[2]
	assert.Equal(t, "Grandchild 1", l3.Text)
	assert.Equal(t, 2, l3.Depth)

	l4 := normalized[3]
	assert.Equal(t, "Grandchild 2", l4.Text)
	assert.Equal(t, 2, l4.Depth)

	l5 := normalized[4]
	assert.Equal(t, "Great-Grandchild 1", l5.Text)
	assert.Equal(t, 3, l5.Depth)

	l6 := normalized[5]
	assert.Equal(t, "Child 2", l6.Text)
	assert.Equal(t, 1, l6.Depth)

	l7 := normalized[6]
	assert.Equal(t, "Grandchild 1", l7.Text)
	assert.Equal(t, 2, l7.Depth)
}
