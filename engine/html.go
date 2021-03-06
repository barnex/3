package engine

// THIS FILE IS AUTO GENERATED FROM gui.html
// EDITING IS FUTILE

const templText = `
<!DOCTYPE html>
<html>

<head>

	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

	<title>mumax3</title>

	<style media="all" type="text/css">

		body  { margin-left: 5%; margin-right:5%; font-family: sans-serif; font-size: 14px; }
		h1    { font-size: 22px; color: gray; }
		h2    { font-size: 18px; color: gray; }
		img   { margin: 10px; }
		table { border-collapse: collapse; }
		tr:nth-child(even) { background-color: white; }
		tr:nth-child(odd)  { background-color: #F7F7FF; }
		td        { padding: 1px 5px; }
		hr        { border-style: none; border-top: 1px solid #CCCCCC; }
		a         { color: #375EAB; text-decoration: none; }
		div       { margin-left: 20px; margin-top: 5px; margin-bottom: 20px; }
		div#footer{ color:gray; font-size:14px; border:none; }
		.ErrorBox { color: red; font-weight: bold; } 
		.TextBox  { border:solid; border-color:#BBBBBB; border-width:1px; padding-left:4px; }
	</style>

	<script>
		function toggle(id) {
	       var el = document.getElementById(id);
	       if(el.style.display != 'none'){
	          el.style.display = 'none';
			} else {
	          el.style.display = 'block';
			}
	    }
	</script>

	{{.JS}}

</head>


<body>

	<h1> mumax<sup>3</sup> web interface </h1> 
		<p> {{.ErrorBox}} </p>
		<p> {{.Span "log"}} </p>
	<hr/>

	
	<h2> geometry </h2><div>

		<table>
			<tr> <td>gridsize: </td> <td>{{.Span "nx"}} </td> <td> &times; {{.Span "ny"}}</td> <td> &times; {{.Span "nz"}}                </td> </tr>
			<tr> <td>cellsize: </td> <td>{{.Span "cx"}} </td> <td> &times; {{.Span "cy"}}</td> <td> &times; {{.Span "cz"}} nm<sup>3</sup> </td> </tr>
			<tr> <td>worldsize:</td> <td>{{.Span "wx"}} </td> <td> &times; {{.Span "wy"}}</td> <td> &times; {{.Span "wz"}} nm<sup>3</sup> </td> </tr>
		</table>
	</div><hr/>


	<h2> solver </h2><div>

		<table>
			<tr style="background-color:white"> <td>
	
		<table>
			<tr> <td> {{.Button "break"}}</td> <td> status: {{.Span "solverstatus" "initializing"}} </td></tr>
			<tr> <td> {{.Button "run"}}  </td> <td> {{.NumBox "runtime" 1e-9}}s</td></tr> 
			<tr> <td> {{.Button "steps"}}</td> <td> {{.IntBox "runsteps" 1000}}</td></tr>
		</table>

		</td><td>
			&nbsp; &nbsp; &nbsp; &nbsp;
		</td><td>

		<table>
			<tr> <td>step:    </td><td>{{.Span   "step"}}     </td><td>time:  </td><td>{{.Span "time"}}s          </td></tr>
			<tr> <td>dt:      </td><td>{{.Span   "dt"}}s      </td><td>fixdt: </td><td>{{.NumBox "fixdt"  0}}s    </td></tr>
			<tr> <td>mindt:   </td><td>{{.NumBox "mindt" 0}}s </td><td>maxdt: </td><td>{{.NumBox "maxdt"  0}}s    </td></tr>
			<tr> <td>err/step:</td><td>{{.Span   "lasterr"}}  </td><td>maxerr:</td><td>{{.NumBox "maxerr" 0}}/step</td></tr>
		</table>

			</td></tr>
		</table>
	</div><hr/>


	<h2> parameters </h2><div>

	<p class=ErrorBox>{{.Span "paramErr" ""}}</p>

	<table>
	<tr> <td> <b>Region </b> </td>
	<td>{{.BeginSelect "sel_region"}}
		{{range .Data.MakeRange 0 256}}
			<option value= {{.}}> {{.}}</option>
		{{end}}
	{{.EndSelect}}</td></tr>

	{{range $k,$v := .Data.Params}}
		<tr> <td>{{$k}}</td> {{range $.Data.CompBoxIds $k}} <td>{{$.TextBox . "0"}}</td> {{end}} <td> {{$v.Unit}}</td> </tr>
	{{end}}
	</table>

	</div><hr/>


	<h2> display </h2><div>

		{{.BeginSelect "sel_render"}}
			{{range $k,$v := .Data.Quants}} {{$.Option $k}} {{end}}
		{{.EndSelect}}
		<br/>
		{{.Img "render" "/render/m"}}

	</div><hr/>


	<h2> gnuplot </h2><div>

		Plot of "table.txt", provided table is being autosaved and gnuplot installed.<br>
		<b>plot "table.txt" using {{.IntBox "usingX" 1}} : {{.IntBox "usingY" 2}} with lines </b><br/>
		<p class=ErrorBox>{{.Span "plotErr"}}</p>
		{{.Img "plot" "/plot/"}}

	</div><hr/>


	<h2> process </h2><div>
		<table>
			<tr> <td>host:     </td> <td>{{.Span "hostname"}} </td>  </tr>
			<tr> <td>gpu:      </td> <td>{{.Span "gpu"}}      </td>  </tr>
			<tr> <td>walltime: </td> <td>{{.Span "walltime"}} </td>  </tr>
		</table>
	</div><hr/>

	<div id="footer">
		<br/>
		<center>
			mumax 3<br/>
			Copyright 2012-2013 <a href="mailto:a.vansteenkiste@gmail.com">Arne Vansteenkiste</a>,
			<a href="http://dynamat.ugent.be">Dynamat group</a>, Ghent University, Belgium.<br/>
			You are free to modify and distribute this software under the terms of the
			<a href="http://www.gnu.org/licenses/gpl-3.0.html">GPLv3 license</a>.
		</center>
	</div>

</body>
</html>
`
