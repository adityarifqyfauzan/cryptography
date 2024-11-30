package docs

import (
	"embed"
)

//go:embed *.md
var DocsMarkdown embed.FS
