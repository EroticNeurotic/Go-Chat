// don't forget to do npm install --save ws
websocket = require ('ws');

PORT = 8080

// broadcast the data to all the connected clients
// except the sender
function broadcast(server, data, sender){
    server.clients.forEach(function (client){
        if (client.readyState == websocket.OPEN && client != sender){
            client.send(data);
        }

    });
}


const server = new websocket.Server({
    port: PORT,
    // I honestly don't remember what this option does
    perMessageDeflate:false
});


// create an event handler for a new connection
// I've created some anonymous functions on the go here
// which is why there are function definition inside functions
server.on('connection', function register(client){

    // assign the client an event handler for when they send a message
    client.on('message', function handleMessage(data){
        broadcast(server, data, client);
    });
})

server.on("listening", function ready(){
    console.log("server running on port " +PORT);
})
