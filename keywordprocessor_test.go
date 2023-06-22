package flashtext

import (
	"fmt"
	"github.com/ayoyu/flashtext"
	"testing"
)

var (
	keys = []string{
		"he", "she", "hers", "his", "share",
	}
	key = "ahishershare"
)

func TestKeywordProcessor_AddKeywordsFromList(t *testing.T) {
	trie := NewKeywordProcessor(false)

	trie.AddKeywordsFromList(keys).Build()
	matches := trie.ExtractKeywords(key)
	for _, match := range matches {
		fmt.Println(match.MatchString())
		fmt.Println("Start:", match.Start())
		fmt.Println("End:", match.End())
		fmt.Println("---")
	}
}

func TestKeywordProcessor_Flash(t *testing.T) {
	var flash = flashtext.NewFlashKeywords(false)

	for _, word := range keys {
		flash.Add(word)
	}
	matches := flash.Search(key)
	for _, match := range matches {
		fmt.Println("match:", match.Key)
		fmt.Println("Start:", match.Start)
		fmt.Println("End:", match.End)
		fmt.Println("---")
	}
}
