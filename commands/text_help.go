package commands

const (
	// DefaultErrorText - displayed when invalide commands are passed
	DefaultErrorText string = "Please supply a valid command, -h, or --help switch"

	// DefaultHelpText - displayed when no commands are passed
	DefaultHelpText string = `Welcome to the vsts_cli

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

    project     Manage VSTS projects
    build       Manage VSTS builds
    release 	Manage VSTS releases
    `
)
