package ars

import (
	"math/rand"
	"time"
)

var (
	LevelDefault = New(1)
	LevelWeak    = New(2)
	LevelMedium  = New(3)
	LevelStrong  = New(4)
	// LevelHash can't be decrypted
	LevelHash = New(5)
)

// Level is a security level structure
type Level struct {
	Spacing        int
	FirstStrBreak  string
	SecondStrBreak string
}

// New returns a new Level
func New(level int) Level {
	var spacing int
	var firstStrBreak, secondStrBreak string
	if level <= 4 {
		spacing = level * 10
		firstStrBreak = genSecString(level)
		secondStrBreak = genSecString(level)
	} else {
		spacing = level * 9
	}
	return Level{
		Spacing:        spacing,
		FirstStrBreak:  firstStrBreak,
		SecondStrBreak: secondStrBreak,
	}
}

// genSecString returns a random string
func genSecString(level int) string {
	rand.Seed(time.Now().UnixNano())
	var secStr = "`"
	for i := 0; i != level; i++ {
		secStr += string(characters[rand.Intn(len(characters)-1)])
	}
	return secStr
}
