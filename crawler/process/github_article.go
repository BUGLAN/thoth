package process

import (
	"fmt"
)

// GithubReadMeProcess TODO 使用context作为一个重构, 加入一些元信息
func (process *KanShu) GithubReadMeProcess(html string, text string) error {
	fmt.Println(html)
	fmt.Println(text)
	return nil
}
