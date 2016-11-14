package transcode

import (
	"io"
	"os/exec"
	"strings"
)

// Transcode takes in a file location, runs ffmpeg with a pipe output
// and the data is collected from the io.Reader
func Transcode(in string) (out io.ReadCloser, err error) {

	c := ffmpegTranscodeCommand(in)

	parts := strings.Split(c, " ")

	head, parts := parts[0], parts[1:]

	cmd := exec.Command(head, parts...)

	rd, err := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return rd, err

}
