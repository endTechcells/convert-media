package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	formats := []string{"mp4", "m4a", "mp3", "wav", "mkv", "webm", "avi", "mov"}
	for i, f := range formats {
		fmt.Printf("%d: %s\n", i+1, f)
	}
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Archivo de entrada: ")
	file, _ := in.ReadString('\n')
	fmt.Print("Formato destino (número): ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(formats) {
		return
	}
	out := file[:len(file)-1] + "." + formats[n-1]
	cmd := exec.Command("ffmpeg", "-i", file[:len(file)-1], out)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
