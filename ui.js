(function (){

function resize() {
	var offset = document.getElementsByTagName("header")[0].clientHeight;
	document.getElementById("chat").style.paddingTop = (offset) + "px";
}
function ready() {
	resize();
}

ready();

})();
