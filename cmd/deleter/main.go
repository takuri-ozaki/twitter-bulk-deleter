package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"twitter-bulk-deleter/internal/prompt"
	"twitter-bulk-deleter/internal/tweet"
	"twitter-bulk-deleter/internal/util"
)

func main() {
	flag.Parse()
	skipPrompt := false
	start := ""
	end := ""
	if len(flag.Args()) >= 2 {
		start = flag.Arg(0)
		end = flag.Arg(1)
		skipPrompt = true
	}
	execute := false
	if flag.Arg(0) == "execute" || flag.Arg(2) == "execute" {
		execute = true
	}

	prompter, err := prompt.NewPrompter(start, end, skipPrompt)
	if err != nil {
		fmt.Println("invalid input")
		return
	}
	prompter.StartPrompt()
	prompter.EndPrompt()
	if !prompter.ConfirmPrompt() {
		fmt.Println("bye")
		return
	}

	err = godotenv.Load(fmt.Sprintf("./.env"))
	if err != nil {
		panic(err)
	}

	targets, errorCount, err := tweet.Parse(prompter.Start(), prompter.End())
	if err != nil {
		panic(err)
	}
	if errorCount != 0 {
		fmt.Println(fmt.Sprint(errorCount) + " tweets failed to parse")
	}

	api := util.NewRealApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	deleter := tweet.NewDeleter(api, execute)
	deleter.Delete(targets)
	prompter.Finish()
}
