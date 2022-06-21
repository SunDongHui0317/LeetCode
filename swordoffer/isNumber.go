package swordoffer

const (
	StateInit = iota
	StateSign
	StateInt
	StatePoint
	StatePointWithoutInt
	StateFraction
	StateExp
	StateExpSign
	StateExpInt
	StateEnd
)

const (
	CharSpace = iota
	CharSign
	CharNum
	CharExp
	CharPoint
	CharEnd
)

var CharType = map[string]int{
	"0": CharNum,
	"1": CharNum,
	"2": CharNum,
	"3": CharNum,
	"4": CharNum,
	"5": CharNum,
	"6": CharNum,
	"7": CharNum,
	"8": CharNum,
	"9": CharNum,
	" ": CharSpace,
	"-": CharSign,
	"+": CharSign,
	"e": CharExp,
	"E": CharExp,
	".": CharPoint,
}

var StateChar = map[int]map[int]int{
	StateInit: {
		CharSpace: StateInit,
		CharSign:  StateSign,
		CharNum:   StateInt,
		CharPoint: StatePointWithoutInt,
	},
	StateSign: {
		CharNum:   StateInt,
		CharPoint: StatePointWithoutInt,
	},
	StateInt: {
		CharSpace: StateEnd,
		CharNum:   StateInt,
		CharExp:   StateExp,
		CharPoint: StatePoint,
		CharEnd:   StateEnd,
	},
	StatePoint: {
		CharSpace: StateEnd,
		CharNum:   StateFraction,
		CharExp:   StateExp,
		CharEnd:   StateEnd,
	},
	StatePointWithoutInt: {
		CharNum: StateFraction,
	},
	StateFraction: {
		CharSpace: StateEnd,
		CharNum:   StateFraction,
		CharExp:   StateExp,
		CharEnd:   StateEnd,
	},
	StateExp: {
		CharSign: StateExpSign,
		CharNum:  StateExpInt,
	},
	StateExpSign: {
		CharNum: StateExpInt,
	},
	StateExpInt: {
		CharSpace: StateEnd,
		CharNum:   StateExpInt,
		CharEnd:   StateEnd,
	},
	StateEnd: {
		CharSpace: StateEnd,
		CharEnd:   StateEnd,
	},
}

func isNumber(s string) bool {
	stat := StateInit
	var ok bool
	var tmpStat int
	for i := 0; i < len(s); i++ {
		cur := string(s[i])
		charType := CharType[cur]
		if _, ok = CharType[cur]; !ok {
			return false
		}
		if tmpStat, ok = StateChar[stat][charType]; ok {
			stat = tmpStat
		} else {
			return false
		}
	}
	if _, ok = StateChar[stat][CharEnd]; ok {
		return true
	}

	return false
}
