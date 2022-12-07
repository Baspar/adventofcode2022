package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Day interface {
	Init(input []string) error
	Part1() (string, error)
	Part2() (string, error)
}

var frames = []string{"⡆", "⠇", "⠋", "⠙", "⠸", "⢰", "⣠", "⣄"}

// ⠁⠂⠃⠄⠅⠆⠇⡀⡁⡂⡃⡄⡅⡆⡇⠈⠉⠊⠋⠌⠍⠎⠏⡈⡉⡊⡋⡌⡍⡎⡏⠐⠑⠒⠓⠔⠕⠖⠗⡐⡑⡒⡓⡔⡕⡖⡗⠘⠙⠚⠛⠜⠝⠞⠟⡘⡙⡚⡛⡜⡝⡞⡟⠠⠡⠢⠣⠤⠥⠦⠧⡠⡡⡢⡣⡤⡥⡦⡧⠨⠩⠪⠫⠬⠭⠮⠯⡨⡩⡪⡫⡬⡭⡮⡯⠰⠱⠲⠳⠴⠵⠶⠷⡰⡱⡲⡳⡴⡵⡶⡷⠸⠹⠺⠻⠼⠽⠾⠿⡸⡹⡺⡻⡼⡽⡾⡿⢀⢁⢂⢃⢄⢅⢆⢇⣀⣁⣂⣃⣄⣅⣆⣇⢈⢉⢊⢋⢌⢍⢎⢏⣈⣉⣊⣋⣌⣍⣎⣏⢐⢑⢒⢓⢔⢕⢖⢗⣐⣑⣒⣓⣔⣕⣖⣗⢘⢙⢚⢛⢜⢝⢞⢟⣘⣙⣚⣛⣜⣝⣞⣟⢠⢡⢢⢣⢤⢥⢦⢧⣠⣡⣢⣣⣤⣥⣦⣧⢨⢩⢪⢫⢬⢭⢮⢯⣨⣩⣪⣫⣬⣭⣮⣯⢰⢱⢲⢳⢴⢵⢶⢷⣰⣱⣲⣳⣴⣵⣶⣷⢸⢹⢺⢻⢼⢽⢾⢿⣸⣹⣺⣻⣼⣽⣾⣿

func SanitizeInput(input string) []string {
	return strings.Split(strings.Trim(input, "\n"), "\n")
}

func Pstr(s string) *string {
	return &s
}

func Time[T any](title string, dim bool, f func(chan<- T, chan<- error), display func(T) *string) error {
	var (
		start time.Time
		err   error
		ans   T
	)

	currentFrame := 0
	ansChan := make(chan T, 1)
	errChan := make(chan error, 1)
	ticker := time.NewTicker(100 * time.Millisecond)

	defer ticker.Stop()
	defer close(ansChan)
	defer close(errChan)

	start = time.Now()
	go f(ansChan, errChan)

	shade := 1
	if dim {
		shade = 0
	}

	fmt.Printf("\033[%d;33m%s: %s\033[0m", shade, title, frames[currentFrame])

	for {
		select {
		case <-ticker.C:
			currentFrame = (currentFrame + 1) % len(frames)
			fmt.Printf("\r\033[%d;33m%s: %s\033[0m", shade, title, frames[currentFrame])
			continue
		case ans = <-ansChan:
			fmt.Printf("\r\033[%d;32m%s: ✓\033[0m (%s)", shade, title, time.Since(start))
			if prettyAns := display(ans); prettyAns != nil {
				fmt.Println("\n\n" + *prettyAns)
			}
		case err = <-errChan:
			fmt.Printf("\r\033[%d;31m%s: 𐄂\033[0m (%s)\n\n", shade, title, time.Since(start))
			fmt.Println(err)
		}
		fmt.Println("")
		return err
	}
}

func Run(day Day, re_init bool) {
	var (
		err       error
		stdinInfo os.FileInfo
		content   []byte
	)

	stdin := os.Stdin
	if stdinInfo, err = stdin.Stat(); err != nil {
		fmt.Printf("Cannot analyze stdin %s\n", err)
		return
	}

	if (stdinInfo.Mode() & os.ModeCharDevice) == 0 {
		if content, err = io.ReadAll(stdin); err != nil {
			fmt.Printf("Cannot read stdin %s\n", err)
		} else {
			fmt.Print("Using stdin\n\n")
		}
	} else if content, err = os.ReadFile("./input.txt"); err != nil {
		fmt.Printf("Input reading failed: %s\n", err)
		return
	}

	lines := SanitizeInput(string(content))

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	processPart := func(part int, f func() (string, error), init bool) error {
		if init {
			if err := Time(
				"Init",
				true,
				func(ansChan chan<- bool, errChan chan<- error) {
					if err = day.Init(lines); err != nil {
						errChan <- err
					} else {
						ansChan <- true
					}
				},
				func(_ bool) *string {
					return nil
				},
			); err != nil {
				return err
			}
		}

		return Time(
			fmt.Sprintf("Part%d", part),
			false,
			func(ansChan chan<- string, errChan chan<- error) {
				if ans, err := f(); err != nil {
					errChan <- err
				} else {
					ansChan <- ans
				}
			},
			func(ans string) *string {
				return &ans
			},
		)
	}

	if processPart(1, day.Part1, true) != nil {
		return
	}
	if processPart(2, day.Part2, re_init) != nil {
		return
	}
}
