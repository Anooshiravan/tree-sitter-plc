package tree_sitter_plc_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-plc"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_plc.Language())
	if language == nil {
		t.Errorf("Error loading Plc grammar")
	}
}
