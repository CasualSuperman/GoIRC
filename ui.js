(function (){

function resize() {
	var height = window.innerHeight;
	var offset = document.getElementById("networks").clientHeight;
	document.getElementById("log").style.maxHeight = (height - offset) + "px";
}
window.onresize = resize;
function ready() {
	resize();	
}

})();
