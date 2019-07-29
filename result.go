package main

import (
	"sync"
)

type Result struct {
	cmap map[string]string
	sync.Mutex
}
