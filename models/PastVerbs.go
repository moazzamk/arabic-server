package models

type Verb struct {
	Beginning string
	Ending string
	Tense string
	Doer string

}

var PastVerbs = []Verb{
	// 3m
	{ Ending: "\u064E", Doer: "He", Tense: "past"},
	{ Ending: "\u064E\u0627", Doer: "Both of them", Tense: "past"},
	{ Ending: "\u0648\u064F\u0627", Doer: "All of them", Tense: "past"},

	//3f
	{ Ending: "\u064E\u062A\u0652", Doer: "She", Tense: "past"},
	{ Ending: "\u064E\u062A\u064E\u0627",
		Doer: "Both of them ladies",
		Tense: "past",
	},
	{
		Ending: "\u0652\u0646\u0618",
		Doer: "All of them ladies",
		Tense:   "past",
	},
	// 2m
	{
		Ending: "\u0652\u062A\u0618",
		Doer: "You",
		Tense:   "past",
	},
	{
		Ending: "\u0652\u062A\u064F\u0645\u0627",
		Doer: "Both of you",
		Tense:   "past",
	},
	{
		Ending: "\u0652\u062A\u064F\u0645\u0652",
		Doer: "All of you",
		Tense:   "past",
	},
	// 2f
	{
		Ending: "\u0652\u062A\u0650",
		Doer: "You (fem.)",
		Tense:   "past",
	},
	{
		Ending: "\u0652\u062A\u064F\u0646\u0651\u064E",
		Doer: "All of you ladies",
		Tense:   "past",
	},
	// 1p
	{
		Ending: "\u0652\u062A\u064F",
		Doer: "I",
		Tense:   "past",
	},
	{
		Ending: "\u0652\u0646\u0627",
		Doer: "We",
		Tense:   "past",
	},
}
