# flashtext

flashtext 是一个使用 Golang 实现的 AC 自动机库，用于在文本中查找和提取关键词。它使用高效的算法和数据结构，可以快速匹配大量的关键词，并返回匹配结果。

## 特点

- 高效的关键词匹配：flashtext 使用 AC 自动机算法，在文本中快速查找和提取关键词，具有良好的性能和扩展性。

- 多模式匹配：支持同时匹配多个关键词，返回所有匹配结果。

- 大小写敏感性：根据需要，可以选择是否区分关键词的大小写。

## 安装

使用以下命令将 flashtext 库添加到你的项目中：

```bash
go get github.com/code-innovator-zyx/flashtext
```

## 使用示例

下面是一个简单的示例，展示了如何使用 flashtext 进行关键词提取：

```go
import (
	"github.com/code-innovator-zyx/flashtext"
	"fmt"
)

func main() {
	// 创建一个关键词处理器，不区分大小写
	kp := flashtext.NewKeywordProcessor(false)

	// 添加关键词
	kp.AddKeyword("apple")
	kp.AddKeyword("banana")
	kp.AddKeyword("orange")

	// 提取关键词
	sentence := "I have an Apple and a Banana."
	matches := kp.ExtractKeywords(sentence)

	// 打印匹配结果
	for _, match := range matches {
		fmt.Println("Match:", match.Match)
		fmt.Println("Start:", match.Start)
		fmt.Println("End:", match.End)
		fmt.Println("---")
	}
}
```

输出：

```
Match: Apple
Start: 9
End: 13
---
Match: Banana
Start: 20
End: 25
---
```

## 与其他库的对比

- flashtext vs python flashtext


python 库用过的都知道，匹配bug何其多
### python demo1  
```python
from flashtext import KeywordProcessor

kp = KeywordProcessor()
words = ["he", "she", "hers", "his", "share"] 
for word in words:
    kp.add_keyword(keyword=word)  # 新增关键词
kp.extract_keywords('ahishehersshare')
```

```
匹配结果:惊喜吧，啥也配匹配上
  []
```
### python demo2
```python
from flashtext import KeywordProcessor

kp = KeywordProcessor()
words = ["he", "she", "hers", "his", "share"] 
for word in words:
    kp.add_keyword(keyword=word)  # 新增关键词
kp.extract_keywords('a his he hers share')  # 这里手动分词一下
```

```
匹配结果: 
  ['his', 'he', 'hers', 'share']
```
#### python flashtext 支持数据清洗  本库暂不支持 见后面说明
- flashtext vs others flashtext
### flashtext
```go

import (
	"github.com/code-innovator-zyx/flashtext"
	"fmt"
)

var (
	keys = []string{
		"he", "she", "hers", "his", "share",
	}
	key = "ahishershare"
)

func main() {
	// 创建一个关键词处理器，不区分大小写
	kp := flashtext.NewKeywordProcessor(false)

	// 添加关键词
	kp.AddKeywordsFromList(keys).Build()

	// 提取关键词
	matches := kp.ExtractKeywords(key)

	// 打印匹配结果
	for _, match := range matches {
		fmt.Println("Match:", match.Match)
		fmt.Println("Start:", match.Start)
		fmt.Println("End:", match.End)
		fmt.Println("---")
	}
}
```
输出
```
his
Start: 1
End: 3
---
she
Start: 3
End: 5
---
he
Start: 4
End: 5
---
hers
Start: 4
End: 7
---
share
Start: 7
End: 11
---
```
### [ayoyu/flashtext](github.com/ayoyu/flashtext)
```go

import (
	"github.com/ayoyu/flashtext"
	"fmt"
)

var (
	keys = []string{
		"he", "she", "hers", "his", "share",
	}
	key = "ahishershare"
)

func main() {
	// 创建一个关键词处理器，不区分大小写
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
```
结果
```
match: his
Start: 1
End: 3
---
match: he
Start: 4
End: 5
---
match: hers
Start: 4
End: 7
```


## 为什么不支持数据清洗

我们的AC自动机库专注于高效的关键词匹配和提取功能，而没有包含数据清洗功能。以下是一些原因解释为什么我们选择不支持数据清洗：

- **重复结果的可能性**：在多模式匹配下，数据清洗功能可能导致结果的重复。考虑到AC自动机的工作原理，一旦匹配到关键词，它会记录匹配的位置。如果在数据清洗过程中删除了匹配结果所在的文本片段，那么可能会导致后续的匹配结果无法被识别，从而引入了重复的结果。

- **专注于核心任务**：AC自动机的设计目标是快速在文本中查找关键词的出现，以实现高效的模式匹配。我们致力于提供一个轻量级且高性能的库，专注于核心任务，以满足用户对快速关键词匹配的需求。

- **单一职责原则**：根据软件工程的设计原则，每个组件或库应该专注于完成单一的任务。将数据清洗功能集成到AC自动机库中会增加复杂性和维护负担，而且数据清洗通常涉及到更广泛的文本处理任务。

- **灵活的组合**：我们鼓励用户根据其具体需求构建完整的文本处理流程。通过将AC自动机与其他专门用于数据清洗的工具或库结合使用，用户可以根据需要选择最适合的工具和方法来完成数据清洗任务，实现更灵活的文本处理流程。

## 贡献

欢迎贡献代码、报告问题或提供改进建议。请在 GitHub 项目页面提交 issue 或 pull 请求。

## 许可证

flashtext 使用 [MIT 许可证](https://github.com/code-innovator-zyx/flashtext）