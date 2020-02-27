package nodejieba

import (
	"github.com/ddliu/motto"
	"github.com/robertkrimen/otto"
	"github.com/yanyiwu/gojieba"
)

func cutModuleLoader(vm *motto.Motto) (otto.Value, error) {
	fs, _ := vm.Object(`({})`)
	x := gojieba.NewJieba()
	//defer x.Free()
	fs.Set("cut", func(call otto.FunctionCall) otto.Value {
		str, _ := call.Argument(0).ToString()
		var words []string
		//s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
		words = x.Cut(str, true)
		//res := strings.Join(words, "@@")
		//fmt.Println(str)
		//fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

		v, _ := call.Otto.ToValue(words)
		return v
	})
	//x.Free()
	return vm.ToValue(fs)
}

func init() {
	motto.AddModule("nodejieba", cutModuleLoader)
}
