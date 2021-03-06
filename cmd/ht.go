package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"github.com/e10ulen/qqw/lib"
)
const DateFormat = "2006/01/02 15:04"
func init(){
	RootCmd.AddCommand(gitcommitCmd)
}

var gitcommitCmd = &cobra.Command{
	Use:	"gc",
	Short:	"Auto Commit And Push.",
	Run:	func(cmd *cobra.Command, args []string){
		tm := time.Now()
		dir, err := os.Getwd()
		if err == nil {
			fmt.Print(dir, "\n")
		}
		fmt.Print("DateTime:", tm.Format(DateFormat), "\n")
		//	Add Git UnStageFiles
		add, err := exec.Command("git", "add", "--all").CombinedOutput()
		lib.Check(err)
		//	Git Commit
		cmt,err := exec.Command("git", "commit", "-m", "Commit:"+tm.Format(DateFormat)).CombinedOutput()
		lib.Check(err)
		push, err :=exec.Command("git", "push").CombinedOutput()
		lib.Check(err)
	fmt.Print("Git Add:", string(add), "\n")
	fmt.Print("Git Commit:", string(cmt), "\n")
	fmt.Print("Git Push:", string(push), "\n")
	},
}

