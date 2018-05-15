package main

import (
	"log"
	"os"
	"os/signal"
	"github.com/Mi7teR/shavronne/discord"
	"syscall"
	"flag"
)

func main() {
	var discordToken string
	flag.StringVar(&discordToken, "discordToken", "exampleToken", "discord auth token")
	flag.Parse()
	sc := make(chan os.Signal, 1)
	dg, err := discord.Run(discordToken)
	if err != nil {
		log.Println(err)
	}
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}
