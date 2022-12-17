package main

import (
	"fmt"
	"strings"
)

// Todo: translate colors
func Format(f string, addr string) string {
	f = strings.ReplaceAll(f, "{server.name}", fmt.Sprintf("%v", addr))
	f = strings.ReplaceAll(f, "{server.ping}", fmt.Sprintf("%d", STATUS.Delay.Milliseconds()))
	f = strings.ReplaceAll(f, "{server.motd}", fmt.Sprintf("%v", STATUS.Description.ClearString()))
	f = strings.ReplaceAll(f, "{players.now}", fmt.Sprintf("%v", STATUS.Players.Online))
	f = strings.ReplaceAll(f, "{players.max}", fmt.Sprintf("%v", STATUS.Players.Max))
	f = strings.ReplaceAll(f, "{version.name}", fmt.Sprintf("%v", STATUS.Version.Name))
	f = strings.ReplaceAll(f, "{version.proto}", fmt.Sprintf("%v", STATUS.Version.Protocol))

	// Fill it up to 2 lines to avoid errors
	l := strings.Split(STATUS.Description.ClearString(), "\n")
	for len(l) < 2 {
		l = append(l, "")
	}

	f = strings.ReplaceAll(f, "{server.motd:0}", fmt.Sprintf("%v", l[0]))
	f = strings.ReplaceAll(f, "{server.motd:1}", fmt.Sprintf("%v", l[1]))

	// Trim away whitespace on sides
	f = strings.ReplaceAll(f, "{server.motd:0.trim}", fmt.Sprintf("%v", strings.TrimSpace(l[0])))
	f = strings.ReplaceAll(f, "{server.motd:1.trim}", fmt.Sprintf("%v", strings.TrimSpace(l[1])))

	return f
}
