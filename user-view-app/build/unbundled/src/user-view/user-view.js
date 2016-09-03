Polymer({is:"user-view",properties:{uuid:{type:String,value:"",notify:!0},isSafe:{type:Boolean,value:!1,notify:!0},safeResponse:{type:Object,value:{},notify:!0},unsafeResponse:{type:Object,value:{},notify:!0},showMap:{type:Boolean,value:!1,notify:!0},showZikaMap:{type:Boolean,value:!1,notify:!0},latitude:{type:Number,value:1.2956504,notify:!0},longtitude:{type:Number,value:103.8566111,notify:!0},updateLocSuccess:{type:Boolean,value:!1,notify:!0}},listeners:{"google-map-ready":"addCircles"},ready:function(){this.uuid=this.getUuid(),console.log(this.uuid)},getUuid:function(){url=window.location.href,name="uuid".replace(/[\[\]]/g,"\\$&");var e=new RegExp("[?&]"+name+"(=([^&#]*)|&|#|$)"),t=e.exec(url);return t?t[2]?decodeURIComponent(t[2].replace(/\+/g," ")):"":null},getParams:function(e){return{uuid:e}},markAsSafe:function(){this.isSafe=!0,this.uuid&&""!==this.uuid&&(document.querySelector("#safe-ajax").generateRequest(),console.log(this.safeResponse))},markAsUnsafe:function(){this.isSafe=!1,this.showMap=!0,this.uuid&&""!==this.uuid&&(document.querySelector("#unsafe-ajax").generateRequest(),console.log(this.unsafeResponse)),this.triggerLocationService()},goToZikaView:function(){this.showZikaMap=!0},triggerLocationService:function(){navigator.geolocation&&navigator.geolocation.getCurrentPosition(function(e){this.latitude=e.coords.latitude,this.longtitude=e.coords.longtitude})},callForHelp:function(){this.uuid&&""!==this.uuid&&document.querySelector("#location-ajax").generateRequest(),this.updateLocSuccess=!0},getLocationParams:function(e){return{uuid:e,latitude:this.latitude,longtitude:this.longtitude,timestamp:Date.now()}},addCircles:function(){zikaMap=document.querySelector("#zika-map");var e={lat:1.316566,lng:103.88294},t=(new window.google.maps.Circle({strokeColor:"#FF0000",strokeOpacity:.8,strokeWeight:2,fillColor:"#FF0000",fillOpacity:.35,map:zikaMap.map,center:e,radius:2500}),{lat:1.33733,lng:103.93392}),o=(new window.google.maps.Circle({strokeColor:"#FF0000",strokeOpacity:.8,strokeWeight:2,fillColor:"#FF0000",fillOpacity:.35,map:zikaMap.map,center:t,radius:1e3}),{lat:1.430087,lng:103.83556});new window.google.maps.Circle({strokeColor:"#FF0000",strokeOpacity:.8,strokeWeight:2,fillColor:"#FF0000",fillOpacity:.35,map:zikaMap.map,center:o,radius:1e3})}});