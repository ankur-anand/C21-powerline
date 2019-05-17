package main

import (
	"os"
)

const goModFile = "./go.mod"

func segmentGolang(p *powerline) {
	stat, err := os.Stat(goModFile)
	if err == nil && !stat.IsDir() {

		p.appendSegment("golang", segment{
			content:    "\ufcd1",
			foreground: 254,
			background: 237,
		})
	}

}
