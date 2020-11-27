package models

var PresentPronouns []Pronoun

func init()  {
	PresentPronouns = []Pronoun{
		{
			Pattern: "\u064A\u064E%s\u0652%s\u064F%s\u064F",
			Pronoun: "He",
			Tense:   "present",
		},
		{
			Pattern: "\u064A\u064E%s\u0652%s\u064F%s\u064E\u0627\u0646\u0650",
			Pronoun: "Both of them",
			Tense:   "present",
		},
		{
			Pattern: "\u064A\u064E%s\u0652%s\u064F%s\u064F\u0648\u0646\u064E",
			Pronoun: "All of them",
			Tense:   "present",
		},

		// 3f
		{
			Pattern: "\u062A\u064E%s\u0652%s\u064F%s\u064F",
			Pronoun: "She / You",
			Tense:   "present",
		},
		{
			Pattern: "\u062A\u064E%s\u0652%s\u064F%s\u064E\u0627\u0646\u0650",
			Pronoun: "Both of you / Both of them ladies",
			Tense:   "present",
		},
		{
			Pattern: "\u064A\u064E%s\u0652%s\u064F%s\u0652\u0646\u064E",
			Pronoun: "Them ladies",
			Tense:   "present",
		},

		// 2m
		{
			Pattern: "\u062A\u064E%s\u0652%s\u064F%s\u064F\u0648\u0646\u064E",
			Pronoun: "All of you",
			Tense:   "present",
		},

		// 2f
		{
			Pattern: "\u062A\u064E%s\u0652%s\u064F%s\u0650\u064A\u0646\u064E",
			Pronoun: "You (fem)",
			Tense:   "present",
		},
		{
			Pattern: "\u062A\u064E%s\u0652%s\u064F%s\u0652\u0646\u064E",
			Pronoun: "You ladies",
			Tense:   "present",
		},

		// 1
		{
			Pattern: "\u0627\u064E%s\u0652%s\u064F%s\u064F",
			Pronoun: "I",
			Tense:   "present",
		},
		{
			Pattern: "\u0646\u064E%s\u0652%s\u064F%s\u064F",
			Pronoun: "We",
			Tense:   "present",
		},
	}
}
