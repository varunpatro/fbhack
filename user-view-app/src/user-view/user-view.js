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
        if (this.uuid && this.uuid !== "") {
            document.querySelector('#unsafe-ajax').generateRequest();
            console.log(this.unsafeResponse);
        }
    },
});