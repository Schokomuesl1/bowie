<h2>Allgemeines</h2>
<div class="row">
	<div class ="col-md-2">
		<table class="table">
			<tr><td>Name</td>     <td>{{.Held.Name}}        </td></tr>
			<tr><td>Spezies</td>  <td>{{.Held.Spezies.Name}}</td></tr>
			<tr><td>Kultur</td>   <td>{{.Held.Kultur.Name}} </td></tr>
		</table>
	</div>
	<div class ="col-md-2">
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
	<div class ="col-md-2">
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
	<div class ="col-md-6">
		<table class="table">
			<tr>
				<td>Vorteile</td>
				<td>
					{{range .Held.Vorteile}}
					{{.Name}}<a href="javascript:removeVTNT('{{.Name}}');"><span class='{{.DeleteButtonIfApplicable}}'></span></a>,
					{{end}}
				</td>
			</tr>
			<tr>
				<td>Nachteile</td>
				<td>
					{{range .Held.Nachteile}}
					{{.Name}}<a href="javascript:removeVTNT('{{.Name}}');"><span class='{{.DeleteButtonIfApplicable}}'></span></a>,
					{{end}}
				</td>
			</tr>
			<tr>
				<td>Allgemeine SF</td>
				<td>
					{{range .Held.Sonderfertigkeiten.Allgemeine}}
					{{.Name}} <a href="javascript:removeSF('SFToAddAllgemein', '{{.Name}}');"><span class="text-danger glyphicon glyphicon-remove"></span></a>, 
					{{end}}
				</td>
			</tr>
			<tr>
				<td>Sprachen</td>
				<td>
					{{range .Held.Sonderfertigkeiten.Sprachen}}
					{{.Name}} <a href="javascript:removeSF('SFToAddSprache', '{{.Name}}');"><span class="text-danger glyphicon glyphicon-remove"></span></a>, 
					{{end}}
				</td>
			</tr>
			<tr>
				<td>Schriften</td>
				<td>
					{{range .Held.Sonderfertigkeiten.Schriften}}
					{{.Name}} <a href="javascript:removeSF('SFToAddSchrift', '{{.Name}}');"><span class="text-danger glyphicon glyphicon-remove"></span></a>, 
					{{end}}
				</td>
			</tr>
		</table>

		<form>
			<div class="form-group">
				<label for="VorteilToAdd">Vorteil:</label>
				<select class="form-control" name="VorteilToAdd" id="VorteilToAdd">
					{{range .Available.Vorteile}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addVorteil()"/></td>
			</div>
			<div class="form-group">
				<label for="NachteilToAdd">Nachteil:</label>
				<select class="form-control" name="NachteilToAdd" id="NachteilToAdd">
					{{range .Available.Nachteile}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addNachteil()"/></td>
			</div>
			<div class="form-group">
				<label for="SFToAddAllgemein">Allgemeine SF:</label>
				<select class="form-control" name="SFToAddAllgemein" id="SFToAddAllgemein">
					{{range .Available.SF_Allgemein}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addSF('SFToAddAllgemein')"/></td>
			</div>
			<div class="form-group">
				<label for="SFToAddSprache">Sprachen:</label>
				<select class="form-control" name="SFToAddSprache" id="SFToAddSprache">
					{{range .Available.SF_Sprachen}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addSF('SFToAddSprache')"/></td>
			</div>
			<div class="form-group">
				<label for="SFToAddSchrift">Schriften:</label>
				<select class="form-control" name="SFToAddSchrift" id="SFToAddSchrift">
					{{range .Available.SF_Schriften}}
					<option value="{{.Name}}"> {{.Name}} <i>({{.APKosten}} AP)</i></option>
					{{end}}
				</select>
				<input type="button" value="+" onClick="Javascript:addSF('SFToAddSchrift')"/></td>
			</div>
		</form>
	</div>
</div>
