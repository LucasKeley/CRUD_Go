package docs

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed swagger.json swagger.yaml
var swaggerFS embed.FS

// RegisterRoutes adiciona as rotas de documentação Swagger no grupo informado.
func RegisterRoutes(r *gin.RouterGroup) {
	docsGroup := r.Group("/docs")
	{
		docsGroup.GET("/swagger.json", serveFile("swagger.json"))
		docsGroup.GET("/swagger.yaml", serveFile("swagger.yaml"))
		docsGroup.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "docs-index", gin.H{})
		})
	}

	r.Engine.SetHTMLTemplate(docTemplate)
}

func serveFile(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := swaggerFS.ReadFile(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		contentType := "application/json"
		if name == "swagger.yaml" {
			contentType = "application/x-yaml"
		}

		c.Data(http.StatusOK, contentType, data)
	}
}

var docTemplate = templateMust(`{{ define "docs-index" }}
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8" />
    <title>CRUD_Go Swagger</title>
    <style>
        body { font-family: sans-serif; margin: 40px; }
        h1 { color: #333; }
        a { color: #0a5; text-decoration: none; font-weight: bold; }
        a:hover { text-decoration: underline; }
        ul { margin-top: 24px; }
    </style>
</head>
<body>
    <h1>Documentação Swagger</h1>
    <p>Escolha um formato para visualizar a especificação:</p>
    <ul>
        <li><a href="/docs/swagger.json">swagger.json</a></li>
        <li><a href="/docs/swagger.yaml">swagger.yaml</a></li>
    </ul>
    <p>Você pode importar qualquer um dos arquivos em ferramentas como <strong>Swagger UI</strong>, <strong>Insomnia</strong> ou <strong>Postman</strong>.</p>
</body>
</html>
{{ end }}`)

func templateMust(html string) *template.Template {
	tmpl, err := template.New("docs-index").Parse(html)
	if err != nil {
		panic(err)
	}
	return tmpl
}
