package html2json

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

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
