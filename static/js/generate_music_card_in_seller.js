var open = true
async function generateAllMusic(){
    getData('http://localhost:8080/collectMusicinSeller')
    .then((data) => {
        data.json().then(data =>{
            console.log(data);
            if (data.Music != null){
                if (open){
                    for (var i = 0; i< data.Music.length; i++){
                        var t = document.createElement('a');
                        t.classList.add('card_music','d-flex', 'justify-content-start');
                        t.appendChild(document.createElement('img'));
                        t.querySelector('img').classList.add('card_music_image','d-flex', 'mr-3', 'col-lg-3');
                        t.querySelector('img').id = 'card_image_' + i;
                        t.appendChild(document.createElement('div'));
                        t.querySelector('div').classList.add('media-body');
                        t.querySelector('div').appendChild(document.createElement('h5'));
                        t.querySelector('div').querySelector('h5').textContent = data.Music[i].Song;
                        t.querySelector('div').querySelector('h5').id = "song";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[0].textContent = data.Music[i].Author;
                        t.querySelector('div').querySelectorAll('p')[0].id = "author";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[1].textContent = data.Music[i].Salon;
                        t.querySelector('div').querySelectorAll('p')[1].id = "salon";
                        let s = data.Music[i].Song
                        let a = data.Music[i].Author
                        let sl = data.Music[i].Salon
                        t.onclick = function(s,a,sl) {return function() { openCompositionCardFromSeller(s,a,sl); }}(s,a,sl);
                        document.getElementById('music_inner').appendChild(t);
                        document.getElementById("card_image_"+i).src="/static/images/divide.jpg";    
                    }
                    open = false
            }else{
                document.getElementById('music_inner').remove();
                document.getElementById('seller_music').appendChild(document.createElement('div'))
                var new_inner = document.getElementById('seller_music').querySelector('div')
                new_inner.classList.add('d-flex', 'flex-wrap')
                new_inner.id = "music_inner"
                open = true
            }
            }
        })
    });
}