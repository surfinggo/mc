package mc

import (
	"github.com/pkg/errors"
	"io"
	"os"
)

type OpenOrCreateOptions struct {
	Template     string
	TemplatePath string
}

type OpenOrCreateOptionFunc func(o *OpenOrCreateOptions) error

func WithTemplate(template string) OpenOrCreateOptionFunc {
	return func(o *OpenOrCreateOptions) error {
		o.Template = template
		return nil
	}
}

func WithTemplatePath(templatePath string) OpenOrCreateOptionFunc {
	return func(o *OpenOrCreateOptions) error {
		o.TemplatePath = templatePath
		return nil
	}
}

// EnsureFileExists ensures the file exist, creates it with a template if not.
func EnsureFileExists(filePath string, optionFuncs ...OpenOrCreateOptionFunc) error {
	options := &OpenOrCreateOptions{}
	for _, optionFunc := range optionFuncs {
		if err := optionFunc(options); err != nil {
			return errors.Wrap(err, "apply option error")
		}
	}

	_, err := os.Stat(filePath)
	if err == nil {
		return nil // already exist
	}
	if !os.IsNotExist(err) {
		return errors.Wrapf(err, "os.Stat: %s error", filePath)
	}

	// not exist, create it
	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "os.Open: %s error", filePath)
	}
	defer func() {
		Must(file.Close())
	}()

	if options.Template != "" {
		if _, err := file.WriteString(options.Template); err != nil {
			return errors.Wrap(err, "file.WriteString error")
		}
		return nil
	}

	if options.TemplatePath != "" {
		if _, err := os.Stat(options.TemplatePath); err != nil {
			if os.IsNotExist(err) {
				return errors.Errorf("template not exist: %s", options.TemplatePath)
			}
			return errors.Wrapf(err, "os.Stat: %s error", options.TemplatePath)
		}

		tmpl, err := os.Open(options.TemplatePath)
		if err != nil {
			return errors.Wrapf(err, "os.Open: %s error", options.TemplatePath)
		}
		defer func() {
			Must(tmpl.Close())
		}()

		_, err = io.Copy(file, tmpl)
		if err != nil {
			return errors.Wrap(err, "io.Copy error")
		}
	}

	// no template, just create an empty file
	return nil
}

// OpenOrCreate opens the file, creates it with a template if not.
func OpenOrCreate(filePath string, optionFuncs ...OpenOrCreateOptionFunc) (*os.File, error) {
	if err := EnsureFileExists(filePath, optionFuncs...); err != nil {
		return nil, err
	}
	return os.Open(filePath)

}
