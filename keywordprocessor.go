package flashtext

import (
	"unicode"
	"unicode/utf8"
)

/*
* @Author: zouyx
* @Email:
* @Date:   2023/5/17 17:27
* @Package:
 */

type WalkFn func(start, end int) bool

type KeywordProcessor struct {
	root              *Node
	whiteSpaceChars   []string // 不知道有啥用，反正python也定义了，就跟着定义一个吧
	nonWordBoundaries string
	caseSensitive     bool // 匹配是否区分大小写
}

func NewKeywordProcessor(caseSensitive bool) *KeywordProcessor {
	return &KeywordProcessor{
		root:            newNode(),
		caseSensitive:   caseSensitive,
		whiteSpaceChars: []string{".", "\t", "\n", "\a", " ", ","}, // python 就是这样定义的，而且它也没有用到，所以先留着这里吧
	}
}

func (kp *KeywordProcessor) setItem(keyword string) {
	if len(keyword) == 0 {
		return
	}

	node := kp.root
	for _, char := range keyword {
		if !kp.caseSensitive {
			char = unicode.ToLower(char)
		}
		if _, ok := node.children[char]; !ok {
			// 子节点已存在 将子节点设为当前节点
			node.children[char] = newNode()
		}
		node = node.children[char]
	}
	// 记录当前匹配词的长度
	node.exist[utf8.RuneCountInString(keyword)] = struct{}{}
}

func (kp *KeywordProcessor) Build() {
	// 创建一个空队列，用于层级遍历 Trie 树节点
	queue := make([]*Node, 0)
	// 将根节点作为初始节点
	queue = append(queue, kp.root)

	for len(queue) > 0 {
		// 获取队列中的第一个节点
		currentNode := queue[0]
		// 弹出队列中的第一个节点
		queue = queue[1:]

		// 遍历当前节点的子节点
		for char, childNode := range currentNode.children {
			// 将子节点添加到队列中
			queue = append(queue, childNode)
			// 当前节点父节点的失败指针
			faFail := currentNode.failure

			// 递归获取父失败指针 直到获取到根节点为止
			for faFail != nil && faFail.children[char] == nil {
				faFail = faFail.failure
			}
			childNode.failure = kp.root
			if faFail != nil {
				childNode.failure = faFail.children[char]
			}
			for key := range childNode.failure.exist {
				childNode.exist[key] = struct{}{}
			}
		}
	}
}
func (kp *KeywordProcessor) AddKeyWord(keyword string) *KeywordProcessor {
	kp.setItem(keyword)
	return kp
}

func (kp *KeywordProcessor) AddKeywordsFromList(keywords []string) *KeywordProcessor {
	for _, keyword := range keywords {
		kp.setItem(keyword)
	}
	return kp
}

func (kp *KeywordProcessor) walk(sentence string, wf WalkFn) {
	// 从根节点开始查找
	currentNode := kp.root
	// 遍历文本 sentence 的每个字符，并记录当前字符的索引为 idx，当前字符为 r
	for idx, r := range sentence {
		if !kp.caseSensitive {
			r = unicode.ToLower(r)
		}
		// 在循环中 判断是否有当前字符子节点
		//如果不存在，则说明匹配失败，需要通过失败路径回溯到前一个节点，直到找到一个匹配的子节点或回溯到根节点。
		for currentNode.children[r] == nil && currentNode.failure != nil {
			currentNode = currentNode.failure
		}
		if currentNode.children[r] == nil {
			continue

		}
		currentNode = currentNode.children[r]
		for length := range currentNode.exist {
			if !wf(idx-length+1, idx) {
				return
			}
		}

	}
}

// ExtractKeywords 匹配关键词
func (kp *KeywordProcessor) ExtractKeywords(sentence string) []Match {
	var matches []Match
	if len(sentence) == 0 {
		return matches
	}

	kp.walk(sentence, func(start, end int) bool {
		matches = append(matches, Match{
			start: start,
			end:   end,
			match: sentence[start : end+1],
		})
		return true
	})
	return matches
}
