var mouseInEmbed = false

document.getElementById("embed").addEventListener("mouseenter", function() {mouseInEmbed=true;console.log(mouseInEmbed)})
document.getElementById("embed").addEventListener("mouseout", function() {mouseInEmbed=false;console.log(mouseInEmbed)})

function change() {
    let range = document.getElementById("zoom-range")
    let rangeLabel = document.getElementById("range-label")

    rangeLabel.innerText = (range.value*10) + '%'
}