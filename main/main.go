package main

import (
	"messenger"
)

func main() {
	m, err := messenger.NewMessenger()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
