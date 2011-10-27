<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>goIRC</title>
		<!--[if lt IE 9]>
		<link rel="stylesheet" href="/html5/fixes/html5.css" />
		<script src="/html5/fixes/html5.js" type="text/javascript"></script>
		<![endif]-->
		<link rel="stylesheet" href="/~rwertman/libs/php/font.php?n=Segoe UI&f=Segoe&b" />
		<link rel="stylesheet" href="main.css" media="screen" />
		<link rel="stylesheet" href="small.css" media="handheld" />
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.5.2/jquery.js"></script>
		<script src="/~rwertman/libs/js/jquery.transition.js"></script>
		<script src="/~rwertman/libs/js/jquery.color.clo.js"></script>
		<script src="ui.js"></script>
<!--		<script src="/libs/js/socket.io.min.js"></script>
		<script src="ui.js"></script>
		<script src="app.js"></script>-->
	</head>
<!--
Copyright 2011 Bobby Wertman.
-->
<body>
<div id="chat">
	<div id="networks">
		<ul>
			<li class="network active">foonetic</li>
			<li class="network nonitem">+connect</li>
		</ul>
		<div id="chats">
			<ul>
				<li class="chat active">#ufeff</li>
				<li class="chat">#xkcd</li>
				<li class="chat nonitem">+join</li>
			</ul>
		</div>
		<div class="clearfix"></div>
	</div>
	<table id="log">
		<tbody>
			<tr>
				<td class="author">CasualSuperman</td>
				<td>sup?</td>
			</tr>
			<tr>
				<td class="author op">@s0urc3</td>
				<td>nm, sup?</td>
			</tr>
			<tr>
				<td class="author">CasualSuperman</td>
				<td>Well, I'm working on this interface for the web IRC thingie, and I need to add a lot of text to what I'm saying right now so I can be sure it will wrap onto multiple lines so I can make sure names will be aligned to the top of text-blocks. That's about it.</td>
			</tr>
		</tbody>
		<tfoot>
			<tr id="input">
				<td id="name">CasualSuperman (irsxw)</td>
				<td><input class="chat" /></td>
			</tr>
		</tfoot>
	</table>
	<style type="text/css" id="style"></style>
</div>
<!--
Kind of TODO:
	* Tab autocompletes names
	* Autoscrolling on new messages
	* Clicking a name opens a dialog with actions
	* Max-height on incoming messages area
	* Also making it actually work.
-->
</body>
</html>
