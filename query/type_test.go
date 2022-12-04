package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	// "中古普明粵上海蘇州閩南閩東客朝越日吳日漢日訓通字"
	assert.Equal(t, "中古", LangType(TYPE_MC).String())
	assert.Equal(t, "普", LangType(TYPE_PU).String())
	assert.Equal(t, "明", LangType(TYPE_MG).String())
	assert.Equal(t, "粵", LangType(TYPE_CT).String())
	assert.Equal(t, "上海", LangType(TYPE_SH).String())
	assert.Equal(t, "蘇州", LangType(TYPE_SZ).String())
	assert.Equal(t, "閩南", LangType(TYPE_MN).String())
	assert.Equal(t, "閩東", LangType(TYPE_MD).String())
	assert.Equal(t, "客", LangType(TYPE_HK).String())
	assert.Equal(t, "朝", LangType(TYPE_KR).String())
	assert.Equal(t, "越", LangType(TYPE_VN).String())
	assert.Equal(t, "日吳", LangType(TYPE_JP_GO).String())
	assert.Equal(t, "日漢", LangType(TYPE_JP_KAN).String())
	assert.Equal(t, "日唐", LangType(TYPE_JP_TOU).String())
	assert.Equal(t, "日慣", LangType(TYPE_JP_KWAN).String())
	assert.Equal(t, "日他", LangType(TYPE_JP_OTHERS).String())
	assert.Equal(t, "日訓", LangType(TYPE_JP_YOMI).String())
	assert.Equal(t, "通字", LangType(TYPE_TZ).String())

	assert.Equal(t, LangType(TYPE_MC), NewLangType("mm"))
	assert.Equal(t, LangType(TYPE_MC), NewLangType("中古"))
	assert.Equal(t, LangType(TYPE_PU), NewLangType("普"))
	assert.Equal(t, LangType(TYPE_MG), NewLangType("明"))
	assert.Equal(t, LangType(TYPE_CT), NewLangType("粵"))
	assert.Equal(t, LangType(TYPE_SH), NewLangType("上海"))
	assert.Equal(t, LangType(TYPE_SZ), NewLangType("蘇州"))
	assert.Equal(t, LangType(TYPE_MN), NewLangType("閩南"))
	assert.Equal(t, LangType(TYPE_MD), NewLangType("閩東"))
	assert.Equal(t, LangType(TYPE_HK), NewLangType("客"))
	assert.Equal(t, LangType(TYPE_KR), NewLangType("朝"))
	assert.Equal(t, LangType(TYPE_VN), NewLangType("越"))
	assert.Equal(t, LangType(TYPE_JP_GO), NewLangType("日吳"))
	assert.Equal(t, LangType(TYPE_JP_KAN), NewLangType("日漢"))
	assert.Equal(t, LangType(TYPE_JP_TOU), NewLangType("日唐"))
	assert.Equal(t, LangType(TYPE_JP_KWAN), NewLangType("日慣"))
	assert.Equal(t, LangType(TYPE_JP_OTHERS), NewLangType("日他"))
	assert.Equal(t, LangType(TYPE_JP_YOMI), NewLangType("日訓"))
	assert.Equal(t, LangType(TYPE_TZ), NewLangType("通字"))
}
