<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Held</title>
		<!--<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.2.0/jquery.min.js"></script>-->
	</head>
	<body>
		<h1>Held</h1>
		<table>
			<tr><td>Name</td><td>{{.Held.Name}}</td></tr>
			<tr><td>Spezies</td><td>{{.Held.Spezies.Name}}</td></tr>
			<tr><td>Kultur</td><td>{{.Held.Kultur.Name}}</td></tr>
			<tr><td>AP</td><td>
				<table>
					<tr><td>Gesamt:</td><td>{{.Held.APGesamt}}</td></tr>
					<tr><td>Verfügbar:</td><td>{{.Held.AP}}</td></tr>
					<tr><td>Ausgegeben:</td><td>{{.Held.AP_spent}}</td></tr>
				</table>
			</td></tr>
		</table>
		<h2>Eigenschaften</h2>
		<table>
		{{range .Held.Eigenschaften.Eigenschaften}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Value}}</td>
				<td>Min: {{.Min}}, Max: {{.Max}}</td>
				<td><input  {{.KannSenken}} type="button" value="-" onClick="Javascript:decEigen({{.Name}})"/></td>
				<td><input  {{.KannSteigern}} type="button" value="+" onClick="Javascript:incEigen({{.Name}})"/></td>
			</tr>
		{{end}}
		</table>
		
		<h2>Kampftechniken</h2>
		<table>
		{{range .Held.Kampftechniken.Kampftechniken}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Value}}</td>
				<td>Min: {{.Min}}, Max: {{.Max}}</td>
				<td>AT: {{.AT}}</td><td>PA: {{.PA}}</td><td>FK: {{.FK}}</td>
				<td><input  {{.KannSenken}} type="button" value="-" onClick="Javascript:decKampftechnik({{.Name}})"/></td>
				<td><input  {{.KannSteigern}} type="button" value="+" onClick="Javascript:incKampftechnik({{.Name}})"/></td>
			</tr>
		{{end}}
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
			{{range .ValidatorMsg}} <li>{{.Msg}}</li>{{end}}
		</ul>
		<h2>Talente</h2>
		<table>
		{{range .Held.Talente.Talente}}
			<tr><td>{{.Name}}</td><td>{{.Value}}</td><td>[{{.Min}},{{.Max}}]</td>
			<td><input  {{.KannSenken}} type="button" value="-" onClick="Javascript:decTalent({{.Name}})"/></td>
			<td><input  {{.KannSteigern}} type="button" value="+" onClick="Javascript:incTalent({{.Name}})"/></td>
			<td>
			{{range .Eigenschaften}}
				<td>{{.}}</td>
			{{end}}
			</tr>
		{{end}}
		</table>
		{{range .Held.Eigenschaften.Eigenschaften}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Value}}</td>
				<td>Min: {{.Min}}, Max: {{.Max}}</td>
				
			</tr>
		{{end}}
	</body>
	<script>
	function decEigen(item) {
		doStuff("decrement", "eigenschaft", item);
	}
	function incEigen(item) {
		doStuff("increment", "eigenschaft", item);
	}
	function decTalent(item) {
		doStuff("decrement", "talent", item);
	}
	function incTalent(item) {
		doStuff("increment", "talent", item);
	}
	function decKampftechnik(item) {
		doStuff("decrement", "kampftechnik", item);
	}
	function incKampftechnik(item) {
		doStuff("increment", "kampftechnik", item);
	}
	
	// this is a hack - each click replaces the whole page. Rework this after switchung to a sensible API
	function doStuff(action, group, item) {
		console.log("/held/action/"+action+"/"+group+"/"+item);
		//$.get( "/held/action/"+action+"/"+group+"/"+item);
		window.location.href = "/held/action/"+action+"/"+group+"/"+item
	};
	</script>
</html>