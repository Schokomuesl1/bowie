<h2>Neuer Held</h2>
Kampftechnik-Auswahl treffen:
<form id="#selectKTAuswahlForm">
	{{range .}}

	<div class="form-group">
        <label for="{{.Label}}">Modifikations-Satz {{.Label}}: Wert {{.Modifikation.Wert}}</label>
		<select class="form-control" id="kampfwertSelect{{.Label}}" name="{{.Label}}" >
			{{range .Modifikation.Wahlmoeglichkeiten}}
				<option value="{{.}}">{{.}}</option>
			{{ end }}
		</select>
        <label for="{{.Label}}_Wert">Wert</label>
        <p class="form-control-static" id="kampfwertSelect{{.Label}}_Wert" name="{{.Label}}_Wert">{{.Modifikation.Wert}}</p>
    </div>
    {{end}}
	<input type="button" class="btn btn-primary" value="Modifikationen &uuml;bernehmen" onClick="Javascript:extractSelectedKampfwertAuswahl()"/>
</form>


