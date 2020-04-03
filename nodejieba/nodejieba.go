package nodejieba

import (
	"github.com/ddliu/motto"
	"github.com/robertkrimen/otto"
	"github.com/huichen/sego"
)

func cutModuleLoader(vm *motto.Motto) (otto.Value, error) {
	fs, _ := vm.Object(`({})`)
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("dictionary.txt")
	fs.Set("cut", func(call otto.FunctionCall) otto.Value {
		str, _ := call.Argument(0).ToString()
		var words []string
		//s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
		segments := segmenter.Segment(str)
		words = sego.SegmentsToSlice(segments, true)

		v, _ := call.Otto.ToValue(words)
		return v
	})
	//x.Free()
	return vm.ToValue(fs)
}

func init() {
	motto.AddModule("nodejieba", cutModuleLoader)
}
