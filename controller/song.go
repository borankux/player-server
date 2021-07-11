package controller

import "github.com/borankux/filemaster/scans"

type SongController struct {}

func GetAll() []scans.FileInfo{
	return scans.Walk("song", 10)
}