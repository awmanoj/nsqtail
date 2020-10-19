package html

type IndexHTMLData struct {
	NSQLookupdAddress string
	Topics []string
}

var IndexHTML = `
	<html>
		<head>
			<title>NSQ Tail Topics via {{ .NSQLookupdAddress }}</title>
			<style>
			</style>
		</head>
		<body>
			<h1>Topics</h1>
			<ul>
				{{range .Topics}}
					<a href="/nsqtail/{{.}}">{{.}}</a>
				{{end}}
			</ul>
		</body>
	</html>
`
