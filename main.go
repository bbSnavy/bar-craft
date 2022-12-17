package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
)

const (
	TIMEOUT = time.Second * 5
)

var (
	DEBUG  = flag.Bool("d", false, "Print debug information")
	STATUS status
)

type status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
	Favicon string
	Delay   time.Duration
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s <address> <format>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	var (
		addr   string
		format string
		data   []byte
		err    error
	)

	addr = flag.Arg(0)
	if addr == "" {
		fmt.Println(GetError("E_ADDR", errors.New("invalid address")))
		os.Exit(1)
	}

	format = flag.Arg(1)
	if format == "" {
		fmt.Println(GetError("E_FMT", errors.New("invalid format")))
		os.Exit(1)
	}

	data, STATUS.Delay, err = bot.PingAndListTimeout(addr, TIMEOUT)
	if err != nil {
		fmt.Println(GetError("E_PING", err))
		os.Exit(1)
	}

	err = json.Unmarshal(data, &STATUS)
	if err != nil {
		fmt.Println(GetError("E_JSON", err))
		os.Exit(1)
	}

	format = Format(format, addr)

	fmt.Println(format)
}
