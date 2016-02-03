<h2>Eigenschaften</h2>
<div class="row">
	<div class ="col-md-4">
		<table class="table">
			{{range .Held.Eigenschaften.Eigenschaften}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Value}}</td>
				<!--<td>Min: {{.Min}}, Max: {{.Max}}</td>-->
				<td><input  {{.KannSenken}} type="button" value="-" onClick="Javascript:decEigen({{.Name}})"/> <input  {{.KannSteigern}} type="button" value="+" onClick="Javascript:incEigen({{.Name}})"/></td>
			</tr>
			{{end}}
		</table>
	</div>
	<div class ="col-md-4">
		<table class="table">
			<tr><td>Lebensenergie</td>  <td>{{.Held.Basiswerte.Lebensenergie.Value}}</td>  </tr>
			<tr><td>Astralenergie</td>  <td>{{.Held.Basiswerte.Astralenergie.Value}}</td>  </tr>
			<tr><td>Karmaenergie</td>   <td>{{.Held.Basiswerte.Karmaenergie.Value}}</td>   </tr>
			<tr><td>Seelenkraft</td>    <td>{{.Held.Basiswerte.Seelenkraft.Value}}</td>    </tr>
			<tr><td>Zaehigkeit</td>     <td>{{.Held.Basiswerte.Zaehigkeit.Value}}</td>     </tr>
			<tr><td>Ausweichen</td>     <td>{{.Held.Basiswerte.Ausweichen.Value}}</td>     </tr>
			<tr><td>Initiative</td>     <td>{{.Held.Basiswerte.Initiative.Value}}</td>     </tr>
			<tr><td>Geschwindigkeit</td><td>{{.Held.Basiswerte.Geschwindigkeit.Value}}</td></tr>
		</table>
	</div>
	<div class ="col-md-4">
		<table class="table">
			<tr>
				<td>Vorteile</td>
				<td>
					{{range .Held.Vorteile}}
					{{.Name}}, 
					{{end}}
				</td>
			</tr>
			<tr>
				<td>Nachteile</td>
				<td>
					{{range .Held.Nachteile}}
					{{.Name}}, 
					{{end}}
				</td>
			</tr>
			<tr>
				<td>Sonderfertigkeiten</td>
				<td>
					{{range .Held.Sonderfertigkeiten}}
					{{.Name}}, 
					{{end}}
				</td>
			</tr>
		</table>

		<form>
			<div class="form-group">
				<label for="VorteilToAdd">Vorteil:</label>
				<select name="VorteilToAdd" id="VorteilToAdd">
					{{range .Available.Vorteile}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addVorteil()"/></td>
			</div>
			<div class="form-group">
				<label for="NachteilToAdd">Nachteil:</label>
				<select name="NachteilToAdd" id="NachteilToAdd">
					{{range .Available.Nachteile}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addNachteil()"/></td>
			</div>
			<div class="form-group">
				<label for="SFToAddAllgemein">Allgemeine SF:</label>
				<select name="SFToAddAllgemein" id="SFToAddAllgemein">
					{{range .Available.SF_Allgemein}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addSF('SFToAddAllgemein')"/></td>
			</div>
			<div class="form-group">
				<label for="SFToAddKarmal">Karmale SF:</label>
				<select name="SFToAddKarmal" id="SFToAddKarmal">
					{{range .Available.SF_Karmal}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addSF('SFToAddKarmal')"/></td>
			</div>
			<div class="form-group">
				<label for="SFToAddMagisch">Magische SF:</label>
				<select name="SFToAddMagisch" id="SFToAddMagisch">
					{{range .Available.SF_Magisch}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addSF('SFToAddMagisch')"/></td>
			</div>
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

<!--
<h2>Grundwerte</h2>
<table class="table">
	<tr><td>Lebensenergie</td>  <td>{{.Held.Basiswerte.Lebensenergie.Value}}</td>  </tr>
	<tr><td>Astralenergie</td>  <td>{{.Held.Basiswerte.Astralenergie.Value}}</td>  </tr>
	<tr><td>Karmaenergie</td>   <td>{{.Held.Basiswerte.Karmaenergie.Value}}</td>   </tr>
	<tr><td>Seelenkraft</td>    <td>{{.Held.Basiswerte.Seelenkraft.Value}}</td>    </tr>
	<tr><td>Zaehigkeit</td>     <td>{{.Held.Basiswerte.Zaehigkeit.Value}}</td>     </tr>
	<tr><td>Ausweichen</td>     <td>{{.Held.Basiswerte.Ausweichen.Value}}</td>     </tr>
	<tr><td>Initiative</td>     <td>{{.Held.Basiswerte.Initiative.Value}}</td>     </tr>
	<tr><td>Geschwindigkeit</td><td>{{.Held.Basiswerte.Geschwindigkeit.Value}}</td></tr>
</table>



	<label for="VorteilToAdd">Hinzufügen:</label>
	<select name="VorteilToAdd" id="VorteilToAdd">
		{{range .Available.Vorteile}}
			<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
		{{end}}
	</select>
	<input type="button" value="+" onClick="Javascript:addVorteil()"/></td>
	<label for="NachteilToAdd">Hinzufügen:</label>
	<select name="NachteilToAdd" id="NachteilToAdd">
		{{range .Available.Nachteile}}
			<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
		{{end}}
	</select>
	<input type="button" value="+" onClick="Javascript:addNachteil()"/></td>
	<label for="SFToAdd">Hinzufügen:</label>
	<select name="SFToAdd" id="SFToAdd">
		{{range .Available.SF_Allgemein}}
			<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
		{{end}}
	</select>
	<input type="button" value="+" onClick="Javascript:addSF()"/></td>
</form>
-->