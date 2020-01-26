// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package util

import (
	"log"
)

// Check .
func Check(err error, msg string) {
	if err != nil {
		log.Panicf(`
#######+
Error: `+msg+`
#######-
%v
#######+
`, err)
	}
}
