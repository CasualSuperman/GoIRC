(function (){

var style = document.getElementById("style");

function resize() {
	var height = window.innerHeight;
	var offset = document.getElementById("networks").clientHeight;
	document.getElementById("log").style.maxHeight = (height - offset) + "px";
}
window.onresize = resize;
function ready() {
	resize();	
}

// This is the style tag reserved for programmatically changing the width of the name column.

})();
