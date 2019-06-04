package main

import (
	"os"
)

const rustcargoFile = "./Cargo.toml"

func segmentRustLang(p *powerline) {
	stat, err := os.Stat(rustcargoFile)
	if err == nil && !stat.IsDir() {

		p.appendSegment("rustlang", segment{
			content:    "\ue7a8",
			foreground: 237,
			background: 254,
		})
	}

}
