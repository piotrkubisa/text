package catalog_test

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func ExampleMessage() {
	message.Set(language.English, "You are %d minute(s) late.",
		catalog.Var("minutes",
			plural.Selectf(1, "",
				plural.One, "minute",
				plural.Other, "minutes",
			),
		),
		catalog.String("You are %[1]d ${minutes} late."),
	)
	p := message.NewPrinter(language.English)
	p.Printf("You are %d minute(s) late.", 5)
	//Output:
	// You are 5 minutes late.
}
