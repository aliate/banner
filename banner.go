package banner

import (
	"fmt"
)

type Ascii []string

func getUppercase(a byte) Ascii {
	switch a {
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
			`  ___  `,
			` / _ \ `,
			`| |_| |`,
			` \___/ `,
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
			` / _ \ `,
			`/____/ `,
			`       `,
		}
	case 'D':
		return Ascii{
			`   __  `,
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
			`__  __`,
			`\ \/ /`,
			` |  | `,
			`/_/\_\`,
			`      `,
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
			`   ___ `,
			`  /  _\`,
			` / /__/`,
			`/_/ \_\`,
			`       `,
		}
	case 'Y':
		return Ascii{
			`__  _`,
			`\ \//`,
			` | | `,
			` |_| `,
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
			`  ____ `,
			` / __/ `,
			`| ||__|`,
			` \___/ `,
			`       `,
		}
	case 'V':
		return Ascii{
			`__    _`,
			`\ \  //`,
			` \ \// `,
			`  \_/  `,
			`       `,
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
			` ___   `,
			`/  _ \ `,
			`| |_|| `,
			`\___/_\`,
			`       `,
		}
	default:
		return Ascii{
			` `, ` `, ` `, ` `, ` `,
		}
	}
}

func getLowercase(a byte) Ascii {
	switch a {
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
			`  ___`,
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
			`     `,
			`__  _`,
			`\ \//`,
			`/_/\\`,
			`     `,
		}
	case 'w':
		return Ascii{
			`       `,
			`  _ _ _`,
			` / ////`,
			`/_/\_/ `,
			`       `,
		}
	case 'z':
		return Ascii{
			`     `,
			` ____`,
			`/_  /`,
			`/___/`,
			`     `,
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
			`     `,
			` _  _`,
			` \\//`,
			` / / `,
			`/_/  `,
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
			`      `,
			`  __ _`,
			` / ///`,
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
			`\ \//`,
			` \_/ `,
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
			`   __`,
			`  / /`,
			` / /_`,
			`/___/`,
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

func getAscii(a byte) Ascii {
	if a >= 'a' && a <= 'z' {
		return getLowercase(a)
	}
	return getUppercase(a)
}

func getAsciiBlankInfo(ascii Ascii, fromBack bool) []int {
	BlankInfo := []int{}
	for _, v := range ascii {
		num := 0
		for i := 0; i < len(v); i++ {
			idx := i
			if fromBack {
				idx = len(v) - i - 1
			}
			if v[idx] == ' ' {
				num++
			} else {
				break
			}
		}
		BlankInfo = append(BlankInfo, num)
	}
	return BlankInfo
}

const MaxInt = int(^uint(0) >> 1)

func getAsciiUnionSpace(firstBlankInfo, secondBlankInfo []int) int {
	space := MaxInt
	for i, v := range firstBlankInfo {
		if space > v + secondBlankInfo[i] {
			space = v + secondBlankInfo[i]
		}
	}
	return space
}

type UnionRecord struct {
	BackPos     int
	ForwardPos  int
	First		byte
	Second		byte
}

func isShouldAdjoin(a byte, b byte) bool {
	return (a == '/' || a == '_') &&
			(b == '/' || b == '_') ||
			a == ' ' ||
			b == ' '
}

func (b Ascii) isEmpty() bool {
	for _, v := range b {
		for _, i  := range v {
			if i != ' ' {
				return false
			}
		}
	}
	return true
}

func (b Ascii) Append(a Ascii, adjoin bool) {
	if b.isEmpty() || a.isEmpty() {
		for i, _ := range b {
			b[i] = b[i] + a[i]
		}
	} else {
		b.joinOne(a, adjoin)
	}
}

func (b Ascii) joinOne(a Ascii, adjoin bool) {
	back := getAsciiBlankInfo(b, true)
	forward := getAsciiBlankInfo(a, false)
	erase := getAsciiUnionSpace(back, forward)

	unionRecords := []UnionRecord{}
	for i, _ := range b {
		var backPos, forwardPos int
		if erase > back[i] {
			backPos = back[i]
			forwardPos = erase - backPos
		} else {
			backPos = erase
			forwardPos = 0
		}
		unionRecords = append(unionRecords, UnionRecord{
			BackPos: backPos,
			ForwardPos: forwardPos,
			First: b[i][len(b[i])-backPos-1],
			Second: a[i][forwardPos],
		})
	}

	hasPair := false
	shouldAraseMoreOne := true
	for _, record := range unionRecords {
		if record.First == '/' && record.Second == '/' {
			hasPair = true
		}
		if !isShouldAdjoin(record.First, record.Second) {
			shouldAraseMoreOne = false
		}
	}

	if !hasPair {
		shouldAraseMoreOne = false
	}


	for i, _ := range b {
		record := unionRecords[i]
		backGet := len(b[i]) - record.BackPos
		forwardGet := record.ForwardPos
		if shouldAraseMoreOne && adjoin {
			if b[i][backGet-1] == ' ' {
				backGet -= 1
			} else {
				forwardGet += 1
			}
		}
		b[i] = b[i][:backGet] + a[i][forwardGet:]
	}
}

func (b Ascii) Show() {
	for _, v := range b {
		fmt.Println(v)
	}
}

type Banner struct {
	Origin	string
	Adjoin	bool
	Ascii	Ascii
}

func NewBanner(str string, adjoin bool) *Banner {
	b := &Banner{
		Origin: str,
		Adjoin: adjoin,
	}
	ascii := getAscii(str[0])
	for i := 1; i < len(str); i++ {
		ascii.Append(getAscii(str[i]), adjoin)
	}
	b.Ascii = ascii
	return b
}

func (b *Banner) Show() {
	b.Ascii.Show()
}
