// Component: FlexDict
// File flexdict.js
// Copyright 2017 Langwiki.org

var version = 1.2;
var disabled_table = [
	new Array(0, 0, 0),
	new Array(1, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1),
	new Array(0, 1, 1)
];

function tabs(part) {
	var tab = ['dictionary', 'converter', 'multidict', 'notepad', 'setting', 'instruction'];
	for (var i = tab.length - 1; i >= 0; i--) {
		$('#'+tab[i]).attr('class', '');
		$('#'+tab[i]+'-wrapper').css('display', 'none');
	}
	$('#'+part).attr('class', 'active');
	$('#'+part+'-wrapper').css('display', 'block');

	// refresh
	if (part == 'notepad') {
		refresh_dictmemo();
	}
}

function switch_init() {
	$.fn.bootstrapSwitch.defaults.size = "small";
	$.fn.bootstrapSwitch.defaults.onColor = "success";
	$.fn.bootstrapSwitch.defaults.onText = "開";
	$.fn.bootstrapSwitch.defaults.offText = "關";
	// initialization
	$("#switch-0").bootstrapSwitch('state', false);
	$("#switch-1").bootstrapSwitch('state', true);
	$("#switch-2").bootstrapSwitch('state', true);
	//$("#switch-2").bootstrapSwitch('disabled', true);
}

function convert_switch_init() {
	//$("#convert-switch-0").bootstrapSwitch('state', false);
}

function change_mode(id) {
	$("#dropdown-selector").val(id);
	$("#dropdown-selector").html($('#selector-'+id).html()+' <span class="caret"></span>');
	for (var i = 0; i < 3; i++) {
		$("#switch-"+i).bootstrapSwitch('disabled', disabled_table[id][i]);
	}
}

function change_dict(id) {
	$("#dropdown-selector-lang").val(id);
	$("#dropdown-selector-lang").html($('#selector-lang-'+id).html()+' <span class="caret"></span>');
}

function change_convert_mode(id) {
		$("#convert-dropdown-selector").val(id);
		$("#convert-dropdown-selector").html($('#convert-selector-'+id).html()+' <span class="caret"></span>');
}

function change_convert_match_mode(id) {
	$("#convert-match-dropdown-selector").val(id);
	$("#convert-match-dropdown-selector").html($('#convert-match-selector-'+id).html()+' <span class="caret"></span>');
}

function set_display(id, lang) {
	$("#dropdown-"+lang).val(id);
	$("#dropdown-"+lang).html($('#'+lang+'-'+id).html()+' <span class="caret"></span>');
}

function getSettings() {
	return [$("#dropdown-pu").val(), // 0
					$("#dropdown-ct").val(), // 1
					$("#dropdown-kr").val(), // 2
					$("#dropdown-vn").val(), // 3
					$("#dropdown-jp").val(), // 4
					$("#dropdown-ltco").val(), // 5
					$("#dropdown-ltcp").val(), // 6
					$("#convert-match-dropdown-selector").val(), //7
					$("#dropdown-mn").val(), // 8
					$("#dropdown-wu").val(), // 9
				];
}

function queryWordZh(word, mode, flag, bot, success) {
	var string = word;
	var setting = getSettings();
	// use: console.log("search: ["+string+"] mode="+mode+", flag="+flag+", setting="+setting+ "...");

	// Query the database
	$.post("php/query.php",
		{
			string: string,
			mode: mode,
			flag: flag,
			setting: setting,
			bot: bot
		},
		function(data, status) {
			// status: "success", "notmodified", "error", "timeout", or "parsererror"
			if (status == "success") {
				if (data.status == "success") {
					success(data.HTML);
				} else {
					success("未找到符合條件的漢字。");
				}
			} else if (status == "timeout") {
				alert("Timeout...");
			} else if (status == "error") {
				alert("Error...");
			} else if (status == "notmodified") {
				alert("Notmodified...");
			} else if (status == "parsererror") {
				alert("Parsererror...");
			}
		},
	"json");
}

function queryMultiDictWord(lang, word, flag, bot, showBookmark, success) {
	var string = word;
	var setting = [
		"0"     // Show manchu characters for manchu
	];

	// Change settings based on lang
	if (lang == "mnc") {
		setting[0] = "1";
	}

	// use: console.log("search: ["+string+"] mode="+mode+", flag="+flag+", setting="+setting+ "...");

	// Query the multi database
	$.post("php/querydict.php",
		{
			lang: lang,
			string: string,
			mode: "0", // dummy
			flag: flag,
			setting: setting,
			bot: bot,
			bm: showBookmark, // "0" do not show bookmark
		},
		function(data, status) {
			// status: "success", "notmodified", "error", "timeout", or "parsererror"
			if (status == "success") {
				if (data.status == "success") {
					success(data.HTML);
				} else {
					success("未找到符合條件的詞語。");
				}
			} else if (status == "timeout") {
				alert("Timeout...");
			} else if (status == "error") {
				alert("Error...");
			} else if (status == "notmodified") {
				alert("Notmodified...");
			} else if (status == "parsererror") {
				alert("Parsererror...");
			}
		},
	"json");
}

