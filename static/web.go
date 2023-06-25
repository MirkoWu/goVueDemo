package static

import "embed"

// 产物的所有路径
//
//go:embed index.html vite.svg assets
var Static embed.FS
