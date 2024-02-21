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

function getChildren(element) {
    if (element.hasChildNodes()) {
        console.log(element.children)
    }
}

class EmbedData {
    constructor(parentElement) {
        this.parent = parentElement
        this.embedJson = this.fillData(this.parent)
    }

    fillData(div) {
        var tag = {}
        tag['class']=div.className
        tag['children'] = []
        for(var i = 0; i< div.children.length;i++){
            tag['children'].push(this.fillData(div.children[i]))
        }
        tag['children'] = tag['children'].flat(Infinity)
        return tag
    }
}

document.getElementById("save-btn").addEventListener("click", async function() {
    let designBody = document.getElementById("designer-body")
    let embedData = new EmbedData(designBody)
    await fetch('/save', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(embedData.embedJson)
    })
       .then(response => response.json())
       .then(response => console.log(JSON.stringify(response)))
})     



