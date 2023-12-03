package email

import (
	"bytes"
	"text/template"
)

func ConstuctMessageHtml(strHtml string, data interface{}) (string, error) {

	t, err := template.New("strHtml").Parse(strHtml)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func BodySendEmailVerification(url, body string) string {
	mapVariable := make(map[string]string)
	mapVariable["url"] = url

	tmp, err := ConstuctMessageHtml(body, mapVariable)
	if err != nil {
		return err.Error()
	}
	return tmp
}
