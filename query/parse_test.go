package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const rawcontent = `<div class="container result-card" value="1867"><div class="row"><div class="col-xs-3 col-xs-char"><span class="h1">啊</span><br><span class="char-code"><small>U+554A</small></span><br/><br/></div><div class="col-xs-9 col-xs-detail"><div class="container"><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_pu.svg" alt="普" width="24px" height="24px"></td><td>ā, á, ǎ, à, a</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mg.svg" alt="明" width="24px" height="24px"></td><td>ō</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_ct.svg" alt="粵" width="24px" height="24px"></td><td>aa (āa), háa (áa), hàa (àa)</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sh.svg" alt="上海" width="24px" height="24px"></td><td>ā, ḣá</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mn.svg" alt="閩南" width="24px" height="24px"></td><td>[ah]</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_md.svg" alt="閩東" width="24px" height="24px"></td><td>ô</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_kr.svg" alt="朝" width="24px" height="24px"></td><td>아</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_vn.svg" alt="越" width="24px" height="24px"></td><td>a</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_go.svg" alt="日吳" width="24px" height="24px"></td><td>[1] あ</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_kan.svg" alt="日漢" width="24px" height="24px"></td><td>[1] あ</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_tou.svg" alt="日唐" width="24px" height="24px"></td><td>[2] あ</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_tz.svg" alt="通字" width="24px" height="24px"></td><td>a</td></tr></table></div></div></div></div></div></div><hr><div class="container result-card" value="1654"><div class="row"><div class="col-xs-3 col-xs-char"><span class="h1">呵</span><br><span class="char-code"><small>U+5475</small></span><br><span class="h5"><small>（啊）</small></span><br/><br/></div><div class="col-xs-9 col-xs-detail"><div class="container"><div class="row"><div class="col-xs-12 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mc.svg" alt="中古" width="24px" height="24px"></td><td>ha (曉果歌一開 下平5歌 虎何切 註:上同)<br>hà (曉果箇一開 去声21箇 呼箇切 註:噓氣呼箇切又呼哥切二)</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_pu.svg" alt="普" width="24px" height="24px"></td><td>hē, a, kē</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mg.svg" alt="明" width="24px" height="24px"></td><td>ho, hō</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_ct.svg" alt="粵" width="24px" height="24px"></td><td>ho (hó), o</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sh.svg" alt="上海" width="24px" height="24px"></td><td>hō</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sz.svg" alt="蘇州" width="24px" height="24px"></td><td>ha, hou</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mn.svg" alt="閩南" width="24px" height="24px"></td><td>[ə]</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_md.svg" alt="閩東" width="24px" height="24px"></td><td>o, ò, ho, hö, hö̀</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_hk.svg" alt="客" width="24px" height="24px"></td><td>ho²⁴, ho⁵⁵</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_kr.svg" alt="朝" width="24px" height="24px"></td><td>가, 하, 아</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_vn.svg" alt="越" width="24px" height="24px"></td><td>a, ha</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_go.svg" alt="日吳" width="24px" height="24px"></td><td>か</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_kan.svg" alt="日漢" width="24px" height="24px"></td><td>か</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_yomi.svg" alt="日訓" width="24px" height="24px"></td><td>しかる, わらう</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_tz.svg" alt="通字" width="24px" height="24px"></td><td>xo</td></tr></table></div></div></div></div></div></div>`

const decodedcontent = `[0]啊
	(0)普: ā, á, ǎ, à, a
	(1)明: ō
	(2)粵: aa (āa), háa (áa), hàa (àa)
	(3)上海: ā, ḣá
	(4)閩南: [ah]
	(5)閩東: ô
	(6)朝: 아
	(7)越: a
	(8)日吳: [1] あ
	(9)日漢: [1] あ
	(10)日唐: [2] あ
	(11)通字: a
[1]呵
	(0)中古: ha (曉果歌一開 下平5歌 虎何切 註:上同)
hà (曉果箇一開 去声21箇 呼箇切 註:噓氣呼箇切又呼哥切二)
	(1)普: hē, a, kē
	(2)明: ho, hō
	(3)粵: ho (hó), o
	(4)上海: hō
	(5)蘇州: ha, hou
	(6)閩南: [ə]
	(7)閩東: o, ò, ho, hö, hö̀
	(8)客: ho²⁴, ho⁵⁵
	(9)朝: 가, 하, 아
	(10)越: a, ha
	(11)日吳: か
	(12)日漢: か
	(13)日訓: しかる, わらう
	(14)通字: xo
`

