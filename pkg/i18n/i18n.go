package i18n

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func SetupI18n() {
	languages := []string{"en", "tr"}

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for _, lang := range languages {
		files, err := ioutil.ReadDir("./pkg/i18n/locale/" + lang)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if _, err := bundle.LoadMessageFile("./pkg/i18n/locale/" + lang + "/" + file.Name()); err != nil {
				panic(err)
			}
		}
	}

	localizer := i18n.NewLocalizer(bundle, language.English.String())
	welcome, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Welcome",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(welcome)
}

func GetLocalizer(lang string) *i18n.Localizer {
	localizer := i18n.NewLocalizer(bundle, lang)
	return localizer
}

func Translate(lang string, messageID string, args ...interface{}) string {
	localizer := GetLocalizer(lang)
	translated, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: args,
	})

	if err != nil {
		fmt.Println(err)
	}

	return translated
}
