<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Held</title>
	</head>
	<body>
		<h1>Neuer Held</h1>
		<form action="/held/modEigenschaften" method="post">
			<div>
		        <label for="heldName">Name:</label>
		        <input type="text" name="heldName" />
		    </div>
		    <div>
		        <label for="spezies">Spezies:</label>
				<select name="spezies" >
					{{range .AlleSpezies}}
						<option value="{{.Name}}">{{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{ end }}
				</select>
		    </div>
		    <div>
		        <label for="kultur">Kultur:</label>
		        <select name="kultur" >
					{{range .AlleKulturen}}
						<option value="{{.Name}}">{{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{ end }}
				</select>
		    </div>
			<div>
		        <label for="erfahrungsgrad">Erfahrungsgrad:</label>
				<select name="erfahrungsgrad" >
					{{range .Grade}}
						<option value="{{.Name}}">{{.Name}} <i>({{.AP}} AP)</i></option>
					{{ end }}
				</select>
		    </div>
		    <div class="button">
			    <button type="submit">Held erstellen</button>
			</div>
		</form>
	</body>
</html>