const rawcontent2 = `<div class="container result-card" value="39"><div class="row"><div class="col-xs-3 col-xs-char"><span class="h1">並</span><br><span class="char-code"><small>U+4E26</small></span><br/><br/></div><div class="col-xs-9 col-xs-detail"><div class="container"><div class="row"><div class="col-xs-12 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mc.svg" alt="中古" width="24px" height="24px"></td><td>ḃéng (並梗迥四開 上声24迥 蒲迥切 註:上同)</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_pu.svg" alt="普" width="24px" height="24px"></td><td>bìng</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mg.svg" alt="明" width="24px" height="24px"></td><td>bing</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_ct.svg" alt="粵" width="24px" height="24px"></td><td>bịng (bīng)</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sh.svg" alt="上海" width="24px" height="24px"></td><td>ḃín</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sz.svg" alt="蘇州" width="24px" height="24px"></td><td>bin, bìn, ḃı̂n</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mn.svg" alt="閩南" width="24px" height="24px"></td><td><span class="light">pịng</span>, <strong>bịng</strong></td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_md.svg" alt="閩東" width="24px" height="24px"></td><td>bǐng</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_hk.svg" alt="客" width="24px" height="24px"></td><td>bin⁵⁵</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_kr.svg" alt="朝" width="24px" height="24px"></td><td>병</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_vn.svg" alt="越" width="24px" height="24px"></td><td>tịnh</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_go.svg" alt="日吳" width="24px" height="24px"></td><td><span class="light">びょう (びゃう)</span></td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_kan.svg" alt="日漢" width="24px" height="24px"></td><td><strong>へい</strong></td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_tz.svg" alt="通字" width="24px" height="24px"></td><td>bhieg</td></tr></table></div></div></div></div></div></div><hr><div class="container result-card" value="11486"><div class="row"><div class="col-xs-3 col-xs-char"><span class="h1">竝</span><br><span class="char-code"><small>U+7ADD</small></span><br><span class="h5"><small>（並）</small></span><br/><br/></div><div class="col-xs-9 col-xs-detail"><div class="container"><div class="row"><div class="col-xs-12 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mc.svg" alt="中古" width="24px" height="24px"></td><td>ḃéng (並梗迥四開 上声24迥 蒲迥切 註:比也蒲迥切四)</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_pu.svg" alt="普" width="24px" height="24px"></td><td>bìng</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mg.svg" alt="明" width="24px" height="24px"></td><td>bing</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sh.svg" alt="上海" width="24px" height="24px"></td><td>ḃín</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_md.svg" alt="閩東" width="24px" height="24px"></td><td>běng</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_kr.svg" alt="朝" width="24px" height="24px"></td><td>병, 방, 반</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_vn.svg" alt="越" width="24px" height="24px"></td><td>tịnh</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_go.svg" alt="日吳" width="24px" height="24px"></td><td><span class="light">びょう (びゃう)</span></td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_kan.svg" alt="日漢" width="24px" height="24px"></td><td><strong>へい</strong></td></tr></table></div></div></div></div></div></div><hr><div class="container result-card" value="4215"><div class="row"><div class="col-xs-3 col-xs-char"><span class="h1">并</span><br><span class="char-code"><small>U+5E76</small></span><br><span class="h5"><small>（並）</small></span><br/><br/></div><div class="col-xs-9 col-xs-detail"><div class="container"><div class="row"><div class="col-xs-12 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mc.svg" alt="中古" width="24px" height="24px"></td><td>bieng (幫梗清三開 下平8庚 府盈切 註:合也亦州名舜分冀州爲幽州并州春秋時爲晉國後屬趙秦爲太原郡魏復置并州又姓出姓苑府盈切四)<br>bièng (幫梗勁三開 去声24敬 畀政切 註:專也)</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_pu.svg" alt="普" width="24px" height="24px"></td><td>bìng, bīng</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_mg.svg" alt="明" width="24px" height="24px"></td><td>bing, bīng</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_ct.svg" alt="粵" width="24px" height="24px"></td><td>bịng (bīng), bing</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sh.svg" alt="上海" width="24px" height="24px"></td><td>ḃín, bīn, bin</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_md.svg" alt="閩東" width="24px" height="24px"></td><td>bǐng</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_kr.svg" alt="朝" width="24px" height="24px"></td><td>병</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_vn.svg" alt="越" width="24px" height="24px"></td><td>tịnh, tinh</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_go.svg" alt="日吳" width="24px" height="24px"></td><td><span class="light">ひょう (ひゃう)</span></td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_kan.svg" alt="日漢" width="24px" height="24px"></td><td>へい</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_yomi.svg" alt="日訓" width="24px" height="24px"></td><td>あわせ, ならぶ</td></tr></table></div></div></div></div></div></div><hr><div class="container result-card" value="4216"><div class="row"><div class="col-xs-3 col-xs-char"><span class="h1">幷</span><br><span class="char-code"><small>U+5E77</small></span><br><span class="h5"><small>（並）</small></span><br/><br/></div><div class="col-xs-9 col-xs-detail"><div class="container"><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_pu.svg" alt="普" width="24px" height="24px"></td><td>bìng, bīng</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_sh.svg" alt="上海" width="24px" height="24px"></td><td>ḃín, bīn, bin</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_md.svg" alt="閩東" width="24px" height="24px"></td><td>bǐng</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_kr.svg" alt="朝" width="24px" height="24px"></td><td>병</td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_vn.svg" alt="越" width="24px" height="24px"></td><td>tịnh, bình</td></tr></table></div><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_go.svg" alt="日吳" width="24px" height="24px"></td><td><span class="light">ひょう (ひゃう)</span></td></tr></table></div></div><div class="row"><div class="col-xs-6 col-xs-orthography"><table><tr style="text-align: left;"><td><img src="res/drawable/lang_jp_kan.svg" alt="日漢" width="24px" height="24px"></td><td>へい</td></tr></table></div></div></div></div></div></div>`

