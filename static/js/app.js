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
    else if (pagename == "Liturgien")
    {
      putContentToDiv('/held/page/liturgien', '#main_window');
      $("#nav_Liturgien").siblings('li').removeClass('active');
        $("#nav_Liturgien").addClass('active');
    }
    else if (pagename == "Zauber")
    {
      putContentToDiv('/held/page/zauber', '#main_window');
      $("#nav_Zauber").siblings('li').removeClass('active');
        $("#nav_Zauber").addClass('active');
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
	console.log("putContentToDiv")
	console.log(pagename)
	console.log(divname)
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
	doStuff("add", "sf", selectedItem);
}

function checkForRedirect(data, status)
{
	//var obj = $.parseJSON(data);
	if (data.hasOwnProperty('redirectTo'))
	{
		putContentToDiv(data['redirectTo'])
		updateProgressBar()
	}
	
}


// this is a hack - each click replaces the whole page. Rework this after switchung to a sensible API
function doStuff(action, group, item) {	
	console.log("/held/action/"+action+"/"+group+"/"+item);

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

function extractSelectedUpdateEigenschaften() {
	var request = new Object();
	var i = 0
	console.log(document.getElementById("modifikationSelect"+i.toString()))
	while (!(!document.getElementById("modifikationSelect"+i.toString()))) {
		console.log(document.getElementById("modifikationSelect"+i.toString()))
		var e = document.getElementById("modifikationSelect"+i.toString());
		request[i.toString()] = e.options[e.selectedIndex].value;
		i++
	}
	request.type = "modEigenschaften"
	console.log(request)
	sendPostWithJSONTo("/held/complexaction", request)
  // now we set the active window and make the relevant items visible

  toggleMenuitemVisibility('Allgemeines', true);
  toggleMenuitemActivity('Allgemeines', true);
  toggleMenuitemActivity('Neu', false);
  toggleMenuitemVisibility('Kampftechniken', true);
  toggleMenuitemVisibility('Talente', true);
  toggleMenuitemVisibility('Zauber', true);
  toggleMenuitemVisibility('Liturgien', true);
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

