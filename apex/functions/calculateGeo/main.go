package main

import (
	"github.com/River-Island/go-util"
	"github.com/apex/go-apex"
)

func main() {
	apex.HandleFunc(util.CalculateGeo())
}
