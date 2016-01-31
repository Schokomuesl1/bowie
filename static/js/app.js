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
function decKampftechnik(item) {
	doStuff("decrement", "kampftechnik", item);
}
function incKampftechnik(item) {
	doStuff("increment", "kampftechnik", item);
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

function addSF() {
	var e = document.getElementById("SFToAdd")	
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
	//$.get( "/held/action/"+action+"/"+group+"/"+item);
	//window.location.href = "/held/action/"+action+"/"+group+"/"+item

	//
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

