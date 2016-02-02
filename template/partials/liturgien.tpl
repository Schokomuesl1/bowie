<h2>Liturgien</h2>
<div class ="row">
<div class ="col-md-12">
</div>
<form>
<div class="form-group">
		<label for="liturgie">Liturgie:</label>
		<select class="form-control" id = "selectLiturgie" name="liturgie" >
			{{range .AlleLiturgien}}
			<option value="{{.Name}}">{{.Name}} Verbreitung: 
			{{ range .Verbreitung }}
				{{.}} 
			{{ end }}
			SF: {{ .Steigerungsfaktor }}
			{{ end }}
			</option>

		</select>
	</div>
	
	<input type="button" class="btn btn-primary" value="+" onClick="Javascript:addLiturgie()"/>
</div>
</form>
<div class ="row">
<div class ="col-md-12">
<table class="table">
<tr>
<th>Liturgie</th>
<th>Wert</th>
<th>Steigern/Senken</th>
<th>Eigenschaften</th>
</tr>
{{range .Held.Liturgien.Sortiert}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{.Value}}</td>
	<!-- as mouseover <td>[{{.Min}},{{.Max}}]</td>-->
	<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decLiturgie({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incLiturgie({{.Name}})"/></td>
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