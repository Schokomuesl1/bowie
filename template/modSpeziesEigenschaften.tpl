<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Held</title>
	</head>
	<body>
		<h1>Neuer Held</h1>
		Eigenschaften-Modifikationen verteilen:
		<form action="/held/new" method="post">
			{{range .}}
		    <div>
		        <label for="{{.Label}}">Modifikation {{.Label}}: Wert {{.Modifikation.Mod}}</label>
				<select name="{{.Label}}" >
					{{range .Modifikation.Eigenschaft}}
						<option value="{{.}}">{{.}}</option>
					{{ end }}
				</select>
		    </div>
		    {{end}}

		    <div class="button">
			    <button type="submit">Modifikationen &uuml;bernehmen</button>
			</div>
		</form>
	</body>
</html>