package out2json

type Line struct {
	Text  string `json:"text"`
	Depth int    `json:"depth"`
}

type Node struct {
	Text     string  `json:"text"`
	Children []*Node `json:"children"`
	Depth    int     `json:"depth"`
}

func NewNode(line Line) *Node {
	return &Node{
		Text:  line.Text,
		Depth: line.Depth,
	}
}
