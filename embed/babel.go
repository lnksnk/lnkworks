package embed

import (
	"io"
	"strings"
)

`

func BabelJS() io.Reader {
	return strings.NewReader(strings.Replace(babeljs, "|'|", "`", -1))
}