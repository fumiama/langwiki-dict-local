package query

import (
	"strconv"
	"strings"
)

// Content 查询结果
type Content struct {
	Cards []Card // Cards 所有匹配字
}

func (c Content) String() string {
	sb := strings.Builder{}
	for i, card := range c.Cards {
		sb.WriteByte('[')
		sb.WriteString(strconv.Itoa(i))
		sb.WriteByte(']')
		sb.WriteRune(card.Char)
		sb.WriteByte('\n')
		for j, lang := range card.Langs {
			sb.WriteByte('\t')
			sb.WriteByte('(')
			sb.WriteString(strconv.Itoa(j))
			sb.WriteByte(')')
			sb.WriteString(lang.Type.String())
			sb.WriteString(": ")
			sb.WriteString(lang.Info)
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

// Card 一个字
type Card struct {
	Char  rune   // Char 字
	Langs []Lang // Langs 各种读音
}

func (c *Card) IsCharValid() bool {
	return c.Char >= 0x4e00 && c.Char <= 0x9fa5
}

// Lang 某语言
type Lang struct {
	Type LangType // Type is TYPE_xx
	Info string   // Info 读音和解释
}
