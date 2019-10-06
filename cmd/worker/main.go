package main

import (
	"time"

	"github.com/bzon/monobuild-example/pkg/db"
	"github.com/docker/docker/pkg/namesgenerator"
)

func main() {
	for {
		db.Add(namesgenerator.GetRandomName(3))
		time.Sleep(1 * time.Second)
	}
}
