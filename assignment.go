package main

import "time"

type Assignment struct {
	product Product
	site    Site
	start   time.Time
	end     time.Time
}
