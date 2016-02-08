<h2>Neuer Held</h2>
Profession ausw&auml;hlen:
<form>
  <div class="form-group">
        <label for="professionsListe">Verf&uuml;gbare Professionen</label>
    <select class="form-control" id="professionsListe" name="professionsListe" >
      <!-- We have the option to chose no profession...-->
      <option selected="selected" value="__none__">Keine</option>
      {{range .Available.ProfessionenNachKulturUndSpezies}}
        <option value="{{.Name}}">{{.Name}} ({{.APKosten}} AP)</option>
      {{ end }}
    </select>
    </div>
  <input type="button" class="btn btn-primary" value="Profession ausw&auml;hlen" onClick="Javascript:extractSelectedProfession()"/>
</form>