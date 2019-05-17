package main

import (
	"os"
)

const pkgfile = "./package.json"

func segmentNode(p *powerline) {
	stat, err := os.Stat(pkgfile)
	if err == nil && !stat.IsDir() {

		p.appendSegment("node-version", segment{
			content:    "\uf898",
			foreground: p.theme.NodeFg,
			background: p.theme.NodeBg,
		})
	}

}
