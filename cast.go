package main

import (
	"html/template"
	"log"
	"net/http"
)

func caster(w http.ResponseWriter, r *http.Request) {

	log.Println("running caster")
	log.Println(tpl)

	tmpl, err := template.New("name").Parse(tpl)

	if err != nil {
		log.Println(err)
	}

	data := struct{}{}

	// Error checking elided
	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
	}

}

var tpl = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Sir Cast-A-Lot</title>

<script type="text/javascript" src="https://www.gstatic.com/cv/js/sender/v1/cast_sender.js?loadCastFramework=1"></script>

<script type="text/javascript" src="/assets/cast.js"></script>

</head>

<body>
<button is="google-cast-button"></button>

<button id="play">Play/Pause</button>

<script type="text/javascript">
var Caster = new CastAlotPlayer();

window['__onGCastApiAvailable'] = function(loaded, errorInfo) {
  if (loaded) {
    Caster.initializeCastApi()

	var p = document.getElementById('play');
    p.onclick = function(){ Caster.playMedia('http://192.168.0.2:4444/f/Hitman.Agent.47.2015.720p.BRRip.x264.AAC-ETRG.mp4', 'video/mp4');};

  } else {
    console.log(errorInfo);
  }
}

</script>

</body>

</html>`
