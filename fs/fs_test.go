// Copyright 2014 dong<ddliuhb@gmail.com>.
// Licensed under the MIT license.
//
// Underscore addon for Motto
package fs

import (
	"github.com/dishuiGit/motto"
	"testing"
)

func TestFsImport(t *testing.T) {
	_, v, _ := motto.Run("tests/index.js")

	s, _ := v.ToString()
	if s != "cat" {
		t.Error("core module test failed: ", s, "!=", "cat")
	}
}
