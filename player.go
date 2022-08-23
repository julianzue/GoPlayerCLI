package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/mitchellh/colorstring"
)

func main() {

	var path string

	colorstring.Print("[yellow][+] [white]Enter absolute path: ")
	fmt.Scan(&path)

	f, _ := os.Open(path)

	filesize, _ := os.Stat(path)

	streamer, format, _ := mp3.Decode(f)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)

	full := strings.Split(path, "/")

	s := fmt.Sprintf("%.2f", float64(filesize.Size())/(1024*1000))

	title := full[len(full)-1]

	colorstring.Println("[light_cyan][*] [white]Playing:[light_cyan] " + title + " [white](" + s + "MB)")

	for {
	}

}
