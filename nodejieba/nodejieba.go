package nodejieba

import (
	"github.com/dishuiGit/motto"
	"github.com/robertkrimen/otto"
	"github.com/huichen/sego"
	"os"
    "os/exec"
	"path/filepath"
)

func cutModuleLoader(vm *motto.Motto) (otto.Value, error) {
	fs, _ := vm.Object(`({})`)
	file, _ := exec.LookPath(os.Args[0])
	dir := filepath.Dir(file)
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(dir+"/dictionary.txt")
	fs.Set("cut", func(call otto.FunctionCall) otto.Value {
		str, _ := call.Argument(0).ToString()
		var words []string
		//s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
		segments := segmenter.Segment([]byte(str))
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
