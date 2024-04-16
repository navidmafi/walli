package validation

import (
	"navidmafi/walli/backends"
	"navidmafi/walli/logger"
	"navidmafi/walli/providers"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	enTrans  ut.Translator
	Validate *validator.Validate
)

type ConfigSchema struct {
	Backend  string `validate:"validBackend"`
	Provider string `validate:"validProvider"`
}

type APIKeys struct {
	Unsplash string
}

func validateBackendName(fl validator.FieldLevel) bool {
	backend := backends.BackendName(fl.Field().String())

	for key := range backends.Backends {
		if backend == key {
			return true
		}
	}

	return false

}

func validateProviderName(fl validator.FieldLevel) bool {
	provider := providers.ProviderName(fl.Field().String())

	for key := range providers.Providers {
		if provider == key {
			return true
		}
	}

	return false

}

func Init() {
	en := en.New()
	uni = ut.New(en, en)

	enTrans, _ = uni.GetTranslator("en")
	Validate = validator.New()
	en_translations.RegisterDefaultTranslations(Validate, enTrans)
	Validate.RegisterValidation("validBackend", validateBackendName)
	Validate.RegisterValidation("validProvider", validateProviderName)
}

func ValidateKeyValueOrPanic(s interface{}, key string, value string) {

	configType := reflect.TypeOf(s)

	keyFound := false
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		validateTag := field.Tag.Get("validate")

		err := Validate.Var(value, validateTag)

		if field.Name == key {
			keyFound = true

			if err != nil {
				validationErrors := err.(validator.ValidationErrors)

				for _, ve := range validationErrors {
					logger.Logger.Fatalf("%s%s\n", key, ve.Translate(enTrans))
				}
			}
		}

	}

	if !keyFound {
		panic("Inavlid Key supplied")
	}
}
