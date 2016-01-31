<h2>Neuer Held</h2>

<form>
	<div class="form-group">
		<label for="heldName">Name:</label>
		<input class="form-control" id="inputHeldName" type="text" name="heldName" />
	</div>
	<div class="form-group">
		<label for="spezies">Spezies:</label>
		<select class="form-control"  id = "selectspezies" name="spezies" >
			{{range .AlleSpezies}}
			<option value="{{.Name}}">{{.Name}} <i>({{.APKosten}} AP)</i></option>
			{{ end }}
		</select>
	</div>
	<div class="form-group">
		<label for="kultur">Kultur:</label>
		<select class="form-control" id = "selectkultur" name="kultur" >
			{{range .AlleKulturen}}
			<option value="{{.Name}}">{{.Name}} <i>({{.APKosten}} AP)</i></option>
			{{ end }}
		</select>
	</div>
	<div class="form-group">
		<label for="erfahrungsgrad">Erfahrungsgrad:</label>
		<select class="form-control"  id = "selecterfahrungsgrad" name="erfahrungsgrad" >
			{{range .Grade}}
			<option value="{{.Name}}">{{.Name}} <i>({{.AP}} AP)</i></option>
			{{ end }}
		</select>
	</div>
	<input type="button" class="btn btn-primary" value="Held erstellen" onClick="Javascript:extractSelectedNewHeld()"/>
</form>