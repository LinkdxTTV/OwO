package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/LinkdxTTV/owo/config"
)

const (
	Sync      string = "sync"
	SyncShort string = "s"
)

func sync(cfg *config.Config) error {

	diffFiles, err := numDiff(cfg)
	if err != nil {
		return err
	}
	if diffFiles == 0 {
		fmt.Println("No changes detected. Nothing to sync owo")
		return nil
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)
	err = os.Chdir(cfg.LocalPath)
	if err != nil {
		return err
	}

	// New Branch

	// Name convention is git username and time

	gitName := exec.Command("git", "config", "user.name")
	outName, err := gitName.Output()
	if err != nil {
		return err
	}
	name := strings.Split(string(outName), " ")[0]
	time := time.Now()
	branchName := "owo" + "/" + name + "/" + time.Format("2006-01-02/15.04.05")

	err = exec.Command("git", "checkout", "-b", branchName).Run()
	defer exec.Command("git", "checkout", "main").Run()

	err = exec.Command("git", "add", "./docs/docs").Run()
	if err != nil {
		return err
	}
	err = exec.Command("git", "commit", "-m", "auto generated by owo").Run()
	if err != nil {
		return err
	}
	err = exec.Command("git", "push", "origin", branchName).Run()
	if err != nil {
		return err
	}

	fmt.Println("Changed synced succesfully: Check it out at: ")
	fmt.Println("https://" + cfg.Git.RemoteURL + "/pulls")
	return nil
}
