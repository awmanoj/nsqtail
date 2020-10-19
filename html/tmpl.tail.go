package html

type TailHTMLData struct {
	NSQLookupdAddress string
	Topic string
	MessageCount int
	Messages []string
}

var TailHTML = `
	<html>
		<head>
			<title>NSQ Tail for {{.Topic}} via {{ .NSQLookupdAddress }}</title>
			<style>
			</style>
		</head>
		<body>
			<h1>Last {{ .MessageCount }} messages for topic {{ .Topic }}</h1>
			<ul>
				{{range .Messages}}
					{{.}}
					<hr/>
				{{end}}
			</ul>
		</body>
	</html>
`