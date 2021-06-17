package ghtrending

var (
	spokenLangCode map[string]string
)

func init() {
	spokenLangCode = map[string]string{
		"abkhazian":             "ab",
		"afar":                  "aa",
		"afrikaans":             "af",
		"akan":                  "ak",
		"albanian":              "sq",
		"amharic":               "am",
		"arabic":                "ar",
		"aragonese":             "an",
		"armenian":              "hy",
		"assamese":              "as",
		"avaric":                "av",
		"avestan":               "ae",
		"aymara":                "ay",
		"azerbaijani":           "az",
		"bambara":               "bm",
		"bashkir":               "ba",
		"basque":                "eu",
		"belarusian":            "be",
		"bengali":               "bn",
		"bihari languages":      "bh",
		"bislama":               "bi",
		"bosnian":               "bs",
		"breton":                "br",
		"bulgarian":             "bg",
		"burmese":               "my",
		"catalan":               "ca",
		"valencian":             "ca",
		"chamorro":              "ch",
		"chechen":               "ce",
		"chichewa":              "ny",
		"chewa":                 "ny",
		"nyanja":                "ny",
		"chinese":               "zh",
		"chuvash":               "cv",
		"cornish":               "kw",
		"corsican":              "co",
		"cree":                  "cr",
		"croatian":              "hr",
		"czech":                 "cs",
		"danish":                "da",
		"divehi":                "dv",
		"dhivehi":               "dv",
		"maldivian":             "dv",
		"dutch":                 "nl",
		"flemish":               "nl",
		"dzongkha":              "dz",
		"english":               "en",
		"esperanto":             "eo",
		"estonian":              "et",
		"ewe":                   "ee",
		"faroese":               "fo",
		"fijian":                "fj",
		"finnish":               "fi",
		"french":                "fr",
		"fulah":                 "ff",
		"galician":              "gl",
		"georgian":              "ka",
		"german":                "de",
		"greek":                 "el",
		"modern":                "el",
		"guarani":               "gn",
		"gujarati":              "gu",
		"haitian":               "ht",
		"haitian creole":        "ht",
		"hausa":                 "ha",
		"hebrew":                "he",
		"herero":                "hz",
		"hindi":                 "hi",
		"hiri motu":             "ho",
		"hungarian":             "hu",
		"interlingua":           "ia",
		"indonesian":            "id",
		"interlingue":           "ie",
		"occidental":            "ie",
		"irish":                 "ga",
		"igbo":                  "ig",
		"inupiaq":               "ik",
		"ido":                   "io",
		"icelandic":             "is",
		"italian":               "it",
		"inuktitut":             "iu",
		"japanese":              "ja",
		"javanese":              "jv",
		"kalaallisut":           "kl",
		"greenlandic":           "kl",
		"kannada":               "kn",
		"kanuri":                "kr",
		"kashmiri":              "ks",
		"kazakh":                "kk",
		"central khmer":         "km",
		"kikuyu":                "ki",
		"gikuyu":                "ki",
		"kinyarwanda":           "rw",
		"kirghiz":               "ky",
		"kyrgyz":                "ky",
		"komi":                  "kv",
		"kongo":                 "kg",
		"korean":                "ko",
		"kurdish":               "ku",
		"kuanyama":              "kj",
		"kwanyama":              "kj",
		"latin":                 "la",
		"luxembourgish":         "lb",
		"letzeburgesch":         "lb",
		"ganda":                 "lg",
		"limburgan":             "li",
		"limburger":             "li",
		"limburgish":            "li",
		"lingala":               "ln",
		"lao":                   "lo",
		"lithuanian":            "lt",
		"luba-katanga":          "lu",
		"latvian":               "lv",
		"manx":                  "gv",
		"macedonian":            "mk",
		"malagasy":              "mg",
		"malay":                 "ms",
		"malayalam":             "ml",
		"maltese":               "mt",
		"maori":                 "mi",
		"marathi":               "mr",
		"marshallese":           "mh",
		"mongolian":             "mn",
		"nauru":                 "na",
		"navajo":                "nv",
		"navaho":                "nv",
		"north ndebele":         "nd",
		"nepali":                "ne",
		"ndonga":                "ng",
		"norwegian bokmål":      "nb",
		"norwegian nynorsk":     "nn",
		"norwegian":             "no",
		"sichuan yi":            "ii",
		"nuosu":                 "ii",
		"south ndebele":         "nr",
		"occitan":               "oc",
		"ojibwa":                "oj",
		"church slavic":         "cu",
		"old slavonic":          "cu",
		"chu...":                "cu",
		"oromo":                 "om",
		"oriya":                 "or",
		"ossetian":              "os",
		"ossetic":               "os",
		"punjabi":               "pa",
		"panjabi":               "pa",
		"pali":                  "pi",
		"persian":               "fa",
		"polish":                "pl",
		"pashto":                "ps",
		"pushto":                "ps",
		"portuguese":            "pt",
		"quechua":               "qu",
		"romansh":               "rm",
		"rundi":                 "rn",
		"romanian":              "ro",
		"moldavian":             "ro",
		"moldovan":              "ro",
		"russian":               "ru",
		"sanskrit":              "sa",
		"sardinian":             "sc",
		"sindhi":                "sd",
		"northern sami":         "se",
		"samoan":                "sm",
		"sango":                 "sg",
		"serbian":               "sr",
		"gaelic":                "gd",
		"scottish gaelic":       "gd",
		"shona":                 "sn",
		"sinhala":               "si",
		"sinhalese":             "si",
		"slovak":                "sk",
		"slovenian":             "sl",
		"somali":                "so",
		"southern sotho":        "st",
		"spanish":               "es",
		"castilian":             "es",
		"sundanese":             "su",
		"swahili":               "sw",
		"swati":                 "ss",
		"swedish":               "sv",
		"tamil":                 "ta",
		"telugu":                "te",
		"tajik":                 "tg",
		"thai":                  "th",
		"tigrinya":              "ti",
		"tibetan":               "bo",
		"turkmen":               "tk",
		"tagalog":               "tl",
		"tswana":                "tn",
		"tonga (tonga islands)": "to",
		"turkish":               "tr",
		"tsonga":                "ts",
		"tatar":                 "tt",
		"twi":                   "tw",
		"tahitian":              "ty",
		"uighur":                "ug",
		"uyghur":                "ug",
		"ukrainian":             "uk",
		"urdu":                  "ur",
		"uzbek":                 "uz",
		"venda":                 "ve",
		"vietnamese":            "vi",
		"volapük":               "vo",
		"walloon":               "wa",
		"welsh":                 "cy",
		"wolof":                 "wo",
		"western frisian":       "fy",
		"xhosa":                 "xh",
		"yiddish":               "yi",
		"yoruba":                "yo",
		"zhuang":                "za",
		"chuang":                "za",
		"zulu":                  "zu",
	}
}

type options struct {
	GitHubURL  string
	SpokenLang string
	Language   string // programming language
	DateRange  string
}

type option func(*options)

func WithDaily() option {
	return func(opt *options) {
		opt.DateRange = "daily"
	}
}

func WithWeekly() option {
	return func(opt *options) {
		opt.DateRange = "weekly"
	}
}

func WithMonthly() option {
	return func(opt *options) {
		opt.DateRange = "monthly"
	}
}

func WithDateRange(dr string) option {
	return func(opt *options) {
		opt.DateRange = dr
	}
}

func WithLanguage(lang string) option {
	return func(opt *options) {
		opt.Language = lang
	}
}

func WithSpokenLanguageCode(code string) option {
	return func(opt *options) {
		opt.SpokenLang = code
	}
}

func WithSpokenLanguageFull(lang string) option {
	return func(opt *options) {
		opt.SpokenLang = spokenLangCode[lang]
	}
}

func WithURL(url string) option {
	return func(opt *options) {
		opt.GitHubURL = url
	}
}
