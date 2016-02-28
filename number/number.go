// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number
package number

import (
	"github.com/gopherjs/gopherjs/js"
)

// IsInteger() method determines whether the passed value is an integer.
// See also https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/isInteger
func IsInteger(n interface{}) bool {
	return js.Global.Get("Number").Call("isInteger", n).Bool()
}
