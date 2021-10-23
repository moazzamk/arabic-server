package models

type Pronoun struct {
	Pattern string
	Pronoun string
	Tense   string
}

type Tense string

const (
	Present Tense = "present"
	Past          = "past"
)

var PastPronouns = []Pronoun{
	{
		Pattern: "%s\u064E%s\u064E%s\u064E",
		Pronoun: "He",
		Tense: "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u064E\u0627",
		Pronoun: "Both of them",
		Tense:   "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u0648\u0627",
		Pronoun: "All of them",
		Tense:   "past",
	},
	// 3f
	{
		Pattern: "%s\u064E%s\u064E%s\u064E\u062A\u0652",
		Pronoun: "She",
		Tense:   "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u064E\u062A\u0627",
		Pronoun: "Both of them ladies",
		Tense:   "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u0646\u0618",
		Pronoun: "All of them ladies",
		Tense:   "past",
	},
	// 2m
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u062A\u0618",
		Pronoun: "You",
		Tense:   "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u062A\u064F\u0645\u0627",
		Pronoun: "Both of you",
		Tense:   "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u062A\u064F\u0645\u0652",
		Pronoun: "All of you",
		Tense:   "past",
	},
	// 2f
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u062A\u0650",
		Pronoun: "You (fem.)",
		Tense:   "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u062A\u064F\u0646\u0651\u064E",
		Pronoun: "All of you ladies",
		Tense:   "past",
	},
	// 1p
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u062A\u064F",
		Pronoun: "I",
		Tense:   "past",
	},
	{
		Pattern: "%s\u064E%s\u064E%s\u0652\u0646\u0627",
		Pronoun: "We",
		Tense:   "past",
	},
}

