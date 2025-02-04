package main

import "time"

func updateLastModifiedTime() {
	lastModifiedTime = time.Now().Format(time.RFC3339)
}
