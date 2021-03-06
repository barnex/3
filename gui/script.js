
// auto-refresh rate
var tick = 300;
var autorefresh = true;

// show error in document (non-intrusive alert())
function showErr(err){
	document.getElementById("ErrorBox").innerHTML = err;
}

// show debug message in document
function msg(err){
	document.getElementById("MsgBox").innerHTML = err;
}

// wraps document.getElementById, shows error if not found
function elementById(id){
	var elem = document.getElementById(id);
	if (elem == null){
		showErr("undefined: " + id);
		return null;
	}
	return elem;
}

// called on change of auto-refresh button
function setautorefresh(){
	autorefresh =  elementById("AutoRefresh").checked;
}

// Id of element that has focus. We don't auto-refresh a focused textbox
// as this would overwrite the users input.
var hasFocus = "";
function notifyfocus(id){hasFocus = id;}
function notifyblur (id){hasFocus = "";}

// called by server to manipulate the DOM
function setAttr(id, attr, value){
	var elem = elementById(id);
	if (elem == null){
		return;
	}
	if (elem[attr] == null){
		showErr("settAttr: undefined: " + id + "[" + attr + "]");
		return;
	}
	elem[attr] = value;
}

// set textbox value unless focused
function setTextbox(id, value){
	if (hasFocus != id){
		elementById(id).value = value;
	}
}

// set select value unless focused
function setSelect(id, value){
	if (hasFocus != id){
		elementById(id).value = value;
	}
}


// onreadystatechange function for update http request.
// refreshes the DOM with new values received from server.
function refreshDOM(req){
	if (req.readyState == 4) { // DONE
		if (req.status == 200) {	
			showErr("");
			var response = JSON.parse(req.responseText);	
			for(var i=0; i<response.length; i++){
				var r = response[i];
				var func = window[r.F];
				if (func == null) {
					showErr("undefined: " + r.F);
				}else{ 
					func.apply(this, r.Args);
				}
			}
		} else {
			showErr("Disconnected");	
		}
	}
}

// refreshes the contents of all dynamic elements.
// periodically called via setInterval()
function refresh(){
		try{
			var req = new XMLHttpRequest();
			req.open("POST", document.URL, true); 
			req.timeout = tick;
			req.onreadystatechange = function(){ refreshDOM(req) };
			req.send(null);
		}catch(e){
			showErr(e); // TODO: same message as refresh
		}
}

function doAutorefresh(){
	if (autorefresh){
		refresh();
	}
}

setInterval(doAutorefresh, tick);

// sends event notification to server, called on button clicks etc.
function notify(id, arg){
	try{
		var req = new XMLHttpRequest();
		req.open("PUT", document.URL, false);
		var map = {"ID": id, "Arg": arg};
		req.send(JSON.stringify(map));
	}catch(e){
		showErr(e); // TODO
	}
	refresh();
}

function notifyButton(id){
	notify(id, elementById(id).innerHTML);
}

function notifytextbox(id){
	notify(id, elementById(id).value);
}

function notifycheckbox(id){
	notify(id, elementById(id).checked);
}

function notifyrange(id){
	notify(id, elementById(id).value);
}

function notifyselect(id){
	var e = elementById(id);
	var value = e.options[e.selectedIndex].text;
	notify(id, value);
}

window.onload = refresh;
