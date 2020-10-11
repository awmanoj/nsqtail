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
				  body {
					width: 60%;
				  }
				
				  h1 {
					font-family: "Geneva";
					background-color: #C0C4E4;
				  }
				
				
				  p {
					font-family: "Geneva";
				  } 
				
				  a {
					font-family: "Geneva"
				  }
				
				  #footer {
					background-color: #DFE1F1;
				  }
			</style>
		</head>
		<body>
			<h1>Last {{ .MessageCount }} messages for topic {{ .Topic }}</h1>
			<ul>
				{{range .Messages}}
					<li>{{.}}</li>
				{{end}}
			</ul>
		</body>
	</html>
`