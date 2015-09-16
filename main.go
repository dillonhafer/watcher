package main

import (
	"log"

	"github.com/go-fsnotify/fsnotify"
	"os/exec"
)

func SortDates() error {
	return exec.Command("ruby", "/Users/dhafer/camera/dates.rb").Run()
}

func ExampleNewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("added file:", event.Name)
					SortDates()
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/Users/dhafer/camera")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	ExampleNewWatcher()
}
