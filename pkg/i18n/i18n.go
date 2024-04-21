package i18n

import (
	"github.com/juju/errors"
	"github.com/k8scat/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

var (
	bundle *i18n.Bundle
)

func InitI18N(localesPath string) error {
	files, err := os.ReadDir(localesPath)
	if err != nil {
		return err
	}
	bundle = i18n.NewBundle(language.English)
	for _, file := range files {
		localeFile := filepath.Join(localesPath, file.Name())
		if _, err := bundle.LoadMessageFile(localeFile); err != nil {
			continue
		}
	}
	return nil
}

func T(tags []language.Tag, key string) string {
	return Translate(tags, key, nil)
}
func Translate(tags []language.Tag, key string, data interface{}) string {
	return translate(bundle, tags, key, data)
}

func translate(bundle *i18n.Bundle, tags []language.Tag, key string, data interface{}) string {
	localizer := i18n.NewLocalizerWithTags(bundle, tags...)
	text, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    ParseKey(key),
		TemplateData: data,
	})
	if err != nil {
		log.Printf("i18n translate failed: %+v, trace: %s", errors.Trace(err), debug.Stack())
	}
	return text
}

func ParseKey(s string) string {
	s = strings.TrimPrefix(s, "{{")
	s = strings.TrimSuffix(s, "}}")
	s = strings.TrimSpace(s)
	return s
}
