package args

import (
	"os"
	"vsts/utils/logger"
)

// printDefaultHelp displays the default help command
func printDefaultHelp() {
	logger.LogMessage(`Welcome to the vsts cli

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

Microsoft (R) Visual Studio (R) Team Foundation Server

	build		Work with VSTS builds
	release 	Work with VSTS Releases
			`)
	os.Exit(0)
}
