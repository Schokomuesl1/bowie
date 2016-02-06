<div class="row">
<div class="col-md-6">
	<h2>Kampftechniken</h2>
</div>
</div>
<div class ="row">
	<div class ="col-md-6">
		<table class="table">
			<thead>
				<tr>
					<th>Name</th>
					<td>Wert</td>
					<td>AT/PA</td>
					<td>+/-</td>
				</tr>
			</thead>
			<tbody>
				{{range .Held.Kampftechniken.Nahkampf}}
				<tr>
					<td>{{.Name}}</td>
					<td>{{.Value}}</td>
					<!--<td>Min: {{.Min}}, Max: {{.Max}}</td>-->
					<td>{{.AT}}/{{.PA}}</td>
					<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decKampftechnik({{.Name}})"/><input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incKampftechnik({{.Name}})"/></td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</div>
	<div class ="col-md-6">
		<table class="table">
			<thead>
				<tr>
					<th>Name</th>
					<td>Wert</td>
					<td>FK</td>
					<td>+/-</td>
				</tr>
			</thead>
			<tbody>
				{{range .Held.Kampftechniken.Fernkampf}}
				<tr>
					<td>{{.Name}}</td>
					<td>{{.Value}}</td>
					<!--<td>Min: {{.Min}}, Max: {{.Max}}</td>-->
					<td>{{.FK}}</td>
					<td><input  {{.KannSenken}} type="button" class="btn btn-default btn-xs" value="-" onClick="Javascript:decKampftechnik({{.Name}})"/> <input  {{.KannSteigern}} type="button" class="btn btn-default btn-xs" value="+" onClick="Javascript:incKampftechnik({{.Name}})"/></td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</div>
</div>
<div class ="row">
	<div class = "col-md-2">
		<h3>Kampf-Sonderfertigkeiten</h3>
	</div>
</div>
<div class="row">
	<div class ="col-md-6">
		<form>
			<div class="form-group">
				<label for="SFToAddKampf">Kampf SF:</label>
				<select name="SFToAddKampf" id="SFToAddKampf">
					{{range .Available.SF_Kampf}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addSF('SFToAddKampf')"/></td>
			</div>
		</form>
	</div>
</div>
<div class ="row">
	<div class ="col-md-6">
		<p>
			{{range .Held.Sonderfertigkeiten.Kampf}}
			{{.Name}}, 
			{{end}}
		</p>
	</div>
</div>