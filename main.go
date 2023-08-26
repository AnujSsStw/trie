package main

import (
	"flag"
	"fmt"
	"strings"
)

type Node struct {
	val      string
	end      bool
	children [26]*Node
}

type Trie struct {
	RootNode *Node
}

func (t *Trie) insertText(text string) {
	zero := []rune("a")[0]
	text = strings.ToLower(text)
	curr := t.RootNode

	for _, e := range text {
		idx := e - zero
		if curr.children[idx] == nil {
			curr.children[idx] = &Node{val: string(rune(e))}
		}
		curr = curr.children[idx]
		curr.end = false
	}
	curr.end = true
}

func Autocomplete(node *Node, prefix string, sugg *[]string) {
	if node == nil {
		return
	}

	if node.end {
		*sugg = append(*sugg, prefix)
		// fmt.Println(prefix)
	}

	for i, child := range node.children {
		if child != nil {
			su := string(rune('a' + i))
			Autocomplete(child, prefix+su, sugg)
		}
	}

}
func printAutoSuggestions(root *Node, text string) []string {
	pCrawl := root
	text = strings.ToLower(text)
	zero := []rune("a")[0]
	sugg := []string{}

	for _, e := range text {
		idx := e - zero
		if pCrawl.children[idx] != nil {
			pCrawl = pCrawl.children[idx]
		} else {
			return sugg
		}
	}
	Autocomplete(pCrawl, text, &sugg)
	return sugg
}

func main() {
	// dist := flag.String("dist", "/mnt/c/Users/anujZ/Downloads/trie.svg", "svg location for output")
	autoCum := flag.String("text", "", "word for autoComplete")

	flag.Parse()
	a := Trie{RootNode: &Node{}}
	for i := 0; i < len(Char); i++ {
		a.insertText(Char[i])
	}

	// fmt.Println("digraph Trie {")
	// PrintTrieDOT(a.RootNode, "", 0)
	// fmt.Println("}")

	if len(*autoCum) > 0 {
		fmt.Println("Autocomplete suggestions:")
		ans := printAutoSuggestions(a.RootNode, *autoCum)
		fmt.Printf("%+q", ans)
	}

}

func PrintTrieDOT(node *Node, parent string, parIdx int) {
	for i, child := range node.children {
		if child != nil {
			var color string
			if child.end {
				color = "green"
			} else {
				color = "pink"
			}
			fmt.Printf(`	"%s" -> "%s" [color=%s];`, parent, child.val, color)
			fmt.Println()
			PrintTrieDOT(child, child.val, i+1)
		}
	}
}