// Queries definition of words for multiple-languages
// Chinese character lookup is a virtual language 'zh'
function queryWord(lang, word, mode, flag, bot, bm, success) {
	switch (lang) {
		case "zh":
		case "汉字":
		{
			queryWordZh(word, mode, flag, bot, success);
			break;
		}

		case "tw":
		case "中文（繁体）": // not longer needed
			queryMultiDictWord("tw", word, flag, bot, bm, success);
			break;

		case "cn":
		case "汉语（简体）":
			queryMultiDictWord("cn", word, flag, bot, bm, success);
			break;

		case "manxi":
			queryMultiDictWord("mnc", word, flag, bot, bm, success);
			break;

		default:
			queryMultiDictWord(lang, word, flag, bot, bm, success);
			break;
	}
}

function query() {
	var string = $.trim($("#_search_").val());
	if (string == "") { // query empty, return empty
		$("#results").html("");
		$("#_search_").val("");
		return;
	}

	var mode = $("#dropdown-selector").val();
	var flag = new Array();
	for (var i = 2; i >= 0; i--) {
		flag[i] = ~disabled_table[mode][i] & $("#switch-"+i).bootstrapSwitch('state');
	}

	var bot = "0";
	var bm = "1";
	queryWord("zh", string, mode, flag, bot, bm,
			function (res) { $("#results").html(res);}
	);
}

function showMultiDictResult(lang, content) {
	var topBottom = lang == "manxi";

	if (topBottom) {
		// set top-to-bottom style
		topDownContent = '<div class="top-to-bottom" style="font-size: 18px; height:300px">' + content + '</div>';
		$("#multidict-results").html(topDownContent);
	} else {
		// set left-to-right style
		$("#multidict-results").html(content);
	}
}

function queryMultiDict(bm) {
	var string = $.trim($("#_search_multi_dict_").val());
	if (string == "") { // query empty, return empty
		$("#multidict-results").html("");
		$("#_search_multi_dict_").val("");
		return;
	}

	var lang = $("#dropdown-selector-lang").val();
	var flag = new Array();
	/*
	for (var i = 2; i >= 0; i--) {
		flag[i] = ~disabled_table[mode][i] & $("#switch-"+i).bootstrapSwitch('state');
	}*/

	var mode = "0";
	var bot = "0";
	queryWord(lang, string, mode, flag, bot, bm,
			function (res) { showMultiDictResult(lang, res); });
}

function getMemoTable() {
	return $('#dictmemo');
}

function refresh_dictmemo() {
	getMemoTable().bootstrapTable('refresh');
}

function dictMemoFormatter(index, row) {
		var word = row['word'];
		var lang = row['langid'];

		var mode = "0";
		var flag = [
							"0", // only
							"0", // variants
							"1"  // annotation
					];

		var $table = getMemoTable();
		$table.on('expand-row.bs.table', function (e, idx, row, $detail) {
			if (idx != index) return;

		  $detail.html('载入数据...');
			var bot = "1";
			var bm = "0";
			queryWord(lang, word, mode, flag, bot, bm, function (res) {
				wrapper = res.replace(/\n/g, '<br>');
				$detail.html(wrapper);
	    });
		});
}

function htmlDecode(input){
  var e = document.createElement('div');
  e.innerHTML = input;
  return e.childNodes[0].nodeValue;
}

function convert() {
	var string = $.trim($("#_converter_").val());
	if (string == "") { // query empty, return empty
		$("#results").html("");
		$("#_search_").val("");
		return;
	}
	var mode = $("#convert-dropdown-selector").val();
	var setting = getSettings();
	// use: console.log("search: ["+string+"] mode="+mode+", flag="+flag+", setting="+setting+ "...");

	// Query the database
	$.post("php/convert.php",
		{
			string: string,
			mode: mode,
			setting: setting
		},
		function(data, status) {
			// status: "success", "notmodified", "error", "timeout", or "parsererror"
		  if (status == "success") {
				if (data.status == "success") {
					$("#_convert_results_").val(htmlDecode(data.HTML));
				} else {
					$("#_convert_results_").val("（抱歉，出错了）");
				}
			} else if (status == "timeout") {
				alert("Timeout...");
			} else if (status == "error") {
				alert("Error...");
			} else if (status == "notmodified") {
				alert("Notmodified...");
			} else if (status == "parsererror") {
				alert("Parsererror...");
			}
		},
	"json");
}

