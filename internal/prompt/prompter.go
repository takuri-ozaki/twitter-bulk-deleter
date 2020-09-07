package prompt

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

const standardFormat = "2006-01-02 15:04:05"

type Prompter struct {
	start      time.Time
	end        time.Time
	skipPrompt bool
}

func NewPrompter(start, end string, skipPrompt bool) (*Prompter, error) {
	if !skipPrompt {
		return &Prompter{}, nil
	}
	startTime, err := time.Parse(standardFormat, start)
	if err != nil {
		return &Prompter{}, err
	}
	endTime, err := time.Parse(standardFormat, end)
	if err != nil {
		return &Prompter{}, err
	}
	if endTime.Before(startTime) {
		return &Prompter{}, errors.New("end time must be after start time")
	}
	return &Prompter{start: startTime, end: endTime, skipPrompt: true}, nil
}

func (p *Prompter) StartPrompt() {
	if p.skipPrompt {
		return
	}
	fmt.Println("input start time")
	fmt.Println("ex) " + standardFormat)
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		start := scanner.Text()
		var err error
		p.start, err = time.Parse(standardFormat, start)
		if err == nil {
			break
		}
		fmt.Println("please input in the correct format")
	}
}

func (p *Prompter) EndPrompt() {
	if p.skipPrompt {
		return
	}
	fmt.Println("input end time")
	fmt.Println("ex) " + standardFormat)
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		end := scanner.Text()
		var err error
		p.end, err = time.Parse(standardFormat, end)
		if err == nil && p.end.After(p.start) {
			break
		} else if err == nil {
			fmt.Println("end time must be after start time")
		}
		fmt.Println("please input in the correct format")
	}
}

func (p *Prompter) ConfirmPrompt() bool {
	fmt.Println("delete tweets " + p.start.Format(standardFormat) + " ~ " + p.end.Format(standardFormat))
	for {
		fmt.Println("input y/n")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		yn := scanner.Text()
		if yn == "y" {
			return true
		}
		if yn == "n" {
			return false
		}
	}
}

func (p *Prompter) Start() time.Time {
	return p.start
}

func (p *Prompter) End() time.Time {
	return p.end
}

func (p *Prompter) Finish() {
	fmt.Println("deleted tweets " + p.start.Format(standardFormat) + " ~ " + p.end.Format(standardFormat))
}