const decodedcontent2 = `[0]並
	(0)中古: ḃéng (並梗迥四開 上声24迥 蒲迥切 註:上同)
	(1)普: bìng
	(2)明: bing
	(3)粵: bịng (bīng)
	(4)上海: ḃín
	(5)蘇州: bin, bìn, ḃı̂n
	(6)閩南: pịng, bịng
	(7)閩東: bǐng
	(8)客: bin⁵⁵
	(9)朝: 병
	(10)越: tịnh
	(11)日吳: びょう (びゃう)
	(12)日漢: へい
	(13)通字: bhieg
[1]竝
	(0)中古: ḃéng (並梗迥四開 上声24迥 蒲迥切 註:比也蒲迥切四)
	(1)普: bìng
	(2)明: bing
	(3)上海: ḃín
	(4)閩東: běng
	(5)朝: 병, 방, 반
	(6)越: tịnh
	(7)日吳: びょう (びゃう)
	(8)日漢: へい
[2]并
	(0)中古: bieng (幫梗清三開 下平8庚 府盈切 註:合也亦州名舜分冀州爲幽州并州春秋時爲晉國後屬趙秦爲太原郡魏復置并州又姓出姓苑府盈切四)
bièng (幫梗勁三開 去声24敬 畀政切 註:專也)
	(1)普: bìng, bīng
	(2)明: bing, bīng
	(3)粵: bịng (bīng), bing
	(4)上海: ḃín, bīn, bin
	(5)閩東: bǐng
	(6)朝: 병
	(7)越: tịnh, tinh
	(8)日吳: ひょう (ひゃう)
	(9)日漢: へい
	(10)日訓: あわせ, ならぶ
[3]幷
	(0)普: bìng, bīng
	(1)上海: ḃín, bīn, bin
	(2)閩東: bǐng
	(3)朝: 병
	(4)越: tịnh, bình
	(5)日吳: ひょう (ひゃう)
	(6)日漢: へい
`

func TestParse(t *testing.T) {
	content, err := NewContent([]byte(rawcontent))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, decodedcontent, content.String())
	content, err = NewContent([]byte(rawcontent2))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, decodedcontent2, content.String())
}
