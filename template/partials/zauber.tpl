<h2>Zauber</h2>
<div class ="row">
<div class ="col-md-12">
</div>
<form>
<div class="form-group">
		<label for="zauber">Zauber:</label>
		<select class="form-control" id = "selectZauber" name="zauber" >
			{{range .AlleZauber}}
			<option value="{{.Name}}">{{.Name}} Verbreitung: 
			{{ range .Verbreitung }}
				{{.}} 
			{{ end }}
			SF: {{ .Steigerungsfaktor }}
			{{ end }}
			</option>

		</select>
	</div>
	
	<input type="button" class="btn btn-primary" value="+" onClick="Javascript:addZauber()"/>
</div>
</form>
<div class ="row">
<div class ="col-md-12">
<table class="table">
<tr>
<th>Zauber</th>
<th>Wert</th>
<th>Steigern/Senken</th>
<th>Eigenschaften</th>
</tr>
{{range .Held.Zauber.Sortiert}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{.Value}}</td>
	<!-- as mouseover <td>[{{.Min}},{{.Max}}]</td>-->
	<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decZauber({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incZauber({{.Name}})"/></td>
	<td>
	{{range .Eigenschaften}}
		{{.Name}} 
	{{end}}
	</td>
	</tr>
{{end}}
</table>
</div>
</div>