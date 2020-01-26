// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package util

import (
	"strings"
)

// B64f swaps the context between rfc3548.3 and rfc3548.4
// https://tools.ietf.org/html/rfc3548#page-6
func B64f(s string) string {
	if strings.Contains(s, "/") || strings.Contains(s, "+") {
		return strings.NewReplacer(
			"/", "_",
			"+", "-",
		).Replace(s)
	} else if strings.Contains(s, "_") || strings.Contains(s, "-") {
		return strings.NewReplacer(
			"_", "/",
			"-", "+",
		).Replace(s)
	}
	return s
}

// B64EnforceOriginal enforces original encoding
// https://tools.ietf.org/html/rfc3548#page-6
func B64EnforceOriginal(s string) string {
	strings.NewReplacer(
			"_", "/",
			"-", "+",
		).Replace(s)
	return s
}

// B64FileSafe enforces file safe encoding
// https://tools.ietf.org/html/rfc3548#page-6
func B64FileSafe(s string) string {
	strings.NewReplacer(
			"/", "_",
			"+", "-",
		).Replace(s)
	return s
}
