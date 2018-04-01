socket = null;
name = "";
mainDiv = null;

// Initialisation code
$(function(){
    // get reference to the div where you put the messages
    mainDiv = $("#maindiv");
    console.log(mainDiv);
})

// called when you click connect
function connectToServer(){
    name = $("#name").val();
    var ipAddress = $("#ip").val();
    var port = $("#port").val();
    // alert("Connecting to " + ipAddress + " " + port);

    var address = "ws://" + ipAddress + ":" + port;
    if (socket && socket.readyState == socket.OPEN){
        alert("Already connected to the server!")
    }
    else{
        socket = new WebSocket(address);

        socket.onopen = function notifyClient(event){
            console.log("Connected!")
        }

        socket.onerror = function connectionProblem(event){
            alert("couldn't connect to "+ipAddress+":"+port)
        }

        socket.onmessage = handleMessage;
    }
}

// called when you click the send button
function sendMessage(){
    if (socket && socket.readyState == socket.OPEN){
        message = {
            name: name,
            message: $("#message").val()
        };
        socket.send(JSON.stringify(message));
    }
    else {
        alert("Your socket isn't connected to anything");
    }

}


// event handler for when the server broadcasts a message to you
function handleMessage(event){
    data = JSON.parse(event.data);
    createMessageElements(data);
    console.log(data.name+": "+data.message);
}


// creatae the html with the images and stuff
function createMessageElements(data){
    var container = $("<div></div>");
    var name = $("<p>"+data.name+":</p>");
    var message = $("<p>"+data.message+"</p>");

    container.append(name);
    container.append(message);
    mainDiv.append(container)


}
