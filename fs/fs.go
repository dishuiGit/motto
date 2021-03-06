package fs

import (
	"github.com/dishuiGit/motto"
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func fsModuleLoader(vm *motto.Motto) (otto.Value, error) {
	fs, _ := vm.Object(`({})`)
	fs.Set("readFile", func(call otto.FunctionCall) otto.Value {
		filename, _ := call.Argument(0).ToString()
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return otto.UndefinedValue()
		}

		v, _ := call.Otto.ToValue(string(bytes))
		return v
	})

	fs.Set("writeFile", func(call otto.FunctionCall) otto.Value {
		filePath, _ := call.Argument(0).ToString()
		str, _ := call.Argument(1).ToString()
		err:=ioutil.WriteFile(filePath, []byte(str),0644)
		if err != nil {
			return otto.UndefinedValue()
		}

		return otto.TrueValue()
	})

	return vm.ToValue(fs)
}

func init() {
	motto.AddModule("fs", fsModuleLoader)
}
