<h2>Neuer Held</h2>
Eigenschaften-Modifikationen verteilen:
<form id="#modEigenForm">
	{{range .}}
	<div class="form-group">
        <label for="{{.Label}}">Modifikation {{.Label}}: Wert {{.Modifikation.Mod}}</label>
		<select class="form-control" id="modifikationSelect{{.Label}}" name="{{.Label}}" >
			{{range .Modifikation.Eigenschaft}}
				<option value="{{.}}">{{.}}</option>
			{{ end }}
		</select>
    </div>
    {{end}}
	<input type="button" class="btn btn-primary" value="Modifikationen &uuml;bernehmen" onClick="Javascript:extractSelectedUpdateEigenschaften()"/>
</form>