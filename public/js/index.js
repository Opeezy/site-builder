function change() {
    let range = document.getElementById("zoom-range")
    let rangeLabel = document.getElementById("range-label")

    rangeLabel.innerText = (range.value*10) + '%'
}

let pages = document.getElementsByClassName("main-page")
for (page of pages) {
    page.addEventListener("click", function() {
    })
}



