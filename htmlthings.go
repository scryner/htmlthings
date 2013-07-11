package htmlthings

import (
	"code.google.com/p/go.net/html"
	"io"
)

type HtmlThings struct {
	docRoot *html.Node
}

func NewHtmlThings(r io.Reader) (*HtmlThings, error) {
	docRoot, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return &HtmlThings{
		docRoot: docRoot,
	}, nil
}

func (things *HtmlThings) FindAttrValues(cond *Condition, attrName string) (ret []string) {
	currentNodes := []*html.Node{things.docRoot}

	for _, ele := range cond.elements {
		var nextNodes []*html.Node

		for _, n := range currentNodes {
			nodes := findMatchedElementNodes(n, ele)
			nextNodes = append(nextNodes, nodes...)
		}

		currentNodes = nextNodes
	}

	for _, n := range currentNodes {
		for _, attr := range n.Attr {
			if attr.Key == attrName {
				ret = append(ret, attr.Val)
			}
		}
	}

	return
}

func (things *HtmlThings) FindTexts(cond *Condition) (ret []string) {
	currentNodes := []*html.Node{things.docRoot}

	for _, ele := range cond.elements {
		var nextNodes []*html.Node

		for _, n := range currentNodes {
			nodes := findMatchedElementNodes(n, ele)
			nextNodes = append(nextNodes, nodes...)
		}

		currentNodes = nextNodes
	}

	for _, n := range currentNodes {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				ret = append(ret, c.Data)
				break
			}
		}
	}

	return
}

func findMatchedElementNodes(root *html.Node, ele *Element) (nodes []*html.Node) {
	var walk func(*html.Node)

	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == ele.name {
			matched := 0

			for _, attr := range n.Attr {
				if val, ok := ele.attrs[attr.Key]; ok && val == attr.Val {
					matched += 1
				}
			}

			if matched == len(ele.attrs) {
				nodes = append(nodes, n)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(root)

	return
}
