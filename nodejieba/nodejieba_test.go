// Copyright 2014 dong<ddliuhb@gmail.com>.
// Licensed under the MIT license.
//
// Underscore addon for Motto
package nodejieba

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
	"testing"
)

func TestFsImport(t *testing.T) {
	var s string
	var words []string
	x := gojieba.NewJieba()
	defer x.Free()

	//s = "我来到北京清华大学"
	//words = x.CutAll(s)
	//fmt.Println(s)
	//fmt.Println("全模式:", strings.Join(words, "/"))

	s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	words = x.Cut(s, true)
	fmt.Println(words)
	fmt.Println("搜索引擎模式:", strings.Join(words, "@"))
}
