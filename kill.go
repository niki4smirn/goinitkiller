package killer

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Init function called. Terminating program...")
	os.Exit(1)
}
