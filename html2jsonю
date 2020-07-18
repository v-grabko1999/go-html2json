package html2json

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Dom struct {
	Node *html.Node
}

type Node struct {
	Name       string            `json:"name,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Class      string            `json:"class,omitempty"`
	ID         string            `json:"id,omitempty"`
	Href       string            `json:"href,omitempty"`
	Text       string            `json:"text,omitempty"`
	Elements   []Node            `json:"elements,omitempty"`
}
type Matcher func(node *html.Node) bool

func New(r *strings.Reader) (*Dom, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return &Dom{Node: doc}, nil
}

func (d *Dom) ToJSON() ([]byte, error) {
	var subset []*html.Node
	subset = append(subset, d.Node)
	return convertToJSON(subset)
}

func (d *Dom) ByID(id string) (*Dom, error) {
	elem, ok := find(d.Node, byID(id))
	if !ok {
		return nil, errors.New("Unable to find an element with the ID " + id)
	}
	return &Dom{Node: elem}, nil
}

func (d *Dom) ByClass(class string) ([]*Dom, error) {
	subset := findAll(d.Node, byClass(class))
	if len(subset) == 0 {
		return nil, errors.New("Unable to find an element with the class " + class)
	}

	return convertTypeHTML(subset), nil
}

func (d *Dom) ByTag(tag atom.Atom) ([]*Dom, error) {
	subset := findAll(d.Node, byTag(tag))
	if len(subset) == 0 {
		return nil, errors.New("Unable to find an element with the tag " + tag.String())
	}

	return convertTypeHTML(subset), nil
}

func (d *Dom) ByAttribute(atr string) ([]*Dom, error) {
	keyVal := strings.SplitN(atr, "=", 2)
	key := keyVal[0]
	var value string
	if len(keyVal) == 1 {
		value = ""
	} else {
		value = keyVal[1]
	}
	subset := findAll(d.Node, matchByAttribute(key, value))
	if len(subset) == 0 {
		return nil, errors.New("Unable to find an element with attribute matcher " + atr)
	}
	return convertTypeHTML(subset), nil
}

func byTag(a atom.Atom) Matcher {
	return func(node *html.Node) bool { return node.DataAtom == a }
}

func byID(id string) Matcher {
	return func(node *html.Node) bool { return attr(node, "id") == id }
}

func byClass(class string) Matcher {
	return func(node *html.Node) bool {
		classes := strings.Fields(attr(node, "class"))
		for _, c := range classes {
			if c == class {
				return true
			}
		}
		return false
	}
}

func attr(node *html.Node, key string) string {
	for _, a := range node.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func find(node *html.Node, matcher Matcher) (n *html.Node, ok bool) {
	if matcher(node) {
		return node, true
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		n, ok := find(c, matcher)
		if ok {
			return n, true
		}
	}
	return nil, false
}

func findAll(node *html.Node, matcher Matcher) []*html.Node {
	return findAllInternal(node, matcher, false)
}

func findAllInternal(node *html.Node, matcher Matcher, searchNested bool) []*html.Node {
	matched := []*html.Node{}

	if matcher(node) {
		matched = append(matched, node)

		if !searchNested {
			return matched
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		found := findAllInternal(c, matcher, searchNested)
		if len(found) > 0 {
			matched = append(matched, found...)
		}
	}
	return matched
}

func convertTypeHTML(s []*html.Node) []*Dom {
	q := make([]*Dom, 0)
	for _, val := range s {
		q = append(q, &Dom{Node: val})
	}
	return q
}

func convertToJSON(nodes []*html.Node) ([]byte, error) {
	rootJSONnodes := make([]Node, len(nodes))
	for i, n := range nodes {
		rootJSONnodes[i].populateFrom(n)
	}

	json, err := json.MarshalIndent(rootJSONnodes, "", "    ")
	if err != nil {
		return nil, err
	}

	return json, err
}

func (n *Node) populateFrom(htmlNode *html.Node) (*Node, error) {
	switch htmlNode.Type {
	case html.ElementNode:
		n.Name = htmlNode.Data
		break

	case html.DocumentNode:
		break

	default:
		return nil, errors.New("Given node needs to be an element or document")
	}

	var textBuffer bytes.Buffer

	if len(htmlNode.Attr) > 0 {
		n.Attributes = make(map[string]string)
		var a html.Attribute
		for _, a = range htmlNode.Attr {
			switch a.Key {
			case "class":
				n.Class = a.Val

			case "id":
				n.ID = a.Val

			case "href":
				n.Href = a.Val

			default:
				n.Attributes[a.Key] = a.Val
			}
		}
	}

	e := htmlNode.FirstChild
	for e != nil {
		switch e.Type {
		case html.TextNode:
			trimmed := strings.TrimSpace(e.Data)
			if len(trimmed) > 0 {
				if textBuffer.Len() > 0 {
					textBuffer.WriteString(" ")
				}
				textBuffer.WriteString(trimmed)
			}

		case html.ElementNode:
			if n.Elements == nil {
				n.Elements = make([]Node, 0)
			}
			var jsonElemNode Node
			jsonElemNode.populateFrom(e)
			n.Elements = append(n.Elements, jsonElemNode)
		}

		e = e.NextSibling
	}

	if textBuffer.Len() > 0 {
		n.Text = textBuffer.String()
	}

	return n, nil
}

func matchByAttribute(key, value string) Matcher {
	return func(node *html.Node) bool {
		if node.Type == html.ElementNode {
			result := attr(node, key)
			if result != "" && (value == "" || value == result) {
				return true
			}
		}
		return false
	}
}
