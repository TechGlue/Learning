var input = document.getElementById("input");
var output = document.getElementById("output");

//api for creating and managing a websocket connection to a server, as well as for sending and receiving data https://developer.mozilla.org/en-US/docs/Web/A:wPI/WebSocket

var socket = new WebSocket("ws://localhost/ws:8080/todo");

socket.onopen = function(){
    output.innerHTML += "Status: Connected\n";
};

socket.onmessage = function (e){
    output.innerHTML += "\nServer: " + e.data + "\n";
};

function send() {
      socket.send(input.value);
      input.value = "";
      console.log("send function was clicked");
}

const el = document.getElementById("btn");
el.addEventListener("click", send, false);

