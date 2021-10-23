package models

var PresentVerbs = []Verb{
	{
		Beginning: "\u064A",
		Ending: "\u064F",
		Doer: "He",
		Tense:   "present",
	},
	{
		Beginning: "\u064A",
		Ending: "\u064E\u0627\u0646\u0650",
		Doer: "Both of them",
		Tense:   "present",
	},
	{
		Beginning: "\u064A",
		Ending: "\u064F\u0648\u0646\u064E",
		Doer: "All of them",
		Tense:   "present",
	},

	// 3f
	{
		Beginning: "\u062A",
		Ending: "\u064F",
		Doer: "She / You",
		Tense:   "present",
	},
	{
		Beginning:"\u062A",
		Ending: "\u064E\u0627\u0646\u0650",
		Doer: "Both of you / Both of them ladies",
		Tense:   "present",
	},
	{
		Beginning: "\u064A",
		Ending: "\u0652\u0646\u064E",
		Doer: "Them ladies",
		Tense:   "present",
	},

	// 2m
	{
		Beginning: "\u062A",
		Ending: "\u064F\u0648\u0646\u064E",
		Doer: "All of you",
		Tense:   "present",
	},

	// 2f
	{
		Beginning: "\u062A",
		Ending: "\u0650\u064A\u0646\u064E",
		Doer: "You (fem)",
		Tense:   "present",
	},
	{
		Beginning: "\u062A",
		Ending: "\u0652\u0646\u064E",
		Doer: "You ladies",
		Tense:   "present",
	},

	// 1
	{
		Beginning: "\u0627",
		Ending: "\u064F",
		Doer: "I",
		Tense:   "present",
	},
	{
		Beginning: "\u0646",
		Ending: "\u064F",
		Doer: "We",
		Tense:   "present",
	},
}
