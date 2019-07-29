package main

import (
	"sync"
)

//type PageDetail struct {
//	title string
//	links []strings
//	err   error
//}

type Result struct {
	cmap map[string]string
	sync.Mutex
}
