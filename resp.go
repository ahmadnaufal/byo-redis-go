package main

import "strings"

const (
	Terminator = "\r\n"

	Echo = "ECHO"
)

func handlePayload(payload string) (string, error) {
	tokens := strings.Fields(payload)
	cmd, args := tokens[0], tokens[1:]
	return handleCommand(cmd, args)
}

func handleCommand(cmd string, args []string) (string, error) {
	switch strings.ToUpper(cmd) {
	case Echo:
		return args[0], nil
	default:
		return "", ErrUnrecognizedCommand
	}
}
