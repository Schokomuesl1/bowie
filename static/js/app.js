function jumpToPage(pagename)
{
	if (!pagename)
		return;
	if (pagename == "Neu")
	{
		putContentToDiv('/held/page/new', '#main_window');
		$("#nav_Neu").siblings('li').removeClass('active');
        $("#nav_Neu").addClass('active');
	}
	else if (pagename == "Allgemeines")
    {
    	putContentToDiv('/held/page/allgemeines', '#main_window');
    	$("#nav_Allgemeines").siblings('li').removeClass('active');
        $("#nav_Allgemeines").addClass('active');

    }
    else if (pagename == "Kampftechniken")
    {
    	putContentToDiv('/held/page/kampftechniken', '#main_window');
    	$("#nav_Kampftechniken").siblings('li').removeClass('active');
        $("#nav_Kampftechniken").addClass('active');

    }
    else if (pagename == "Talente")
    {
    	putContentToDiv('/held/page/talente', '#main_window');
    	$("#nav_Talente").siblings('li').removeClass('active');
        $("#nav_Talente").addClass('active');
    }
    else if (pagename == "Karmales")
    {
      putContentToDiv('/held/page/karmales', '#main_window');
      $("#nav_Karmales").siblings('li').removeClass('active');
        $("#nav_Karmales").addClass('active');
    }
    else if (pagename == "Magie")
    {
      putContentToDiv('/held/page/magie', '#main_window');
      $("#nav_Magie").siblings('li').removeClass('active');
        $("#nav_Magie").addClass('active');
    }
}

function toggleMenuitemActivity(pagename, active)
{
  $("#nav_"+pagename).toggleClass('active', active);
}


function toggleMenuitemVisibility(pagename, visible)
{
  $("#nav_"+pagename).toggleClass('hidden', !visible);
}

// used to load stuff somewhere
function putContentToDiv(pagename, divname)
{
	if (! divname)
	{
		$.get( pagename , function( data )	{ $( "#main_window" ).html( data ); } );
	}
	else
	{
		$.get( pagename , function( data )	{ $( divname ).html( data ); } );	
	}
	
}

function decEigen(item) {
	doStuff("decrement", "eigenschaft", item);
}
function incEigen(item) {
	doStuff("increment", "eigenschaft", item);
}
function decTalent(item) {
	doStuff("decrement", "talent", item);
}
function incTalent(item) {
	doStuff("increment", "talent", item);
}
function decZauber(item) {
  doStuff("decrement", "zauber", item);
}
function incZauber(item) {
  doStuff("increment", "zauber", item);
}
function decLiturgie(item) {
  doStuff("decrement", "liturgie", item);
}
function incLiturgie(item) {
  doStuff("increment", "liturgie", item);
}

function decKampftechnik(item) {
	doStuff("decrement", "kampftechnik", item);
}
function incKampftechnik(item) {
	doStuff("increment", "kampftechnik", item);
}

function addZauber() {
  var e = document.getElementById("selectZauber") 
  var selectedItem = e.options[e.selectedIndex].value;
  doStuff("add", "zauber", selectedItem);
}

function addLiturgie() {
  var e = document.getElementById("selectLiturgie") 
  var selectedItem = e.options[e.selectedIndex].value;
  doStuff("add", "liturgie", selectedItem);
}

function addVorteil() {
	var e = document.getElementById("VorteilToAdd")	
	var selectedItem = e.options[e.selectedIndex].value;
	doStuff("add", "vorteil", selectedItem);
}

function addNachteil() {
	var e = document.getElementById("NachteilToAdd")	
	var selectedItem = e.options[e.selectedIndex].value;
	doStuff("add", "nachteil", selectedItem);
}

function addSF(bereich) {
	var e = document.getElementById(bereich)	
	var selectedItem = e.options[e.selectedIndex].value;
  doStuff("add", bereich, selectedItem);
}

function removeVTNT(name) {
	doStuff("remove", "vorteilnachteil", name);
}

function removeSF(bereich, name) {
	doStuff("remove", bereich, name);
}

