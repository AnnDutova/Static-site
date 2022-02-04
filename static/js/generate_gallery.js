async function generateElemInGalleryPage(){
    getData('http://localhost:8080/generateGalleryPage')
    .then((data) => {
        data.json().then(data =>{
            console.log(data);
            if (data != undefined){
                if (data.IsAuthorised != null){
                    if( data.IsAuthorised.IsAdministrator != null  || data.IsAuthorised.IsSeller != null  ){
                        document.getElementById('search_id').remove()
                        document.getElementById('user_id').remove()
                        document.getElementById('bucket_id').remove()
                    }else if (data.IsAuthorised.IsCustomer == null){
                        document.getElementById('search_id').remove()
                        document.getElementById('user_id').remove()
                        document.getElementById('bucket_id').remove()

                        var sign = document.createElement('a')
                        sign.classList.add('sign_button')
                        sign.id = "gallery_sign_up"
                        sign.textContent = "sign up"
                        var log = document.createElement('a')
                        log.classList.add('log_button')
                        log.id = "gallery_log_in"
                        log.textContent = "log in"

                        document.getElementById('right_header_corner').appendChild(sign)
                        document.getElementById('right_header_corner').appendChild(log)
                    }
                }else{
                    document.getElementById('search_id').remove()
                    document.getElementById('user_id').remove()
                    document.getElementById('bucket_id').remove()

                    var sign = document.createElement('a')
                    sign.classList.add('sign_button')
                    sign.id = "gallery_sign_up"
                    sign.textContent = "sign up"
                    var log = document.createElement('a')
                    log.classList.add('log_button')
                    log.id = "gallery_log_in"
                    log.textContent = "log in"

                    document.getElementById('right_header_corner').appendChild(sign)
                    document.getElementById('right_header_corner').appendChild(log)
                }
                console.log(data.Tracks.Music)

                if (data.Tracks != null){
                    for (var i = 0; i< data.Tracks.Music.length; i++){
                        var t = document.createElement('a');
                        t.classList.add('card_href','d-flex', 'justify-content-start');
                        t.appendChild(document.createElement('img'));
                        t.querySelector('img').classList.add('gallery_card_img','d-flex', 'mr-3', 'col-lg-3');
                        t.querySelector('img').id = 'card_image_' + i;
                        t.appendChild(document.createElement('div'));
                        t.querySelector('div').classList.add('media-body');
                        t.querySelector('div').appendChild(document.createElement('h5'));
                        t.querySelector('div').querySelector('h5').textContent = data.Tracks.Music[i].Song;
                        t.querySelector('div').querySelector('h5').id = "song";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[0].textContent = data.Tracks.Music[i].Author;
                        t.querySelector('div').querySelectorAll('p')[0].id = "author";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[1].textContent = data.Tracks.Music[i].Salon;
                        t.querySelector('div').querySelectorAll('p')[1].id = "salon";
                        let s = data.Tracks.Music[i].Song
                        let a = data.Tracks.Music[i].Author
                        let sl = data.Tracks.Music[i].Salon
                        t.onclick = function(s,a,sl) {return function() { openCompositionCard(s,a,sl); }}(s,a,sl);
                        document.getElementById('gallery_track_inner').appendChild(t);
                        document.getElementById("card_image_"+i).src="/static/images/divide.jpg";
                            
                    }
                }

                console.log(data.Collections)
                if (data.Collections.Collection != null){
                    for (var i = 0; i< data.Collections.Collection.length; i++){
                        var t = document.createElement('div');
                        t.classList.add('collection', 'collection_main_card')
                        t.appendChild(document.createElement('a'));
                        t.querySelector('a').classList.add('d-flex', 'justify-content-start', "align-items-center", 'collection_header');
                        t.querySelector('a').appendChild(document.createElement('img'));
                        t.querySelector('a').querySelector('img').classList.add('d-flex', 'mr-3', 'col-lg-3', 'collection_card_image')
                        t.querySelector('a').querySelector('img').src = "/static/images/equils.jpg"
                        t.querySelector('a').appendChild(document.createElement('div'));
                        t.querySelector('a').querySelector('div').classList.add('media-body')
                        t.querySelector('a').querySelector('div').appendChild(document.createElement('h5'));
                        t.querySelector('a').querySelector('div').querySelector('h5').textContent = data.Collections.Collection[i].Title
                        t.querySelector('a').querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('a').querySelector('div').querySelector('p').textContent = data.Collections.Collection[i].Owner
                        for (var j = 0; j< data.Collections.Collection[i].Collection.length; j++){
                            t.appendChild(document.createElement('div'))
                            t.querySelectorAll('div')[j+1].classList.add('extra_track', 'd-flex', 'justify-content-between', 'align-items-center')
                            t.querySelectorAll('div')[j+1].appendChild(document.createElement('p'))
                            t.querySelectorAll('div')[j+1].appendChild(document.createElement('p'))
                            t.querySelectorAll('div')[j+1].querySelectorAll('p')[0].textContent = data.Collections.Collection[i].Collection[j].Title
                            t.querySelectorAll('div')[j+1].querySelectorAll('p')[1].textContent = data.Collections.Collection[i].Collection[j].Artist
                        }
    
                        var last = t.appendChild(document.createElement('div'))
                        last.classList.add('d-flex', 'justify-content-end', 'align-items-center')
                        last.appendChild(document.createElement('p'))
                        last.querySelector('p').style.margin = "0 10px"
                        last.querySelector('p').textContent = data.Collections.Collection[i].Price + "x" + data.Collections.Collection[i].Count
                        document.getElementById('all_collections_inner').appendChild(t)
                    }
                }
            }
        })
    });
}