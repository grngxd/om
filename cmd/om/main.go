// package main

// import (
// 	"fmt"
// 	"om/internal/lib/logger"
// 	"os"
// 	"strings"
// )

// func printBanner() {
// 	avatar := strings.Trim(`
// ██  ██  ██
// ██████████
//   ██  ██
// ██      ██
// ██      ██`, "\n")

// 	fmt.Printf("\n%s om 0.0.1a\n\n", avatar)
// }

// func main() {
// 	printBanner()

// 	if len(os.Args) < 2 {
// 		fmt.Println("usage: om <i|run|use> [args]")
// 		return
// 	}

// 	cmd := os.Args[1]

// 	switch cmd {
// 	case "i":
// 		if len(os.Args) < 3 {
// 			fmt.Println("usage: om i <package>")
// 			return
// 		}
// 		logger.Log("install:", os.Args[2])

// 	case "run":
// 		if len(os.Args) < 3 {
// 			fmt.Println("usage: om run <package>")
// 			return
// 		}
// 		logger.Log("run:", os.Args[2])

// 	case "use":
// 		if len(os.Args) < 3 {
// 			fmt.Println("usage: om use <distro>")
// 			return
// 		}
// 		logger.Log("use:", os.Args[2])

// 	default:
// 		logger.Log("unknown command:", cmd)
// 	}
// }

package main

import "om/internal/cli"

func main() {
	cli.Execute()
}
