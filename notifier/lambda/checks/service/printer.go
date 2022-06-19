package service

import (
	"html/template"
	"strings"

	"github.com/yunomu/checks-codepipeline/lambda/checks/pipelinestat"
)

var defaultTemplate = template.Must(template.New("template").Parse(`{{range .StageStats}}<h3>Stage: {{.Name}}</h3>{{range .ActionStats}}<h4>Action: {{.Name}}</h4><ul><li>ActionExecutionId: {{.ExecutionId}}</li><li>Status: {{.Status}}</li><li>Message: {{.ErrorDetail}}</li></ul>{{end}}{{end}}`))

type DetailPrinter struct {
	template *template.Template
}

func NewDetailPrinter() *DetailPrinter {
	return &DetailPrinter{
		template: defaultTemplate,
	}
}

func (p *DetailPrinter) Print(stat *pipelinestat.Stat) (string, error) {
	var buf strings.Builder
	if err := p.template.Execute(&buf, stat); err != nil {
		return "", err
	}

	return buf.String(), nil
}