function convertToolbarHandler(event) {
	var convertRes = $("#_convert_results_");
	if (!convertRes) return;
	convertRes.focus(); // for IE
	var sel = convertRes.getSelection();

	if (sel.length == 0) {
		convert_toolbar_update();
		return;
	}

	var text = sel.text;
	var newText = "";
	switch (event.data) {
		case "toUpper":
		  newText = convertCase(text, 1);
			break;
		case "toLower":
		  newText = convertCase(text, 0);
			break;
		case "toProper":
		  newText = convertCase(text, 2);
		  break;
	}
	convertRes.replaceSelectedText(newText);

	convert_toolbar_update();
}

function convert_toolbar_init() {
	$("#convert_tool_to_upper").click("toUpper", convertToolbarHandler);
	$("#convert_tool_to_lower").click("toLower", convertToolbarHandler);
	$("#convert_tool_to_proper").click("toProper", convertToolbarHandler);

	$(document).on("selectionchange", function() {
	  convert_toolbar_update();
	});
}

function buttonUpdate(buttonId) {
	var convertRes = $("#_convert_results_");
	if (!convertRes) return;
	if (convertRes.getSelection == undefined) return;

	var sel = convertRes.getSelection();
	if (sel == null) return;
	var disableCaseButtons = sel.length == 0 ? "disable" : "enable";
	var color = sel.length == 0 ? "lightgray" : "black";

	var btn = $(buttonId);
	btn.button(disableCaseButtons);
	btn.css('color', color);
}

function convert_toolbar_update() {
	buttonUpdate("#convert_tool_to_upper");
	buttonUpdate("#convert_tool_to_lower");
	buttonUpdate("#convert_tool_to_proper");
}

/**
 * Converts the string text to lowercase (mode = 0), uppercase (mode = 1), or
 * proper case (mode = 2, the first character of each word is uppercase otherwise
 * lowercase).
 *
 * Please support basic latin letters and the following special characters:
 *   Lowercase: ḃ ḋ ġ ṡ ż ȷ̈ ẋ ḣ ṗ ṫ k̇
 *   Uppercase: Ḃ Ḋ Ġ Ṡ Ż J̈ Ẋ Ḣ Ṗ Ṫ K̇
 *   Lowercase: l̥ n̥ m̥ ü ë ɿ ö ä ṙ ñ
 *   Uppercase: L̥ N̥ M̥ Ü Ë 𐐓/⅂ Ö Ä Ṙ Ñ
 */
function convertCase(text, mode) {
	var lines = text.split("\n");
	var res = "";
	for (var i = 0; i < lines.length; ++i) {
		if (i != 0) res += "\n";
		var words = lines[i].split(" ");
		for (var j = 0; j < words.length; ++j) {
			words[j] = convertCaseWord(words[j], mode);
		}
		res += words.join(" ");
	}
	return res;
}

/******************** CASE CONVERSION UTILITY from 葉劍飛 ********************/

function toUpperCase(text)
{
	var ret = '';
	for (var i = 0; i < text.length; ++i)
	{
		if (text[i] == 'ɿ')
			ret += '⅂';
		else if (text[i] == 'ȷ')
			ret += 'J';
		else
			ret += text[i].toUpperCase();
	}
	return ret;
}

function toLowerCase(text)
{
	var ret = '';
	for (var i = 0; i < text.length; ++i)
	{
		if (text.substring(i, i + 1) == '⅂')
		{
			ret += 'ɿ';
		}
		else if (text[i] == 'J' && text[i + 1] == '\u0308')
			ret += 'ȷ';
		else
			ret += text[i].toLowerCase();
	}
	return ret;
}

String.prototype.toProperCase = function()
{
	if (this.length == 0) {
		return this;
	}
	var ret = '';
	if (this[0] == 'ȷ')
		ret += 'J';
	else if (this[0] == 'ɿ')
		ret += '⅂';
	else
		ret += this[0].toUpperCase();
	for (var i = 1; i < this.length; ++i)
	{
		if (this[i-1] == ' ')
		{
			if (this[i] == 'ɿ')
				ret += '⅂';
			else if (this[i] == 'ȷ')
				ret += 'J';
			else
				ret += this[i].toUpperCase();
		}
		else
		{
			if (this.substring(i, i + 1) == '⅂')
			{
				ret += 'ɿ';
			}
			else if (this[i] == 'J' && this[i + 1] == '\u0308')
				ret += 'ȷ';
			else
				ret += this[i].toLowerCase();
		}
	}
	return ret;
}

function convertCaseWord(text, mode) {
	switch (mode) {
	case 0:
		return toLowerCase(text);
	case 1:
		return toUpperCase(text);
	case 2:
		text = text.toProperCase();
		return text;
	default:
		return text;
	}
}

/******************* END OF CASE CONVERSION UTILITY from 葉劍飛 ********************/