function checkForRedirect(data, status)
{
	if (data.hasOwnProperty('redirectTo'))
	{
		if (data['redirectTo'] != "") {
			putContentToDiv(data['redirectTo'])
			updateProgressBar()
		}
	}
	if (data.hasOwnProperty('magie'))
	{
		toggleMenuitemVisibility('Magie', data['magie']);
  	}
	if (data.hasOwnProperty('karmal'))
	{
	  	toggleMenuitemVisibility('Karmales', data['karmal']);
	}
	var content = "";
	if (data.hasOwnProperty('validatorMessages'))
	{
		content = generateWarningsAndErrors(data['validatorMessages']);
	}
	$( "#warnings_and_errors" ).html( content );
	var notification = "";
	if (data.hasOwnProperty('notificationMsg'))
	{
		notification = data['notificationMsg'];
	}
	if (notification != "") {
		$.notify({
			// options
			icon: 'glyphicon glyphicon-warning-sign',
			title: 'Warnung:',
			message: notification
		},{
			// settings
			element: 'body',
			position: null,
			type: "warning",
			allow_dismiss: true,
			newest_on_top: false,
			showProgressbar: false,
			placement: {
				from: "top",
				align: "right"
			},
			offset: 20,
			spacing: 10,
			z_index: 1031,
			delay: 5000,
			timer: 1000,
			url_target: '_blank',
			mouse_over: null,
			animate: {
				enter: 'animated bounceInLeft',
				exit: 'animated bounceOutDown'
			},
			onShow: null,
			onShown: null,
			onClose: null,
			onClosed: null,
			icon_type: 'class',
			template: '<div data-notify="container" class="col-xs-11 col-sm-3 alert alert-{0}" role="alert">' +
				'<button type="button" aria-hidden="true" class="close" data-notify="dismiss">×</button>' +
				'<span data-notify="icon"></span> ' +
				'<span data-notify="title">{1}</span> ' +
				'<span data-notify="message">{2}</span>' +
				'<a href="{3}" target="{4}" data-notify="url"></a>' +
			'</div>' 
		});
	}
}


// this is a hack - each click replaces the whole page. Rework this after switchung to a sensible API
function doStuff(action, group, item) {	
    $.post("/held/action/"+action+"/"+group+"/"+item,
    	{},
    	checkForRedirect
    );
};

function extractSelectedNewHeld() {
	var request = new Object();
	var e = document.getElementById("selectspezies");
	request[e.name] = e.options[e.selectedIndex].value;
	e = document.getElementById("selectkultur");
	request[e.name] = e.options[e.selectedIndex].value;
	 e = document.getElementById("selecterfahrungsgrad");
	request[e.name] = e.options[e.selectedIndex].value;
	e = document.getElementById("inputHeldName");
	request[e.name] = e.value;
	request.type = "createHeld"
	sendPostWithJSONTo("/held/complexaction", request)
}

function extractSelectedProfession() {
  var request = new Object();
  var e = document.getElementById("professionsListe")
  request["profession"] = e.options[e.selectedIndex].value
  request.type = "selectProfession"
  sendPostWithJSONTo("/held/complexaction", request)

  // now we set the active window and make the relevant items visible

  toggleMenuitemVisibility('Allgemeines', true);
  toggleMenuitemActivity('Allgemeines', true);
  toggleMenuitemActivity('Neu', false);
  toggleMenuitemVisibility('Kampftechniken', true);
  toggleMenuitemVisibility('Talente', true);
}

function extractSelectedUpdateEigenschaften() {
	var request = new Object();
	var i = 0
	while (!(!document.getElementById("modifikationSelect"+i.toString()))) {
		var e = document.getElementById("modifikationSelect"+i.toString());
		request[i.toString()] = e.options[e.selectedIndex].value;
		i++
	}
	request.type = "modEigenschaften"
	sendPostWithJSONTo("/held/complexaction", request)
}

function extractSelectedKampfwertAuswahl() {
    var request = new Object();
    var i = 0
    while (!(!document.getElementById("kampfwertSelect"+i.toString()))) {
        var e = document.getElementById("kampfwertSelect"+i.toString());
        var w = document.getElementById("kampfwertSelect"+i.toString()+"_Wert");
        request["kw" + i.toString()] = e.options[e.selectedIndex].value;
        request["kw_wert_" + i.toString()] = w.innerHTML;
        i++
    }
    request.type = "selectKampfwerte"
    sendPostWithJSONTo("/held/complexaction", request)
}


function sendPostWithJSONTo(target, jsonContent)
{
	$.post(target,
    	jsonContent,
    	checkForRedirect
	);	
}

// update progress bar
function updateProgressBar()
{
	putContentToDiv("/held/page/footer", "#page_footer")
}

function generateWarningsAndErrors(data)
{
	var retVal = ""
	for (i in data) {
		if (data[i]['Type'] == '2') {
			// warning
			retVal += "<div class='alert alert-warning'> <strong>Warnung!</strong> ";
			retVal += data[i]['Msg'];
			retVal += "</div>";
		}
		else if (data[i]['Type'] == '3') {
			retVal += "<div class='alert alert-danger'> <strong>Fehler!</strong> ";
			retVal += data[i]['Msg'];
			retVal += "</div>";
		}
	}
	return retVal
}
