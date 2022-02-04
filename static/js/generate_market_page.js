var start = 0
var end = 0
var already_in_page = 0

var start_music = 0
var end_music = 0
var already_in_page_music = 0

async function generateMarketPage(){
    getData('http://localhost:8080/generateMarketPage')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (data != undefined){
                if (data.IsAuthorised != undefined){
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
                if (data.Genres != null){
                    for (var i = 0; i< data.Genres.length; i++){
                        var t = document.createElement('div')
                        t.classList.add('d-flex', 'align-items-center', 'justify-content-start')
                        t.appendChild(document.createElement('input'))
                        t.querySelector('input').type = "checkbox"
                        t.querySelector('input').value =  data.Genres[i]
                        t.appendChild(document.createElement('p'))
                        t.querySelector('p').style.margin = "0 5px"
                        t.querySelector('p').textContent = data.Genres[i]
                        document.getElementById('filter_music_genre').appendChild(t);
                    }
                }
                var t = document.createElement('p')
                t.textContent = data.MinPrice + "-"+ data.MaxPrice
                document.getElementById('price_range_slider').appendChild(t);
                document.getElementById('myRange').min = data.MinPrice
                document.getElementById('myRange').value = data.MinPrice
                document.getElementById('myRange').max = data.MaxPrice
                document.getElementById('answer').value = data.MinPrice

                if (data.Collections.Collection != null){
                    if (start > 0){
                        document.getElementById('see_more_btn_collection').remove()
                    }
                    if (data.Collections.Collection.length > 2){
                        start = already_in_page
                        end = already_in_page + 2
                    }
                    for (var i = start; i< end; i++){
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
                            var main_div = t.appendChild(document.createElement('div'))
                            main_div.classList.add('extra_track', 'd-flex', 'align-items-center')
                            main_div.appendChild(document.createElement('div'))
                            main_div.querySelector('div').classList.add('d-flex', 'flex-row', 'justify-content-between')
                            main_div.querySelector('div').style.width = "100%"
                            main_div.querySelector('div').style.margin = " 0 10px 0 0"
                            main_div.querySelector('div').appendChild(document.createElement('p'))
                            main_div.querySelector('div').appendChild(document.createElement('p'))
                            main_div.querySelector('div').querySelectorAll('p')[0].textContent = data.Collections.Collection[i].Collection[j].Title
                            main_div.querySelector('div').querySelectorAll('p')[1].textContent = data.Collections.Collection[i].Collection[j].Artist
                            main_div.appendChild(document.createElement('button'))
                            main_div.querySelector('button').classList.add('w-50%', 'btn', 'btn-lg', 'btn-primary', 'buy_in_collection_card')
                            main_div.querySelector('button').type="submit"
                            main_div.querySelector('button').textContent="Buy"
                            let s = data.Collections.Collection[i].Collection[j].Title
                            let a = data.Collections.Collection[i].Collection[j].Artist
                            let sl = data.Collections.Collection[i].Owner
                            main_div.querySelector('button').onclick = function(s,a,sl) {return function() { addMusicCardToBucketFromCollection(s,a,sl); }}(s,a,sl);
                        }
    
                        var last = t.appendChild(document.createElement('div'))
                        last.classList.add('d-flex', 'justify-content-end', 'align-items-center')
                        last.appendChild(document.createElement('p'))
                        last.querySelector('p').style.margin = "0 10px"
                        last.querySelector('p').textContent = data.Collections.Collection[i].Price + "x" + data.Collections.Collection[i].Count
                        last.appendChild(document.createElement('button'))
                        last.querySelector('button').classList.add('w-50%','btn','btn-lg', 'btn-primary', 'buy_all_in_collection_card')
                        last.querySelector('button').type = "submit"
                        last.querySelector('button').textContent = "Buy all"
                        let title = data.Collections.Collection[i].Title
                        let s = data.Collections.Collection[i].Owner
                        last.querySelector('button').onclick = function(title,s) {return function() { addAllCollectionToBucket(title,s); }}(title,s);
                        document.getElementById('collections_inner').appendChild(t)
                    }
                    already_in_page = end
                    if (data.Collections.Collection.length > already_in_page){
                        btn = document.createElement('div');
                        btn.classList.add('d-flex', 'justify-content-center');
                        btn.appendChild(document.createElement('button'));
                        btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary', 'buy_all_in_collection_card');
                        btn.querySelector('button').type = "submit";
                        btn.querySelector('button').id = "see_more_btn_collection"
                        btn.querySelector('button').textContent = "See more";
                        btn.onclick = function() {return function() { openNewCollections(); }}();
                        document.getElementById('collections_inner').appendChild(btn)
                    }
                }

                if (data.Tracks != null){
                    if (start_music >= 0){
                        document.getElementById('see_more_btn_tracks').remove()
                    }
                    if (data.Tracks.Music.length > 2){
                        start_music = already_in_page_music
                        end_music = already_in_page_music + 5
                        if (end_music > data.Tracks.Music.length){
                            end_music = data.Tracks.Music.length
                        }
                    }
                    for (var i = start_music; i< end_music; i++){
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
                        let isAdmin = data.IsAuthorised.IsAdministrator
                        let s = data.Tracks.Music[i].Song
                        let a = data.Tracks.Music[i].Author
                        let sl = data.Tracks.Music[i].Salon
                        t.onclick = function(isAdmin, s,a,sl) {return function() { openCardFromMarket(isAdmin, s,a,sl); }}(isAdmin, s,a,sl);
                        document.getElementById('tracks_inner').appendChild(t);
                        document.getElementById("card_image_"+i).src="/static/images/divide.jpg";
                    }
                    already_in_page_music = end_music
                    if (data.Tracks.Music.length > already_in_page_music){
                        btn = document.createElement('div');
                        btn.classList.add('d-flex', 'justify-content-center');
                        btn.appendChild(document.createElement('button'));
                        btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary', 'buy_all_in_collection_card');
                        btn.querySelector('button').type = "submit";
                        btn.querySelector('button').id = "see_more_btn_tracks"
                        btn.querySelector('button').textContent = "See more";
                        btn.onclick = function() {return function() { openNewTracks(); }}();
                        document.getElementById('btn_see_more_tracks').appendChild(btn)
                    }
                }

            }

        })
    })
}


