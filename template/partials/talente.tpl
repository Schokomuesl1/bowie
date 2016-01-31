<h2>Talente</h2>
<div class ="row">
<div class ="col-md-6">
<h3>Körperlich</h3>
<table class="table">
<tr>
<th>Talent</th>
<th>Wert</th>
<th>Steigern/Senken</th>
<th>Eigenschaften</th>
</tr>
{{range .Held.Talente.Koerpertalente}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{.Value}}</td>
	<!-- as mouseover <td>[{{.Min}},{{.Max}}]</td>-->
	<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decTalent({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incTalent({{.Name}})"/></td>
	<td>
	{{range .Eigenschaften}}
		{{.Name}} 
	{{end}}
	</td>
	</tr>
{{end}}
</table>
<h3>Natur</h3>
<table class="table">
<tr>
<th>Talent</th>
<th>Wert</th>
<th>Steigern/Senken</th>
<th>Eigenschaften</th>
</tr>
{{range .Held.Talente.Naturtalente}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{.Value}}</td>
	<!-- as mouseover <td>[{{.Min}},{{.Max}}]</td>-->
	<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decTalent({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incTalent({{.Name}})"/></td>
	<td>
	{{range .Eigenschaften}}
		{{.Name}} 
	{{end}}
	</td>
	</tr>
{{end}}
</table>
<h3>Gesellschaft</h3>
<table class="table">
<tr>
<th>Talent</th>
<th>Wert</th>
<th>Steigern/Senken</th>
<th>Eigenschaften</th>
</tr>
{{range .Held.Talente.Gesellschaftstalente}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{.Value}}</td>
	<!-- as mouseover <td>[{{.Min}},{{.Max}}]</td>-->
	<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decTalent({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incTalent({{.Name}})"/></td>
	<td>
	{{range .Eigenschaften}}
		{{.Name}} 
	{{end}}
	</td>
	</tr>
{{end}}
</table>
</div>

<div class ="col-md-6">
<h3>Wissen</h3>
<table class="table">
<tr>
<th>Talent</th>
<th>Wert</th>
<th>Steigern/Senken</th>
<th>Eigenschaften</th>
</tr>
{{range .Held.Talente.Wissenstalente}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{.Value}}</td>
	<!-- as mouseover <td>[{{.Min}},{{.Max}}]</td>-->
	<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decTalent({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incTalent({{.Name}})"/></td>
	<td>
	{{range .Eigenschaften}}
		{{.Name}} 
	{{end}}
	</td>
	</tr>
{{end}}
</table>
<h3>Handwerk</h3>
<table class="table">
<tr>
<th>Talent</th>
<th>Wert</th>
<th>Steigern/Senken</th>
<th>Eigenschaften</th>
</tr>
{{range .Held.Talente.Handwerkstalente}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{.Value}}</td>
	<!-- as mouseover <td>[{{.Min}},{{.Max}}]</td>-->
	<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decTalent({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incTalent({{.Name}})"/></td>
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
