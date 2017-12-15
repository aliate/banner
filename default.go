package banner

type DefaultTheme struct{}

func (d DefaultTheme) Convert(b byte) Ascii {
	switch b {
	// Start Uppercase convert:
	case 'A':
		return Ascii{
			`   ___ `,
			`  /   |`,
			` / _| |`,
			`/_/ |_|`,
			`       `,
		}
	case 'C':
		return Ascii{
			`  ____`,
			` / __/`,
			`/ /___`,
			`\____/`,
			`      `,
		}
	case 'O':
		return Ascii{
			`  ___ `,
			` / _ \`,
			`| |_||`,
			`\___/ `,
			`       `,
		}
	case 'E':
		return Ascii{
			`   ____`,
			`  / __/`,
			` / _/  `,
			`/___/  `,
			`       `,
		}
	case 'S':
		return Ascii{
			`  ___ `,
			` / __|`,
			`\__ \ `,
			`|___/ `,
			`      `,
		}
	case 'T':
		return Ascii{
			` _____`,
			`/_  _/`,
			` / /  `,
			`/_/   `,
			`      `,
		}
	case 'F':
		return Ascii{
			`   ____`,
			`  / __/`,
			` / __/ `,
			`/_/    `,
			`       `,
		}
	case 'B':
		return Ascii{
			`   ___ `,
			`  / _ \`,
			` / _  /`,
			`/____/ `,
			`       `,
		}
	case 'D':
		return Ascii{
			`   ___ `,
			`  / _ \`,
			` / /_//`,
			`/____/ `,
			`       `,
		}
	case 'I':
		return Ascii{
			`   __`,
			`  / /`,
			` / / `,
			`/_/  `,
			`     `,
		}
	case 'X':
		return Ascii{
			` __  __`,
			` \ \/ /`,
			` |   | `,
			`/_/\_\ `,
			`       `,
		}
	case 'W':
		return Ascii{
			`___    _`,
			`| | _ //`,
			`| |//// `,
			`|_/\_/  `,
			`        `,
		}
	case 'Z':
		return Ascii{
			`  _____`,
			` /__  /`,
			`  / /_ `,
			`/____/ `,
			`       `,
		}
	case 'R':
		return Ascii{
			`   ____ `,
			`  / _  \`,
			` / _  / `,
			`/_/ \_\ `,
			`        `,
		}
	case 'Y':
		return Ascii{
			`__  _`,
			`\ \//`,
			` / / `,
			`/_/  `,
			`     `,
		}
	case 'U':
		return Ascii{
			`  __ __`,
			` / // /`,
			`/ // / `,
			`\___/  `,
			`       `,
		}
	case 'K':
		return Ascii{
			`   __ _`,
			`  / ///`,
			` /   / `,
			`/_/\_\ `,
			`       `,
		}

	case 'J':
		return Ascii{
			`     __`,
			`    / /`,
			` __/ / `,
			`/___/  `,
			`       `,
		}
	case 'H':
		return Ascii{
			`   __ __`,
			`  / // /`,
			` / _  / `,
			`/_//_/  `,
			`        `,
		}
	case 'G':
		return Ascii{
			`  ____`,
			` / __/`,
			`| ||_|`,
			`\___/ `,
			`      `,
		}
	case 'V':
		return Ascii{
			`__   _`,
			`| | //`,
			`| |// `,
			`|__/  `,
			`      `,
		}
	case 'M':
		return Ascii{
			`   __ ___`,
			`  / |/  /`,
			` /  _  / `,
			`/_/ /_/  `,
			`         `,
		}
	case 'P':
		return Ascii{
			`   ___ `,
			`  / _ \`,
			` / ___/`,
			`/_/    `,
			`       `,
		}
	case 'L':
		return Ascii{
			`   __`,
			`  / /`,
			` / /_`,
			`/___/`,
			`     `,
		}
	case 'N':
		return Ascii{
			`   __ __`,
			`  / // /`,
			` /  \ / `,
			`/_//_/  `,
			`        `,
		}
	case 'Q':
		return Ascii{
			`  ____ `,
			` / __ \`,
			`| |__||`,
			`\___\_\`,
			`       `,
		}
		// Start Lowercase convert:
	case 'a':
		return Ascii{
			`      `,
			` ___  `,
			`/ _ | `,
			`\____/`,
			`      `,
		}
	case 'c':
		return Ascii{
			`     `,
			` ____`,
			`/ __/`,
			`\__/ `,
			`     `,
		}
	case 'o':
		return Ascii{
			`     `,
			` ___ `,
			`/ _ \`,
			`\___/`,
			`     `,
		}

	case 'e':
		return Ascii{
			`     `,
			` ___ `,
			`/ __\`,
			`\___/`,
			`     `,
		}
	case 's':
		return Ascii{
			`     `,
			` ____`,
			`/ __\`,
			`/___/`,
			`     `,
		}
	case 't':
		return Ascii{
			`    __ `,
			` __/ /_`,
			`/_  __/`,
			` /__/  `,
			`       `,
		}
	case 'f':
		return Ascii{
			`    ___`,
			` __/ _/`,
			`/_  _/ `,
			` / /   `,
			`/_/    `,
		}
	case 'b':
		return Ascii{
			`   __ `,
			`  / / `,
			` / _ \`,
			`/____/`,
			`      `,
		}
	case 'd':
		return Ascii{
			`    __`,
			` __/ /`,
			`/ _ / `,
			`\__/  `,
			`      `,
		}
	case 'i':
		return Ascii{
			`  __`,
			` /_/`,
			` / /`,
			`/_/ `,
			`    `,
		}
	case 'x':
		return Ascii{
			`      `,
			` __  _`,
			` \ \//`,
			`/_/\\ `,
			`      `,
		}
	case 'w':
		return Ascii{
			`       `,
			`__ _ __`,
			`| /// /`,
			`|_/\_/ `,
			`       `,
		}
	case 'z':
		return Ascii{
			`      `,
			`  ____`,
			` /_  /`,
			`/___/ `,
			`      `,
		}
	case 'r':
		return Ascii{
			`      `,
			`  ___ `,
			` / __\`,
			`/_/   `,
			`      `,
		}
	case 'y':
		return Ascii{
			`      `,
			` __  _`,
			` \ \//`,
			`  / / `,
			` /_/  `,
		}
	case 'u':
		return Ascii{
			`      `,
			` __ __`,
			`/ // /`,
			`\____/`,
			`      `,
		}
	case 'k':
		return Ascii{
			`   __ `,
			`  / /_`,
			` /  //`,
			`/_/\_\`,
			`      `,
		}
	case 'j':
		return Ascii{
			`     __`,
			`    /_/`,
			`    / /`,
			` __/ / `,
			`/___/  `,
		}
	case 'h':
		return Ascii{
			`   __ `,
			`  / / `,
			` / _ \`,
			`/_//_/`,
			`      `,
		}
	case 'g':
		return Ascii{
			`      `,
			`  ____`,
			` / _ /`,
			`_\_ / `,
			`\__/  `,
		}
	case 'v':
		return Ascii{
			`     `,
			`__  _`,
			`| \//`,
			`|__/ `,
			`     `,
		}
	case 'm':
		return Ascii{
			`       `,
			`  __ __`,
			` / |/ /`,
			`/_//// `,
			`       `,
		}
	case 'p':
		return Ascii{
			`       `,
			`   ___ `,
			`  / _ \`,
			` / ___/`,
			`/_/    `,
		}
	case 'l':
		return Ascii{
			`  __ `,
			` / / `,
			`/ /_ `,
			`\___/`,
			`     `,
		}
	case 'n':
		return Ascii{
			`      `,
			`  __  `,
			` / _ \`,
			`/_//_/`,
			`      `,
		}
	case 'q':
		return Ascii{
			`      `,
			` _____`,
			`/ _  /`,
			`\_  / `,
			` /_/  `,
		}
	default:
		return Ascii{
			` `, ` `, ` `, ` `, ` `,
		}
	}
}
