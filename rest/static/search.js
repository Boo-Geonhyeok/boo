const button = document.querySelectorAll("#addButton")
const searchBar = document.querySelector("#songTitle")

for(var i = 0; i<button.length; i++) {
    button[i].addEventListener("click", handleAddButtonClick);
}

searchBar.value = searchBar.placeholder

function handleAddButtonClick(event) {
    // event.preventDefault()
    const songData = String(event.target.offsetParent.innerText).slice(0,-3)
    console.log(songData);
    const before = songData.split('⫘')[0].slice(0,-1)
    console.log(before);
    const after = songData.split('⫘')[1].slice(1,)
    console.log(after);

    const songObj = {
        title: before,
        artist: after
    }

    fetch('', {
        method: 'post',
        body: JSON.stringify(songObj),
        headers: {'Content-Type': 'application/json'}
    }).then(function (response){
        return response.text();
    }).then(function (text) {
        console.log(text);
    }).catch(function (error) {
        console.error(error);
    })
}

