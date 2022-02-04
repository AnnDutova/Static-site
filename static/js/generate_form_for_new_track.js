var openFormForTrack = true
async function crateFormForGenerationNewMusic(){
    getData('http://localhost:8080/getArtistGenres')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (openFormForTrack){
                var f = document.createElement('form')
                f.classList.add('myForm', 'col-lg-4', 'col-md-5', 'col-sm-6')
                f.onsubmit = "createMusicCard(event)"
                f.id = "add_music_form"
                var title = f.appendChild(document.createElement('div'))
                title.classList.add('form-floating', 'm-2')
                title.appendChild(document.createElement('input'))
                title.appendChild(document.createElement('label'))
                title.querySelector('input').type = "text"
                title.querySelector('input').classList.add('form-control', 'form_style')
                title.querySelector('input').id = "title"
                title.querySelector('input').placeholder =  "title"
                title.querySelector('label').for = "floatingInput"
                title.querySelector('label').style.opacity = "0.4"
                title.querySelector('label').textContent = 'Composition title'

                var duration = f.appendChild(document.createElement('div'))
                duration.classList.add('form-floating', 'm-2')
                duration.appendChild(document.createElement('input'))
                duration.appendChild(document.createElement('label'))
                duration.querySelector('input').type = "text"
                duration.querySelector('input').classList.add('form-control', 'form_style')
                duration.querySelector('input').id = "duration"
                duration.querySelector('input').placeholder =  "duration"
                duration.querySelector('label').for = "floatingInput"
                duration.querySelector('label').style.opacity = "0.4"
                duration.querySelector('label').textContent = 'Duration of Song'

                var artist = f.appendChild(document.createElement('div'))
                artist.classList.add('form-floating', 'm-2')
                artist.appendChild(document.createElement('p'))
                artist.querySelector('p').textContent = "Artist"
                artist.appendChild(document.createElement('select'))
                artist.querySelector('select').id = "artists_inner"

                var genre = f.appendChild(document.createElement('div'))
                genre.classList.add('form-floating', 'm-2')
                genre.appendChild(document.createElement('p'))
                genre.querySelector('p').textContent = "Artist"
                genre.appendChild(document.createElement('select'))
                genre.querySelector('select').id = "genre_inner"

                var count = f.appendChild(document.createElement('div'))
                count.classList.add('form-floating', 'm-2')
                count.appendChild(document.createElement('input'))
                count.appendChild(document.createElement('label'))
                count.querySelector('input').type = "number"
                count.querySelector('input').classList.add('form-control', 'form_style')
                count.querySelector('input').id = "count"
                count.querySelector('input').placeholder =  "count"
                count.querySelector('label').for = "floatingInput"
                count.querySelector('label').style.opacity = "0.4"
                count.querySelector('label').textContent = 'Count'

                var price = f.appendChild(document.createElement('div'))
                price.classList.add('form-floating', 'm-2')
                price.appendChild(document.createElement('input'))
                price.appendChild(document.createElement('label'))
                price.querySelector('input').type = "number"
                price.querySelector('input').classList.add('form-control', 'form_style')
                price.querySelector('input').id = "price"
                price.querySelector('input').placeholder =  "price"
                price.querySelector('label').for = "floatingInput"
                price.querySelector('label').style.opacity = "0.4"
                price.querySelector('label').textContent = 'Price'

                var sale = f.appendChild(document.createElement('div'))
                sale.classList.add('form-floating', 'm-2')
                sale.appendChild(document.createElement('input'))
                sale.appendChild(document.createElement('label'))
                sale.querySelector('input').type = "number"
                sale.querySelector('input').classList.add('form-control', 'form_style')
                sale.querySelector('input').id = "sale"
                sale.querySelector('input').placeholder =  "sale"
                sale.querySelector('label').for = "floatingInput"
                sale.querySelector('label').style.opacity = "0.4"
                sale.querySelector('label').textContent = 'Sale'

                var btn = f.appendChild(document.createElement('button'))
                btn.classList.add('w-50%','btn', 'btn-lg', 'btn-primary', 'create_music_btn')
                btn.type="submit"
                btn.textContent = "Create Music card"
                btn.onclick = function() {return function() { createMusicCard(); }}();

                if (data.Artists != null){
                    for (var i = 0; i < data.Artists.length; i++){
                        var t = document.createElement('option');
                        t.value = i 
                        t.textContent = data.Artists[i]
                        artist.querySelector('select').appendChild(t);
                    }
                }
                if (data.Genre != null){
                    for (var i = 0; i < data.Genre.length; i++){
                        var t = document.createElement('option');
                        t.value = i 
                        t.textContent = data.Genre[i]
                        genre.querySelector('select').appendChild(t);
                    }
                }

                document.getElementById('create_new_music_card').appendChild(f)
                document.getElementById('create_new_music_card').style.backgroundColor = "white"
                openFormForTrack = false 
            }else{
                document.getElementById('add_music_form').remove()
                document.getElementById('create_new_music_card').style = null
                openFormForTrack = true
            }
        })
    })
}