'use strict';

console.log('here')

var CastAlotPlayer = function() {
};

CastAlotPlayer.prototype.initializeCastApi = function() {

    console.log("initialising..")

    cast.framework.CastContext.getInstance().setOptions({
        receiverApplicationId: chrome.cast.media.DEFAULT_MEDIA_RECEIVER_APP_ID,
        autoJoinPolicy: chrome.cast.AutoJoinPolicy.ORIGIN_SCOPED
    });

    this.setupPlayer();
  
};

CastAlotPlayer.prototype.setupPlayer = function() {

    var currentMediaURL = 'http://192.168.0.2:4444/f/Hitman.Agent.47.2015.720p.BRRip.x264.AAC-ETRG.mp4';
    var contentType = 'video/mp4';
    
    var player = new cast.framework.RemotePlayer();
    var playerController = new cast.framework.RemotePlayerController(player);

    playerController.addEventListener(
        cast.framework.RemotePlayerEventType.ANY_CHANGE, function(event) {
            console.log(event);         
    });

    document.getElementById("play").onclick = function() {
        playerController.playOrPause()
    };
    
    this.playMedia(currentMediaURL, contentType)
    
};


CastAlotPlayer.prototype.playMedia = function(media, type) {

    console.log("trying to play "+media)

    var mediaInfo = new chrome.cast.media.MediaInfo(media, type);
    var request = new chrome.cast.media.LoadRequest(mediaInfo);

    var castSession = cast.framework.CastContext.getInstance().getCurrentSession();

    if (!castSession) {
        console.log("cast session not available")
        return;
    }

    castSession.loadMedia(request).then(
        function() { console.log('Load succeed'); },
        function(errorCode) { console.log('Error code: ' + errorCode); 
    });



};

window.CastAlotPlayer = CastAlotPlayer;



