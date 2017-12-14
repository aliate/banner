package banner

import (
	"fmt"
)

type Banner []string

func getUppercase(a byte) Banner {
	switch a {
	case 'A':
		return Banner{
			`   ___ `,
			`  /   |`,
			` / _| |`,
			`/_/ |_|`,
		}
	case 'C':
		return Banner{
			`  ____`,
			` / __/`,
			`/ /___`,
			`\____/`,
		}
	case 'O':
		return Banner{
			`  ___  `,
			` / _ \ `,
			`| |_| |`,
			` \___/ `,
		}
	case 'E':
		return Banner{
			`   ____`,
			`  / __/`,
			` / _/  `,
			`/___/  `,
		}
	case 'S':
		return Banner{
			`  ___ `,
			` / __|`,
			`\__ \ `,
			`|___/ `,
		}
	case 'T':
		return Banner{
			` _____`,
			`/_  _/`,
			` / /  `,
			`/_/   `,
		}
	case 'F':
		return Banner{
			`   ____`,
			`  / __/`,
			` / __/ `,
			`/_/    `,
		}
	case 'B':
		return Banner{
			`   ___ `,
			`  / _ \`,
			` / _ \ `,
			`/____/ `,
		}
	case 'D':
		return Banner{
			`   __  `,
			`  / _ \`,
			` / _/ /`,
			`/____/ `,
		}
	case 'I':
		return Banner{
			`   __`,
			`  / /`,
			` / / `,
			`/_/  `,
		}
	case 'X':
		return Banner{
			`__  __`,
			`\ \/ /`,
			` |  | `,
			`/_/\_\`,
		}
	case 'W':
		return Banner{
			`___    _`,
			`| | _ //`,
			`| |//// `,
			`|_/\_/  `,
		}
	case 'Z':
		return Banner{
			`  _____`,
			` /__  /`,
			`  / /_ `,
			`/____/ `,
		}
	case 'R':
		return Banner{
			`   ___ `,
			`  /  _\`,
			` / /__/`,
			`/_/ \_\`,
		}
	case 'Y':
		return Banner{
			`__  _`,
			`\ \//`,
			` | | `,
			` |_| `,
		}
	case 'U':
		return Banner{
			`  __ __`,
			` / // /`,
			`/ // / `,
			`\___/  `,
		}
	case 'K':
		return Banner{

			`   __ _`,
			`  / ///`,
			` /   / `,
			`/_/\_\ `,
		}

	case 'J':
		return Banner{
			`     __`,
			`    / /`,
			` __/ / `,
			`/___/  `,
		}
	case 'H':
		return Banner{
			`   __ __`,
			`  / // /`,
			` / _  / `,
			`/_//_/  `,
		}
	case 'G':
		return Banner{
			`  ____ `,
			` / __/ `,
			`| ||__|`,
			` \___/ `,
		}
	case 'V':
		return Banner{

			`__    _`,
			`\ \  //`,
			` \ \// `,
			`  \_/  `,
		}
	case 'M':
		return Banner{

			`   __ __`,
			`  / |/ /`,
			` / //// `,
			`/_////  `,
		}
	case 'P':
		return Banner{
			`   ___ `,
			`  / _ \`,
			` / ___/`,
			`/_/    `,
		}
	case 'L':
		return Banner{

			`   __`,
			`  / /`,
			` / /_`,
			`/___/`,
		}
	case 'N':
		return Banner{
			`   __ __`,
			`  / // /`,
			` / /\ / `,
			`/_//_/  `,
		}
	case 'Q':
		return Banner{
			` ___  `,
			`/  _ \`,
			`| |_||`,
			`\____\`,
		}
	default:
		return Banner{
			``,``,``,``,
		}
	}
}

func getLowercase(a byte) Banner {
	switch a {
	case 'a':
		return Banner{
			`      `,
			` ___  `,
			`/ _ | `,
			`\____/`,
		}
	case 'c':
		return Banner{
			`     `,
			`  ___`,
			`/ __/`,
			`\__/ `,
		}
	case 'o':
		return Banner{
			`     `,
			` ___ `,
			`/ _ \`,
			`\___/`,
		}

	case 'e':
		return Banner{
			`     `,
			` ___ `,
			`/ __|`,
			`\___/`,
		}
	case 's':
		return Banner{
			`     `,
			` ____`,
			`/ __/`,
			`/___/`,
		}
	case 't':
		return Banner{
			`    __ `,
			` __/ /_`,
			`/_  __/`,
			` /__/  `,
		}
	case 'f':
		return Banner{
			`    ___`,
			` __/ _/`,
			`/_  _/ `,
			` /_/   `,
		}
	case 'b':
		return Banner{
			`   __ `,
			`  / / `,
			` / _ \`,
			`/____/`,
		}
	case 'd':
		return Banner{
			`    __`,
			` __/ /`,
			`/ _ / `,
			`\__/  `,
		}
	case 'i':
		return Banner{
			`  __`,
			` /_/`,
			` / /`,
			`/_/ `,
		}
	case 'x':
		return Banner{
			`     `,
			`__  _`,
			`\ \//`,
			`/_/\\`,
		}
	case 'w':
		return Banner{
			`       `,
			`  _ _ _`,
			` / ////`,
			`/_/\_/ `,
		}
	case 'z':
		return Banner{
			`     `,
			` ____`,
			`/_  /`,
			`/___/`,
		}
	case 'r':
		return Banner{
			`      `,
			`  ___ `,
			` / __\`,
			`/_/   `,
		}
	case 'y':
		return Banner{
			`     `,
			`__  _`,
			`\ \//`,
			`/__/ `,
		}
	case 'u':
		return Banner{
			`      `,
			` __ __`,
			`/ // /`,
			`\____/`,
		}
	case 'k':
		return Banner{
			`      `,
			`  __ _`,
			` / ///`,
			`/_/\_\`,
		}
	case 'j':
		return Banner{
			`    __`,
			`   /_/`,
			` __/ /`,
			`/___/ `,
		}
	case 'h':
		return Banner{
			`   __ `,
			`  / / `,
			` / _ \`,
			`/_//_/`,
		}
	case 'g':
		return Banner{
			`  ____`,
			` / _ /`,
			`_\_ / `,
			`\__/  `,
		}
	case 'v':
		return Banner{
			`     `,
			`__  _`,
			`\ \//`,
			` \_/ `,
		}
	case 'm':
		return Banner{
			`       `,
			`  __ __`,
			` / |/ /`,
			`/_//// `,
		}
	case 'p':
		return Banner{
			`   ___ `,
			`  / _ \`,
			` / ___/`,
			`/_/    `,
		}
	case 'l':
		return Banner{
			`   __`,
			`  / /`,
			` / /_`,
			`/___/`,
		}
	case 'n':
		return Banner{
			`      `,
			`  __  `,
			` / _ \`,
			`/_//_/`,
		}
	case 'q':
		return Banner{
			` _____`,
			`/ _  /`,
			`\_  / `,
			` /_/  `,
		}
	default:
		return Banner{
			``,``,``,``,
		}
	}
}

func getCase(a byte) Banner{
	if a >= 'a' && a <= 'z' {
		return getLowercase(a)
	}
	return getUppercase(a)
}

func GetBanner(str string) Banner {
	result := []Banner{}
	for i := 0; i < len(str); i++ {
		result = append(result, getCase(str[i]))
	}
	banner := Banner{}
	for i := 0; i < 4; i++ {
		ss := ""
		for _, v := range result {
			ss += v[i]
		}
		banner = append(banner, ss)
	}
	return banner
}

func (b Banner) Show() {
	for _, v := range b {
		fmt.Println(v)
	}
}

