Vue.component('message', { //template not used yet, want to find out how to implement
	props: ['username', 'content'],
	template: `

    <div style="width:60%">
        <div class="w3-container w3-green">
            {{username}}
        </div>
        <div class="w3-container w3-white w3-border">
            {{content}}
        </div>
    </div>

    `
});

new Vue({
	el: '#messageBox',

	data: {
		ws: null,
		port: 9999,
		newMsg: '',
		chatContent: '',
		username: 'Dave', //create login to set this yourself
		joined: true //should depend on login status
	},

	created: function() {
		console.log('Instance created');
		var self = this;
		var ip = prompt("Enter IP of websocket","localhost")
		this.ws = new WebSocket('ws://' + ip+":"+this.port+'/ws');

		this.ws.onopen = function success(event){
			console.log('WebSocket initialised')
		}
		this.ws.onerror = function handleError(){
			alert("couldn't connect to WebSocket")
		}

		this.ws.addEventListener('message', function(e) {

			console.log('message received');
			var msg = JSON.parse(e.data);
			console.log('Message: ' + msg.content);

			self.chatContent +=
				'<div style="width:60%; margin-bottom:20px">' +
				'<div class="w3-container w3-green">' +
				msg.username +
				'</div>' +
				'<div class="w3-container w3-white w3-border">' +
				msg.content +
				'</div>' +
				'</div>';

			console.log(self.chatContent);
		})
	},

	methods: {
		send: function() {
			console.log('trying to send');
			if (this.newMsg != '') {
				this.ws.send(
					JSON.stringify({
						username: this.username,
						content: this.newMsg
					})
				);
				this.newMsg = '';
			}
		}
	},
})
