package main

import (
	"log"
	"strings"

	"github.com/golang/gddo/gosrc"
)

var (
	scrawDeamon *deamon
)

type deamon struct {
	c chan string
}

func (d *deamon) Push(path string) {
	d.c <- path
}

func (d *deamon) Start() {
	go func() {
		for {
			path := <-d.c
			log.Println("Deamon Recive:", path)
			if !strings.HasSuffix(path, "/vendor") && (gosrc.IsValidRemotePath(path) || gosrc.IsGoRepoPath(path)) {
				p, _, _ := getDoc(path, queryRequest)
				if p != nil {
					log.Println("Deamon", path, p.Etag)
				}
			}
		}
	}()
}
