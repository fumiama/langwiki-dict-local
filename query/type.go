package query

import "strings"

const (
	TYPE_MC        = iota // TYPE_MC 中古
	TYPE_PU               // TYPE_PU 普
	TYPE_MG               // TYPE_MG 明
	TYPE_CT               // TYPE_CT 粤
	TYPE_SH               // TYPE_SH 上海
	TYPE_SZ               // TYPE_SZ 苏州
	TYPE_MN               // TYPE_MN 闽南
	TYPE_MD               // TYPE_MD 闽东
	TYPE_HK               // TYPE_HK 客家
	TYPE_KR               // TYPE_KR 朝鲜
	TYPE_VN               // TYPE_VN 越南
	TYPE_JP_GO            // TYPE_JP_GO 吴音
	TYPE_JP_KAN           // TYPE_JP_KAN 汉音
	TYPE_JP_TOU           // TYPE_JP_TOU 唐音
	TYPE_JP_KWAN          // TYPE_JP_KWAN 惯用音
	TYPE_JP_OTHERS        // TYPE_JP_OTHERS 日他
	TYPE_JP_YOMI          // TYPE_JP_YOMI 训读
	TYPE_TZ               // TYPE_TZ 通字
	TYPE_END              // TYPE_END 边界
)

type LangType uint8

const langstr = "中古普明粵上海蘇州閩南閩東客朝越日吳日漢日唐日慣日他日訓通字"

var langstrlen = [...]int{0, 6, 9, 12, 15, 21, 27, 33, 39, 42, 45, 48, 54, 60, 66, 72, 78, 84, 90}

func (lt LangType) String() string {
	if lt >= TYPE_END {
		return ""
	}
	return langstr[langstrlen[lt]:langstrlen[lt+1]]
}

func NewLangType(name string) LangType {
	if name == "mm" { // 特判
		return TYPE_MC
	}
	i := strings.Index(langstr, name)
	if i < 0 {
		return TYPE_END
	}
	for n, l := range langstrlen {
		if i == l {
			return LangType(n)
		}
	}
	return TYPE_END
}
