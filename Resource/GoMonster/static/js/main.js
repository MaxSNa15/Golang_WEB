let imgArray = [
    "static/img/m-fire.png",
    "static/img/m-grass.png",
    "static/img/m-water.png",
]

function choose(x){
    fetch("/play?c=" + x)
    .then(response => response.json())
    .then(data => {
        if (x == 0){
            document.getElementById("player_choise").innerHTML = "El jugador eligio al monstro de fuego"
        }else if (x == 1){
            document.getElementById("player_choise").innerHTML = "El jugador eligio al monstro de agua"
        }else{
            document.getElementById("player_choise").innerHTML = "El jugador eligio al monstro de pasto"
        }

        document.getElementById("player_score").innerHTML = data.player_score;
        document.getElementById("computer_score").innerHTML = data.computer_score;

        document.getElementById("computer_choise").innerHTML = data.computer_choice;
        document.getElementById("round_result").innerHTML = data.round_result;
        document.getElementById("round_message").innerHTML = data.message;

        var imgElement = document.getElementById("img_computer");
        imgElement.src = imgArray[data.computer_choice_int];
    })
}