async function openNewCollections(){
    getData('http://localhost:8080/generateMarketPage')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (data.Collections.Collection != null){
                if (start >= 0){
                    document.getElementById('see_more_btn_collection').remove()
                }
                if (data.Collections.Collection.length > 2){
                    start = already_in_page
                    end = already_in_page + 2
                    if (end > data.Collections.Collection.length){
                        end = data.Collections.Collection.length
                    }
                }
                for (var i = start; i< end; i++){
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
                        var main_div = t.appendChild(document.createElement('div'))
                        main_div.classList.add('extra_track', 'd-flex', 'align-items-center')
                        main_div.appendChild(document.createElement('div'))
                        main_div.querySelector('div').classList.add('d-flex', 'flex-row', 'justify-content-between')
                        main_div.querySelector('div').style.width = "100%"
                        main_div.querySelector('div').style.margin = " 0 10px 0 0"
                        main_div.querySelector('div').appendChild(document.createElement('p'))
                        main_div.querySelector('div').appendChild(document.createElement('p'))
                        main_div.querySelector('div').querySelectorAll('p')[0].textContent = data.Collections.Collection[i].Collection[j].Title
                        main_div.querySelector('div').querySelectorAll('p')[1].textContent = data.Collections.Collection[i].Collection[j].Artist
                        main_div.appendChild(document.createElement('button'))
                        main_div.querySelector('button').classList.add('w-50%', 'btn', 'btn-lg', 'btn-primary', 'buy_in_collection_card')
                        main_div.querySelector('button').type="submit"
                        main_div.querySelector('button').textContent="Buy"
                        let s = data.Collections.Collection[i].Collection[j].Title
                        let a = data.Collections.Collection[i].Collection[j].Artist
                        let sl = data.Collections.Collection[i].Owner
                        main_div.querySelector('button').onclick = function(s,a,sl) {return function() { addMusicCardToBucketFromCollection(s,a,sl); }}(s,a,sl);
                    }

                    var last = t.appendChild(document.createElement('div'))
                    last.classList.add('d-flex', 'justify-content-end', 'align-items-center')
                    last.appendChild(document.createElement('p'))
                    last.querySelector('p').style.margin = "0 10px"
                    last.querySelector('p').textContent = data.Collections.Collection[i].Price + "x" + data.Collections.Collection[i].Count
                    last.appendChild(document.createElement('button'))
                    last.querySelector('button').classList.add('w-50%','btn','btn-lg', 'btn-primary', 'buy_all_in_collection_card')
                    last.querySelector('button').type = "submit"
                    last.querySelector('button').textContent = "Buy all"
                    console.log("Buy all " + data.Collections.Collection[i].Title+ " "+ data.Collections.Collection[i].Owner)
                    let title = data.Collections.Collection[i].Title
                    let s = data.Collections.Collection[i].Owner
                    last.querySelector('button').onclick = function(title,s) {return function() { addAllCollectionToBucket(title,s); }}(title,s);
                    document.getElementById('collections_inner').appendChild(t)
                }
                already_in_page = end
                if (data.Collections.Collection.length > already_in_page){
                    btn = document.createElement('div');
                    btn.classList.add('d-flex', 'justify-content-center');
                    btn.appendChild(document.createElement('button'));
                    btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary', 'buy_all_in_collection_card');
                    btn.querySelector('button').type = "submit";
                    btn.querySelector('button').id = "see_more_btn_collection"
                    btn.querySelector('button').textContent = "See more";
                    btn.onclick = function() {return function() { openNewCollections(); }}();
                    document.getElementById('collections_inner').appendChild(btn)
                }

            }
        })
    })
}

