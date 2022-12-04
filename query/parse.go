package query

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
)

const (
	langcharregstr = `<div class="col-xs-3 col-xs-char"><span class="h1">([一-龥])`
	langinforegstr = `<td><img src="res/drawable/lang_\w+\.svg" alt="([一-龥\w]+)" width="24px" height="24px"></td><td>([^<>]+)</td>`
)

var (
	langcharre = regexp.MustCompile(langcharregstr)
	langinfore = regexp.MustCompile(langinforegstr)
)

var (
	ErrDataTooShort = errors.New("data too short")
	ErrNilChar      = errors.New("nil char")
	ErrInvChar      = errors.New("invalid char")
	ErrZeroLang     = errors.New("zero lang")
	ErrInvLang      = errors.New("invalid lang")
)

func NewContent(data []byte) (content Content, err error) {
	data = bytes.ReplaceAll(data, []byte("<br>"), []byte("\n"))
	data = bytes.ReplaceAll(data, []byte("<br />"), []byte("\n"))
	data = bytes.ReplaceAll(data, []byte("<br/>"), []byte("\n"))
	data = bytes.ReplaceAll(data, []byte(`<span class="light">`), nil)
	data = bytes.ReplaceAll(data, []byte(`</span>`), nil)
	data = bytes.ReplaceAll(data, []byte(`<strong>`), nil)
	data = bytes.ReplaceAll(data, []byte(`</strong>`), nil)
	if len(data) < 256 {
		err = ErrDataTooShort
		return
	}
	cards := bytes.Split(data[41:], []byte(`<div class="container result-card" value=`))
	if len(cards) == 0 {
		return
	}
	content.Cards = make([]Card, len(cards))
	for i, card := range cards {
		char := langcharre.FindSubmatch(card)
		if len(char) == 0 {
			err = ErrNilChar
			return
		}
		content.Cards[i].Char = []rune(string(char[1]))[0]
		if !content.Cards[i].IsCharValid() {
			err = ErrInvChar
			return
		}
		lang := langinfore.FindAllSubmatch(card, -1)
		if len(lang) == 0 {
			err = ErrZeroLang
			return
		}
		content.Cards[i].Langs = make([]Lang, len(lang))
		for j, l := range lang {
			lang := Lang{Type: NewLangType(string(l[1])), Info: string(l[2])}
			if lang.Type >= TYPE_END {
				fmt.Println(string(card))
				return content, ErrInvLang
			}
			content.Cards[i].Langs[j] = lang
		}
	}
	return
}
