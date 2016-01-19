<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Held</title>
	</head>
	<body>
		<h1>Held</h1>
		<table>
		<tr><td>Name</td><td>{{.Held.Name}}</td></tr>
		<tr><td>Spezies</td><td>{{.Held.Spezies.Name}}</td></tr>
		<tr><td>Kultur</td><td>{{.Held.Kultur.Name}}</td></tr>
		</table>
		<h2>Eigenschaften</h2>
		<table>
			{{range .Held.Eigenschaften.Eigenschaften}}<tr><td>{{.Name}}</td><td>{{.Value}}</td></tr>{{end}}
		</table>
		<h2>Grundwerte</h2>
		<table>
			<tr><td>Lebensenergie</td>  <td>{{.Held.Basiswerte.Lebensenergie.Value}}</td>  </tr>
			<tr><td>Astralenergie</td>  <td>{{.Held.Basiswerte.Astralenergie.Value}}</td>  </tr>
			<tr><td>Karmaenergie</td>   <td>{{.Held.Basiswerte.Karmaenergie.Value}}</td>   </tr>
			<tr><td>Seelenkraft</td>    <td>{{.Held.Basiswerte.Seelenkraft.Value}}</td>    </tr>
			<tr><td>Zaehigkeit</td>     <td>{{.Held.Basiswerte.Zaehigkeit.Value}}</td>     </tr>
			<tr><td>Ausweichen</td>     <td>{{.Held.Basiswerte.Ausweichen.Value}}</td>     </tr>
			<tr><td>Initiative</td>     <td>{{.Held.Basiswerte.Initiative.Value}}</td>     </tr>
			<tr><td>Geschwindigkeit</td><td>{{.Held.Basiswerte.Geschwindigkeit.Value}}</td></tr>
		</table>
		<h2>Validierung</h2>
		<ul>
			{{range .Msg}} <li>{{.Msg}}</li>{{end}}
		</ul>
		<h2>Talente</h2>
		<table>
			{{range .Held.Talente.Talente}}<tr><td>{{.Name}}</td><td>{{.Value}}</td>{{range .Eigenschaften}}<td>{{.}}</td>{{end}}</tr>{{end}}
		</table>
	</body>
</html>