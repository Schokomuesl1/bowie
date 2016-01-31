<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>DSA 5 Heldengenerator</title>

    <!-- Bootstrap -->
    <link href="static/bootstrap/css/bootstrap.min.css" rel="stylesheet">
    <link href="static/bootstrap/css/bootstrap-theme.min.css" rel="stylesheet">
    <link href="static/css/normalize.css" rel="stylesheet">

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
     <!-- Fixed navbar -->
    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container">

        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav">
            <li id="nav_Neu" class="active"><a href="javascript:jumpToPage('Neu')">Neu</a></li>
            <li id="nav_Allgemeines"><a href="javascript:jumpToPage('Allgemeines')">Allgemeines</a></li>
            <li id="nav_Kampftechniken"><a href="javascript:jumpToPage('Kampftechniken')">Kampftechniken</a></li>
            <li id="nav_Talente"><a href="javascript:jumpToPage('Talente')">Talente</a></li>
          </ul>
        </div><!--/.nav-collapse -->
      </div>
    </nav>
	<footer id="page_footer" class="footer" style="padding-top: 50px">
      
    </footer>
    <div id="main_window" class="container-fluid" role="main" style="padding-top: 55px">
    <input type="button" class="btn btn-primary" value="Neuen Held erstellen" onClick="Javascript:jumpToPage('Neu')"/>
	</div>
	<script src = "static/js/app.js"> </script>
    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/bootstrap/js/bootstrap.min.js"></script>
  </body>
</html>