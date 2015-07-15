package template

var Template = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>{{.Title}}</title>

    <link href="{{.Style}}" rel="stylesheet">

    <style>
        body { margin: 75px }
        ul { list-style-type: none; }
    </style>
  </head>

  <body role="document">

    <!-- Fixed navbar -->
    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <a class="navbar-brand" href="#">{{.Title}}</a>
        </div>
      </div>
    </nav>

    <div class="container-fluid" role="main">
        <div class="row">
            <div class="col-md-3">
                <h3>Navigation</h3>
                <hr/>
                {{ $path := .Path }}
                {{ range $dir, $files := .Files }}
                <h5>{{$dir}}</h5>
                <ul>
                    {{ range $i, $file := $files }}
                    {{ $this := print $dir $file }}
                    <li{{ if eq $path $this }} class="active"{{ end }}><a href="{{$dir}}{{$file}}">{{$file}}</a></li>
                    {{ end }}
                </ul>
                {{ end }}
            </div>
            <div class="col-md-9">
                {{.Body}}
            </div>
        </div>
    </div> <!-- /container -->

    <hr>
    <footer>
        <div class="text-center">
            built by <a href="http://github.com/jmervine/getdown">gitdown</a>
        </div>
    </footer>
  </body>
</html>
`
