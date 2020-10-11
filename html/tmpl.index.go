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
			<h1>Topics</h1>
			<ul>
				{{range .Topics}}
					<li><a href="/nsqtail/{{.}}">{{.}}</a></li>
				{{end}}
			</ul>
		</body>
	</html>
`
