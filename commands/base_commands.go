package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

func sanitizeCommand(command string) string {
	// Get rid of weird capitals
	command = strings.ToLower(command)
	// Allow one or two dashes
	if len(command) < 2 {
		return command
	}
	if command[0:2] == "--" {
		return command[2:]
	}
	if command[0:1] == "-" {
		return command[1:]
	}
	return command
}

func IsBaseCommand(args []string) bool {
	if len(args) != 2 {
		return false
	}
	return string(args[1][0]) == "-"
}

func HandleBaseCommand(cfg *config.Config, args []string) error {
	command := sanitizeCommand(args[1])
	switch command {
	case Checkup, CheckupShort:
		needsUpdate, err := CmdCheckup(cfg)
		if err != nil {
			return err
		}
		if !needsUpdate {
			fmt.Println("owo you're up to date :)")
		} else {
			fmt.Println("Please run: owo -update")
		}
	case About, AboutShort:
		CmdAbout()
	case Update, UpdateShort:
		err := CmdUpdate(cfg)
		if err != nil {
			return err
		}
	case Config, ConfigShort:
		err := CmdConfig(cfg)
		if err != nil {
			return err
		}
	case Help, HelpShort:
		CmdHelp()
	case Sync, SyncShort:
		err := sync(cfg)
		if err != nil {
			return err
		}
		os.Exit(0) // Wont run defer
	case Reset, ResetShort:
		err := reset(cfg)
		if err != nil {
			return err
		}
		os.Exit(0) // Wont run defer
	case Diff:
		err := checkDiff(cfg)
		if err != nil {
			return err
		}
		os.Exit(0) // Wont run defer
	default:
		fmt.Println("command", args[1], "not recognized. Perhaps you need owo -help ?")
	}
	return nil
}
