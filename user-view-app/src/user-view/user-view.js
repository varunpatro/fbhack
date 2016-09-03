Polymer({
    is: 'user-view',
    properties: {
        uuid: {
            type: String,
            value: "",
            notify: true
        },
        isSafe: {
            type: Boolean,
            value: false,
            notify: true
        },
        safeResponse: {
            type: Object,
            value: {},
            notify: true,
        },
        unsafeResponse: {
            type: Object,
            value: {},
            notify: true,
        },
        showMap: {
            type: Boolean,
            value: false,
            notify: true
        },
        latitude: {
            type: Number,
            value: 1.2956504,
            notify: true
        },
        longtitude: {
            type: Number,
            value: 103.8566111,
            notify: true
        },
        updateLocSuccess: {
          type: Boolean,
          value: false,
          notify: true
        }
    },
    ready: function() {
        this.uuid = this.getUuid();
        console.log(this.uuid);
    },
    getUuid: function() {
		url = window.location.href;
		name = ("uuid").replace(/[\[\]]/g, "\\$&");
		var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
			results = regex.exec(url);
		if (!results) return null;
		if (!results[2]) return '';
		return decodeURIComponent(results[2].replace(/\+/g, " ")); 
    },
    getParams: function(uuid) {
        return {uuid: uuid};
    },
    markAsSafe: function() {
        this.isSafe = true;
        if (this.uuid && this.uuid !== "") {
            document.querySelector('#safe-ajax').generateRequest();
            console.log(this.safeResponse);
        }
    },
    markAsUnsafe: function() {
        this.isSafe = false;
        this.showMap = true;
        if (this.uuid && this.uuid !== "") {
            document.querySelector('#unsafe-ajax').generateRequest();
            console.log(this.unsafeResponse);
        }
        this.triggerLocationService()
    },
    triggerLocationService: function(){
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(function(position){
                this.latitude = position.coords.latitude;
                this.longtitude = position.coords.longtitude;
            });
        }
    },
    callForHelp: function() {
        if(this.uuid && this.uuid !== "") {
            document.querySelector('#location-ajax').generateRequest();
        }
      this.updateLocSuccess = true;
    },
    getLocationParams: function(uuid){
        return {
            uuid: uuid,
            latitude: this.latitude,
            longtitude: this.longtitude,
            timestamp: Date.now()
        }
    }
});
