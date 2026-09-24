package docs

import (
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

//go:embed swagger.yaml
var swaggerYAML []byte

// specJSON holds the spec pre-converted to JSON at startup so it can be
// inlined into the HTML page — no separate HTTP endpoint is exposed.
var specJSON []byte

func init() {
	var obj any
	if err := yaml.Unmarshal(swaggerYAML, &obj); err != nil {
		panic("docs: failed to parse swagger.yaml: " + err.Error())
	}
	b, err := json.Marshal(obj)
	if err != nil {
		panic("docs: failed to marshal swagger spec to JSON: " + err.Error())
	}
	specJSON = b
}

// SwaggerUIHandler serves the Swagger UI with the spec baked into the page.
func SwaggerUIHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, buildSwaggerHTML(specJSON))
}
