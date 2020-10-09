package html

type IndexHTMLData struct {
	NSQLookupdAddress string
	Topics []string
}

var IndexHTML = `
	<html>
		<head>
			<title>NSQ Tail Topics via {{ .NSQLookupdAddress }}</title>
		</head>
		<body>
			<h1>Topics</h1>
			<ul>
				{{range .Topics}}
					<li><a href="/nsqtail/{{.}}">{{.}}</a></li>
				{{end}}
			</ul>
		</body>
	</html>
`
