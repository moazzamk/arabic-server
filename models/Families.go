package models

type Family struct {
	Name string
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

var BigFamilies = []Family{
	// aslama - af3ala
	{
		Name: "\u0623\u064E\u0633\u0652\u0644\u064E\u0645\u064E",
		ActivePast: "\u0623\u064E%s\u0652%s\u064E%s",
		ActivePresent: "\u064F%s\u0652%s\u064E%s",
		Idea: "\u0625%s\u0652%s\u0627%s\u064C",
		PassivePast: "\u0623\u064F%s\u0652%s\u0650%s\u064E",
		PassivePresent: "\u064F%s\u0652%s\u064E%s",
	},
	// 3allama
	{
		Name: "\u0639\u064E\u0644\u0651\u064E\u0645\u064E",
		ActivePast: "%s\u064E%s\u0651\u064E%s",
		ActivePresent: "\u064F%s\u064E%s\u0651\u0650%s",
	},
	// jaahada
	{
		Name: "\u062C\u064E\u0627\u0647\u064E\u062F\u064E",
		ActivePast: "%s\u064E\u0627%s\u064E%s",
		ActivePresent: "\u064F%s\u064E\u0627%s\u0650%s",
	},

	// ta3allama - ta-fa3-3ala
	{
		Name: "\u062A\u064E\u0639\u064E\u0644\u0651\u064E\u0645\u064E",
		ActivePast: "\u062A\u064E%s\u064E%s\u0651\u064E%s",
		ActivePresent: "\u064E\u062A\u064E%s\u064E%s\u0651\u064E%s",
	},
	// tasaa ala - tafaa3ala
	{
		Name: "تَسَاءَلَ",
		ActivePast: "\u062A\u064E%s\u064E\u0627%s\u064E%s",
		ActivePresent: "\u064E\u062A\u064E%s\u064E\u0627%s\u064E%s",
	},
	// iqtaraba - ifta3ala
	{
		Name: "إقترب",
		ActivePast: "\u0625%s\u0652\u062A\u064E%s\u064E%s",
		ActivePresent: "\u064E%s\u0652\u062A\u064E%s\u0650\u0644%s",
	},
	// inqalaba - infa3ala
	{
		Name: "إنقعل",
		ActivePast: "\u0625\u0646\u0652%s\u064E%s\u064E%s",
		ActivePresent: "\u064E\u0646\u0652%s\u064E%s\u0650%s",
	},
	//istaghfara - istaf3ala
	{
		Name: "إسْتَغْفَرَِ",
		ActivePast: "\u0625\u0633\u0652\u062A\u064E%s\u0652%s\u064E%s",
		ActivePresent: "\u064E\u0633\u0652\u062A\u064E%s\u0652%s\u064E%s",
	},
}
