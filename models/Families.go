package models

type Family struct {
	ActivePast string
	ActivePresent string
	Idea string
	Doer string
	PassivePast string
	PassivePresent string
	Command string
	Forbid string

	IsmMaf3ool string

}

var BigFamilies []Family

func init()  {
	BigFamilies = []Family {
		// aslama
		{
			ActivePast: "\u0623\u064E%s\u0652%s\u064E%s\u064E",
			ActivePresent: "\u064F%s\u0652%s\u064E%s\u064F",
			Idea: "\u0625%s\u0652%s\u0627%s\u064C",
			PassivePast: "\u0623\u064F%s\u0652%s\u0650%s\u064E",
			PassivePresent: "\u064F%s\u0652%s\u064E%s\u064F",
		},
		// 3allama
		{
			ActivePast: "%s\u064E%s\u0651\u064E%s\u064E",
			ActivePresent: "\u064F%s\u064E%s\u0651\u0650%s\u064F",
		},
		//
		{

		},
	}
}
