package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultHelpText string = `
	Welcome to the vsts_cli

                      $$$$$$$$$$$
       ?????        $$$$$$$$ZZZ$$$
    ~?++==++?I    7$$$$$$$$$$Z???Z$
   ~?+==~~~=+I$Z 777$$$$$$$  +OOOOO$
   ??+=        $77777$$$I     $ZZZZZ
   7I?        II777777?       77$$$$
    ~77      +IIIII777         777777
     II~  =IIIIIIIIIZZ        III777
      77IIIIIIIIII   77II?   IIIIII
         7IIIII        $777777777     (TM)

Microsoft (R) Visual Studio (R) Team Foundation Server`
)

var rootCmd = &cobra.Command{
	Use:   "vsts",
	Short: "Welcome to the vsts_cli",
	Long:  defaultHelpText,
}

// Execute starts the fun
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
