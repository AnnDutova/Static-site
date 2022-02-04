async function generateSelectors(){
    getData('http://localhost:8080/getArtistGenres')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (data.Artists != null){
                for (var i = 0; i < data.Artists.length; i++){
                    var t = document.createElement('option');
                    t.value = i 
                    t.textContent = data.Artists[i]
                    document.getElementById('artists_inner').appendChild(t);
                }
            }
            if (data.Genre != null){
                for (var i = 0; i < data.Genre.length; i++){
                    var t = document.createElement('option');
                    t.value = i 
                    t.textContent = data.Genre[i]
                    document.getElementById('genre_inner').appendChild(t);
                }
            }
        })
    })
}