package main

import (
	"bytes"

	"util.mleku.dev/lol"
)

type (
	B = []byte
	S = string
	E = error
	N = int
)

var (
	log, chk, errorf = lol.Main.Log, lol.Main.Check, lol.Main.Errorf
	equals           = bytes.Equal
)
