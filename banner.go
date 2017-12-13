package main

import (
	"fmt"
)

func getLowercase(a byte) []string {
	switch string(a) {
	case "a":
		return []string{
			`      `,
			` ___  `,
			`/ _ | `,
			`\____/`,
		}
	case "c":
		return []string{
			`     `,
			`  ___`,
			`/ __/`,
			`\__/ `,
		}
	case "o":
		return []string{
			`     `,
			` ___ `,
			`/ _ \`,
			`\___/`,
		}

	case "e":
		return []string{
			`     `,
			` ___ `,
			`/ __|`,
			`\___/`,
		}
	case "s":
		return []string{
			`     `,
			` ___ `,
			`/ __\`,
			`|___/`,
		}
	case "t":
		return []string{
			`    __ `,
			` __/ /_`,
			`/_  __/`,
			` /__/  `,
		}
	case "f":
		return []string{
			`    ___`,
			` __/ _/`,
			`/_  _/ `,
			` /_/   `,
		}
	case "b":
		return []string{
			`   __ `,
			`  / / `,
			` / _ \`,
			`/____/`,
		}
	case "d":
		return []string{
			`    __`,
			` __/ /`,
			`/ _ / `,
			`\__/  `,
		}
	case "i":
		return []string{
			`  __`,
			` /_/`,
			` / /`,
			`/_/ `,
		}
	case "x":
		return []string{
			`     `,
			`__  _`,
			`\ \//`,
			`/_/\\`,
		}
	case "w":
		return []string{
			`       `,
			`  _ _ _`,
			` / ////`,
			`/_/\_/ `,
		}
	case "z":
		return []string{
			`     `,
			` ____`,
			`/_  /`,
			`/___/`,
		}
	case "r":
		return []string{
			`      `,
			`  ___ `,
			` / __\`,
			`/_/   `,
		}
	case "y":
		return []string{
			`     `,
			`__  _`,
			`\ \//`,
			`/__/ `,
		}
	case "u":
		return []string{
			`      `,
			` __ __`,
			`/ // /`,
			`\____/`,
		}
	case "k":
		return []string{
			`      `,
			`  __ _`,
			` / ///`,
			`/_/\_\`,
		}
	case "j":
		return []string{
			`    __`,
			`   /_/`,
			` __/ /`,
			`/___/ `,
		}
	case "h":
		return []string{
			`   __ `,
			`  / / `,
			` / _ \`,
			`/_//_/`,
		}
	case "g":
		return []string{
			`  ____`,
			` / _ /`,
			`_\_ / `,
			`\__/  `,
		}
	case "v":
		return []string{
			`     `,
			`__  _`,
			`\ \//`,
			` \_/ `,
		}
	case "m":
		return []string{
			`       `,
			`  __ __`,
			` / |/ /`,
			`/_//// `,
		}
	case "p":
		return []string{
			`   ___ `,
			`  / _ \`,
			` / ___/`,
			`/_/    `,
		}
	case "l":
		return []string{
			`   __`,
			`  / /`,
			` / /_`,
			`/___/`,
		}
	case "n":
		return []string{
			`      `,
			`  __  `,
			` / _ \`,
			`/_//_/`,
		}
	case "q":
		return []string{
			` _____`,
			`/ _  /`,
			`\_  / `,
			` /_/  `,
		}
	default:
		return []string{
			``,
			``,
			``,
			``,
		}
	}
}

func GetBanner(str string) []string {
	result := [][]string{}
	for i := 0; i < len(str); i++ {
		result = append(result, getLowercase(str[i]))
	}
	banner := []string{}
	for i := 0; i < 4; i++ {
		ss := ""
		for _, v := range result {
			ss += v[i]
		}
		banner = append(banner, ss)
	}
	return banner
}

func main() {
	ss := "autopattern"
	result := GetBanner(ss)
	for _, v := range result {
		fmt.Println(v)
	}
}
