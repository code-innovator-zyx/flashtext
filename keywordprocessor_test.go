package flashtext

import (
	"fmt"
	"github.com/ayoyu/flashtext"
	"testing"
)

var (
	keys = []string{
		"he", "she", "hers", "his", "share", "毛泽东",
	}
	key = "ahishershare毛泽东dsadsadsagfasdasd12332154"
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
	//fmt.Println(key)
	fmt.Println(key[12:21])
	//fmt.Println(utf8.DecodeRuneInString("h"))
}

/**
his
Start: 1
End: 4
---
he
Start: 4
End: 6
---
she
Start: 3
End: 6
---
hers
Start: 4
End: 8
---
share
Start: 7
End: 12
---
毛泽东
Start: 12
End: 21
---
*/

func TestKeywordProcessor_AddKeywordsFromList_Bytes(t *testing.T) {
	trie := NewKeywordProcessor(false)

	trie.AddKeywordsFromList(keys).Build()
	matches := trie.ExtractKeywordsFromBytes([]byte(key))
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
