/*
Author Axoford12
Team Freeze
Org Rubiginosu
  _____                        ____
|  ___| __ ___ _______ _ __  / ___| ___
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/

 */

package initial

import (
	"fmt"
	"time"
)

const VERSION string = "v0.0"

func printInfo() {
	fmt.Println("  _____                        ____")
	fmt.Println("|  ___| __ ___ _______ _ __  / ___| ___")
	fmt.Println("| |_ | '__/ _\\_  / _ \\ '_ \\| |  _ / _ \\")
	fmt.Println("|  _|| | | (_) / /  __/ | | | |_| | (_) |")
	fmt.Println("|_|  |_|\\___/___\\___|_| |_|\\____|\\___/")
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------")
	time.Sleep(100 * time.Microsecond)
	fmt.Print("Powered by ")
	for _,v := range []byte("Axoford12"){
		time.Sleep(300 * time.Millisecond)
		fmt.Print(string(v))
	}
	fmt.Println()
	time.Sleep(1000 * time.Millisecond)
	time.Sleep(100 * time.Microsecond)
	fmt.Println("---------------------")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("version:" + VERSION)
}
