package docs

import "fmt"

func buildSwaggerHTML(specJSON []byte) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Portfolio Go – API Reference</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

    :root {
      --bg:          #0d1117;
      --surface:     #161b22;
      --border:      #30363d;
      --text:        #e6edf3;
      --text-muted:  #8b949e;
      --accent:      #10a37f;
      --accent-soft: #0d8f70;
      --get:         #238636;
      --get-bg:      #0d2a1a;
      --post:        #1f6feb;
      --post-bg:     #0d1f3a;
      --delete:      #da3633;
      --delete-bg:   #2d1215;
      --put:         #9e6a03;
      --put-bg:      #2a1f00;
      --patch:       #6e40c9;
      --patch-bg:    #1e1040;
    }

    html, body { height: 100%%; background: var(--bg); color: var(--text); font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif; }

    .navbar {
      position: sticky; top: 0; z-index: 100;
      display: flex; align-items: center; gap: 12px;
      padding: 0 24px; height: 56px;
      background: var(--surface); border-bottom: 1px solid var(--border);
    }
    .navbar-logo { display: flex; align-items: center; gap: 8px; font-weight: 700; font-size: 16px; color: var(--text); text-decoration: none; }
    .navbar-logo svg { width: 28px; height: 28px; fill: var(--accent); }
    .navbar-badge { font-size: 11px; font-weight: 600; padding: 2px 8px; border-radius: 999px; background: var(--accent); color: #fff; }
    .navbar-spacer { flex: 1; }
    .navbar-version { font-size: 12px; color: var(--text-muted); }

    #swagger-ui { max-width: 1080px; margin: 0 auto; padding: 32px 16px 80px; }

    .swagger-ui .scheme-container { background: var(--surface) !important; border: 1px solid var(--border) !important; border-radius: 8px !important; padding: 16px 20px !important; margin-bottom: 24px !important; box-shadow: none !important; }
    .swagger-ui .info { margin-bottom: 28px !important; }
    .swagger-ui .info .title { color: var(--text) !important; font-size: 28px !important; font-weight: 700 !important; }
    .swagger-ui .info p, .swagger-ui .info li, .swagger-ui .info .description { color: var(--text-muted) !important; font-size: 14px !important; line-height: 1.6 !important; }
    .swagger-ui .info a { color: var(--accent) !important; }

    .swagger-ui .opblock { border-radius: 8px !important; border: 1px solid var(--border) !important; margin-bottom: 12px !important; overflow: hidden !important; }
    .swagger-ui .opblock-summary { padding: 12px 16px !important; background: transparent !important; cursor: pointer; }
    .swagger-ui .opblock-summary:hover { filter: brightness(1.06); }
    .swagger-ui .opblock-summary-method { border-radius: 4px !important; font-size: 12px !important; font-weight: 700 !important; min-width: 62px !important; text-align: center !important; padding: 4px 8px !important; }
    .swagger-ui .opblock-summary-path { color: var(--text) !important; font-size: 14px !important; font-weight: 600 !important; }
    .swagger-ui .opblock-summary-description { color: var(--text-muted) !important; font-size: 13px !important; }

    .swagger-ui .opblock-get    { background: var(--get-bg)    !important; border-color: var(--get)    !important; }
    .swagger-ui .opblock-get    .opblock-summary-method { background: var(--get)    !important; }
    .swagger-ui .opblock-post   { background: var(--post-bg)   !important; border-color: var(--post)   !important; }
    .swagger-ui .opblock-post   .opblock-summary-method { background: var(--post)   !important; }
    .swagger-ui .opblock-delete { background: var(--delete-bg) !important; border-color: var(--delete) !important; }
    .swagger-ui .opblock-delete .opblock-summary-method { background: var(--delete) !important; }
    .swagger-ui .opblock-put    { background: var(--put-bg)    !important; border-color: var(--put)    !important; }
    .swagger-ui .opblock-put    .opblock-summary-method { background: var(--put)    !important; }
    .swagger-ui .opblock-patch  { background: var(--patch-bg)  !important; border-color: var(--patch)  !important; }
    .swagger-ui .opblock-patch  .opblock-summary-method { background: var(--patch)  !important; }

    .swagger-ui .opblock-body { background: var(--bg) !important; border-top: 1px solid var(--border) !important; padding: 16px !important; }
    .swagger-ui .opblock-section-header { background: var(--surface) !important; border-radius: 4px !important; padding: 8px 12px !important; }
    .swagger-ui .opblock-section-header label, .swagger-ui .opblock-section-header h4 { color: var(--text) !important; font-size: 13px !important; }

    .swagger-ui table { background: transparent !important; }
    .swagger-ui table thead tr td, .swagger-ui table thead tr th { background: var(--surface) !important; color: var(--text-muted) !important; border-bottom: 1px solid var(--border) !important; font-size: 12px !important; }
    .swagger-ui table tbody tr td { background: transparent !important; color: var(--text) !important; border-bottom: 1px solid var(--border) !important; font-size: 13px !important; }
    .swagger-ui .parameter__name { color: var(--text) !important; font-weight: 600 !important; }
    .swagger-ui .parameter__type { color: var(--accent) !important; font-size: 12px !important; }
    .swagger-ui .parameter__in   { color: var(--text-muted) !important; font-size: 11px !important; }
    .swagger-ui .required-label  { color: var(--delete) !important; }
    .swagger-ui .prop-type        { color: var(--accent) !important; }

    .swagger-ui input[type=text], .swagger-ui input[type=number], .swagger-ui textarea, .swagger-ui select {
      background: var(--bg) !important; border: 1px solid var(--border) !important; border-radius: 6px !important;
      color: var(--text) !important; font-size: 13px !important; padding: 6px 10px !important;
    }
    .swagger-ui input[type=text]:focus, .swagger-ui input[type=number]:focus, .swagger-ui textarea:focus { border-color: var(--accent) !important; outline: none !important; }

    .swagger-ui .btn { border-radius: 6px !important; font-size: 13px !important; font-weight: 600 !important; padding: 6px 14px !important; cursor: pointer !important; }
    .swagger-ui .btn.execute { background: var(--accent) !important; border-color: var(--accent) !important; color: #fff !important; }
    .swagger-ui .btn.execute:hover { background: var(--accent-soft) !important; }
    .swagger-ui .btn.try-out__btn, .swagger-ui .btn.cancel { background: transparent !important; border-color: var(--border) !important; color: var(--text) !important; }
    .swagger-ui .btn.authorize { background: var(--surface) !important; border-color: var(--accent) !important; color: var(--accent) !important; }

    .swagger-ui .responses-inner { background: var(--bg) !important; }
    .swagger-ui .response-col_status { color: var(--text) !important; font-weight: 700 !important; }
    .swagger-ui .response-col_description__inner p { color: var(--text-muted) !important; }
    .swagger-ui .highlight-code pre, .swagger-ui .microlight { background: var(--surface) !important; border: 1px solid var(--border) !important; border-radius: 6px !important; color: var(--text) !important; font-size: 12.5px !important; }

    .swagger-ui section.models { background: var(--surface) !important; border: 1px solid var(--border) !important; border-radius: 8px !important; padding: 16px !important; margin-top: 32px !important; }
    .swagger-ui section.models .model-container { background: var(--bg) !important; border: 1px solid var(--border) !important; border-radius: 6px !important; margin-bottom: 8px !important; }
    .swagger-ui .model-title { color: var(--text) !important; font-weight: 700 !important; }
    .swagger-ui .model       { color: var(--text-muted) !important; font-size: 13px !important; }
    .swagger-ui .prop-name   { color: var(--accent) !important; }

    .swagger-ui .opblock-tag { color: var(--text) !important; border-bottom: 1px solid var(--border) !important; font-size: 18px !important; font-weight: 700 !important; padding-bottom: 8px !important; margin-bottom: 16px !important; }
    .swagger-ui .opblock-tag small { color: var(--text-muted) !important; font-size: 13px !important; font-weight: 400 !important; }

    .swagger-ui .loading-container { background: var(--bg) !important; }
    .swagger-ui .loading-container .loading::after { border-color: var(--accent) transparent transparent !important; }

    .swagger-ui label { color: var(--text-muted) !important; }
    .swagger-ui h4    { color: var(--text) !important; }
    .swagger-ui p     { color: var(--text-muted) !important; }
    .swagger-ui a     { color: var(--accent) !important; }
    .swagger-ui svg   { fill: var(--text-muted) !important; }
    .swagger-ui .topbar { display: none !important; }

    ::-webkit-scrollbar { width: 6px; height: 6px; }
    ::-webkit-scrollbar-track { background: var(--bg); }
    ::-webkit-scrollbar-thumb { background: var(--border); border-radius: 3px; }
  </style>
</head>
<body>

<nav class="navbar">
  <a class="navbar-logo" href="#">
    <svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
      <path d="M16 2 L29 9.5 L29 22.5 L16 30 L3 22.5 L3 9.5 Z"/>
    </svg>
    Portfolio Go
  </a>
  <span class="navbar-badge">API</span>
  <span class="navbar-spacer"></span>
  <span class="navbar-version">OpenAPI 3.0</span>
</nav>

<div id="swagger-ui"></div>

<script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
<script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
<script>
  window.onload = function () {
    SwaggerUIBundle({
      spec: %s,
      dom_id: "#swagger-ui",
      presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
      layout: "StandaloneLayout",
      deepLinking: true,
      defaultModelsExpandDepth: 1,
      defaultModelExpandDepth: 2,
      displayRequestDuration: true,
      filter: true,
      tryItOutEnabled: false,
    });
  };
</script>
</body>
</html>`, specJSON)
}
