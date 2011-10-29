(function (){

function resize() {
	var offset = document.getElementsByTagName("header")[0].clientHeight;
	document.getElementById("log").style.paddingTop = (offset) + "px";
}
function ready() {
	resize();
	stickToBottom();
}

function stickToBottom() {
	var maxSize = document.getElementById("chat").clientHeight;
	var log = document.getElementById("log");
	var defaultSize = log.clientHeight - log.style.paddingTop.split("px")[0];
	var marginSize  = log.style.marginTop.split("px")[0];
	if(defaultSize > maxSize) {
		var offset = document.getElementsByTagName("header")[0].clientHeight;
		offset -= document.getElementById("gradient").clientHeight / 2;
		log.style.paddingTop = offset + "px";
	} else {
		var val = maxSize - defaultSize - marginSize;
		log.style.paddingTop = val + "px";
	}
	console.log("maxSize: ", maxSize);
	console.log("size: ", defaultSize - marginSize);
	setTimeout(stickToBottom, 300);
}

ready();

})();
