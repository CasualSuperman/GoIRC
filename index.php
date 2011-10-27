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
	<header>
		<ul id="networks">
			<li class="network active">foonetic</li>
			<li class="network nonitem">+connect</li>
		</ul>
		<ul id="rooms">
			<li class="chat active">#ufeff</li>
			<li class="chat">#xkcd</li>
			<li class="chat nonitem">+join</li>
		</ul>
	</header>
	<table id="chat">
		<tbody id="log" valign="bottom">
			<tr class="status">
				<td class="author">User</td>
				<td>
					<span class="me">CasualSuperman</span> has joined.
				</td>
			</tr>
			<tr>
				<td class="author me">CasualSuperman</td>
				<td>sup?</td>
			</tr>
			<tr>
				<td class="author op mention">@s0urc3</td>
				<td>nm, sup <span class="mention">CasualSuperman</span>?</td>
			</tr>
			<tr>
				<td class="author me">CasualSuperman</td>
				<td>Well, I'm working on this interface for the web IRC thingie, and I need to add a lot of text to what I'm saying right now so I can be sure it will wrap onto multiple lines so I can make sure names will be aligned to the top of text-blocks. That's about it.</td>
			</tr>
			<tr>
				<td class="author op">@s0urc3</td>
				<td>Here ya go.</td>
			</tr>
			<tr>
				<td class="author op">@s0urc3</td>
				<td>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean sed quam quam. Vivamus id lectus sit amet libero porta porta. Nullam vulputate eleifend dapibus. Nulla iaculis nibh eu ipsum pretium ultricies eu vitae nunc. Aliquam ullamcorper quam in est dapibus semper. Aliquam erat volutpat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ac lorem consectetur sapien imperdiet scelerisque. In et convallis tellus. Quisque vitae ligula in mauris vestibulum dapibus. Nam hendrerit congue felis a iaculis. Donec vel neque leo. Donec imperdiet sapien eget magna commodo id ornare mauris pretium. Nam nulla massa, dictum eget lacinia vel, fermentum commodo ipsum. Sed nunc magna, sodales sit amet suscipit non, pellentesque a libero. Praesent nulla sem, adipiscing id elementum ut, volutpat vitae purus. Duis leo est, cursus nec mollis sit amet, suscipit quis leo. Vivamus sodales justo nec odio pulvinar adipiscing. Nulla facilisi. Morbi eget lacus vel nunc tincidunt fringilla in quis tortor. Cras volutpat pretium sem a lobortis. Proin sit amet lacus eu metus rutrum tincidunt at non ligula. Donec eu tellus augue. Vestibulum vitae augue a nibh tincidunt mollis eget eu ipsum. Nam sollicitudin tincidunt nunc adipiscing molestie. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut in feugiat tellus. Phasellus ullamcorper vulputate pellentesque. Donec id lorem augue, a cursus lorem. Integer posuere blandit ornare. Nam venenatis suscipit pretium. Pellentesque a orci in arcu gravida tincidunt pulvinar eu libero. Integer a nisl non neque euismod suscipit nec a arcu.</td>
			</tr>
		</tbody>
	</table>
	<input id="talk" autofocus="autofocus" />
	<style type="text/css" id="style"></style>
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