async function openNewTracks(){
    getData('http://localhost:8080/generateMarketPage')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)

            if (data.Tracks != null){
                if (start_music >= 0){
                    document.getElementById('see_more_btn_tracks').remove()
                }
                if (data.Tracks.Music.length > 2){
                    start_music = already_in_page_music
                    end_music = already_in_page_music + 5
                    if (end_music > data.Tracks.Music.length){
                        end_music = data.Tracks.Music.length
                    }
                }
                for (var i = start_music; i< end_music; i++){
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
                    let isAdmin = data.IsAuthorised.IsAdministrator
                    let s = data.Tracks.Music[i].Song
                    let a = data.Tracks.Music[i].Author
                    let sl = data.Tracks.Music[i].Salon
                    t.onclick = function(isAdmin, s,a,sl) {return function() { openCardFromMarket(isAdmin, s,a,sl); }}(isAdmin, s,a,sl);
                    document.getElementById('tracks_inner').appendChild(t);
                    document.getElementById("card_image_"+i).src="/static/images/divide.jpg";
                }
                already_in_page_music = end_music
                if (data.Tracks.Music.length > already_in_page_music){
                    btn = document.createElement('div');
                    btn.classList.add('d-flex', 'justify-content-center');
                    btn.appendChild(document.createElement('button'));
                    btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary', 'buy_all_in_collection_card');
                    btn.querySelector('button').type = "submit";
                    btn.querySelector('button').id = "see_more_btn_tracks"
                    btn.querySelector('button').textContent = "See more";
                    btn.onclick = function() {return function() { openNewTracks(); }}();
                    document.getElementById('btn_see_more_tracks').appendChild(btn)
                }
            }

        })
    })
}