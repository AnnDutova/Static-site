var start_music = 0
var already_in_page_music = 0
var end_music = 0

var start_collection = 0
var already_in_page_collection = 0
var end_collection = 0

async function generateElemInUserPage(){
    getData('http://localhost:8080/collectInformation')
    .then((data) => {
        data.json().then(data =>{
            console.log(data);
            if (data.Music != null){
                if (data.Music.Music.length > 0){
                    if (start_music == 0){
                        document.getElementById('empty_tracks_message').remove();
                        document.getElementById('buy_tracks_empty').remove();
                    }
                    else if(start_music > 0){
                        document.getElementById('see_more_btn_music').remove()
                    }
                    
                    if (already_in_page_music + 4 <= data.Music.Music.length){
                        start_music = already_in_page_music
                        end_music = already_in_page_music + 4
                    }else{
                        end_music = data.Music.Music.length
                    }
                    for (var i = start_music; i< end_music; i++){
                        var t = document.createElement('a');
                        t.classList.add('card_music','d-flex', 'justify-content-start');
                        t.appendChild(document.createElement('img'));
                        t.querySelector('img').classList.add('card_music_image','d-flex', 'mr-3', 'col-lg-3');
                        t.querySelector('img').id = 'card_image_music' + i;
                        t.appendChild(document.createElement('div'));
                        t.querySelector('div').classList.add('media-body');
                        t.querySelector('div').appendChild(document.createElement('h5'));
                        t.querySelector('div').querySelector('h5').textContent = data.Music.Music[i].Song;
                        t.querySelector('div').querySelector('h5').id = "song";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[0].textContent = data.Music.Music[i].Author;
                        t.querySelector('div').querySelectorAll('p')[0].id = "author";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[1].textContent = data.Music.Music[i].Salon;
                        t.querySelector('div').querySelectorAll('p')[1].id = "salon";
                        document.getElementById('card_inner_music').appendChild(t);
                        document.getElementById("card_image_music"+i).src="/static/images/divide.jpg";
                        let s = data.Music.Music[i].Song
                        let a = data.Music.Music[i].Author
                        let sl = data.Music.Music[i].Salon
                        t.onclick = function(s,a,sl) {return function() { openMusicCard(s,a,sl); }}(s,a,sl);
                    }
                    already_in_page_music = end_music
                    start_music = already_in_page_music
                    if (data.Music.Music.length > already_in_page_music){
                        btn = document.createElement('div');
                        btn.classList.add('d-flex', 'justify-content-center');
                        btn.appendChild(document.createElement('button'));
                        btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                        btn.querySelector('button').type = "submit";
                        btn.querySelector('button').textContent = "See more";
                        btn.querySelector('button').id = "see_more_btn_music"
                        btn.onclick = function() {return function() { openRemainingMusicCard(); }}();
                        document.getElementById('music_collections_inner').appendChild(btn)
                    }
                }
            }
            if (data.Collection != null){
                if (data.Collection.Collection.length > 0){
                    if (start_collection == 0){
                        document.getElementById('empty_collections_message').remove();
                        document.getElementById('buy_collections_empty').remove();
                    }
                    else if(start_collection > 0){
                        document.getElementById('see_more_btn_collection').remove()
                    }
                    
                    if (already_in_page_collection + 4 <= data.Collection.Collection.length){
                        start_collection = already_in_page_collection
                        end_collection = already_in_page_collection + 4
                    }else{
                        end_collection = data.Collection.Collection.length
                    }
                    for (var i = start_collection; i< end_collection; i++){
                        var t = document.createElement('a');
                        t.classList.add('card_music','d-flex', 'justify-content-start');
                        t.appendChild(document.createElement('img'));
                        t.querySelector('img').classList.add('card_music_image','d-flex', 'mr-3', 'col-lg-3');
                        t.querySelector('img').id = 'card_image_' + i;
                        t.appendChild(document.createElement('div'));
                        t.querySelector('div').classList.add('media-body');
                        t.querySelector('div').appendChild(document.createElement('h5'));
                        t.querySelector('div').querySelector('h5').textContent = data.Collection.Collection[i].Song;
                        t.querySelector('div').querySelector('h5').id = "title";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[0].textContent = data.Collection.Collection[i].Author;
                        t.querySelector('div').querySelectorAll('p')[0].id = "value";
                        t.querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('div').querySelectorAll('p')[1].textContent = data.Collection.Collection[i].Salon;
                        t.querySelector('div').querySelectorAll('p')[1].id = "salon";
                        document.getElementById('card_inner_collections').appendChild(t);
                        document.getElementById("card_image_"+i).src="/static/images/divide.jpg";
                        let s = data.Collection.Collection[i].Song
                        let a = data.Collection.Collection[i].Author
                        let sl = data.Collection.Collection[i].Salon
                        t.onclick = function(s,a,sl) {return function() { openCollectionCard(s,a,sl); }}(s,a,sl);
                    }
                    already_in_page_collection = end_collection
                    start_collection = already_in_page_collection
                    if (data.Collection.Collection.length > already_in_page_collection){
                        var btn = document.createElement('div');
                        btn.classList.add('d-flex', 'justify-content-center');
                        btn.appendChild(document.createElement('button'));
                        btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                        btn.querySelector('button').type = "submit";
                        btn.querySelector('button').textContent = "See more";
                        btn.querySelector('button').id = "see_more_btn_collection"
                        btn.onclick = function() {return function() { openRemainingCollectionCard(); }}();
                        document.getElementById('collections_inner').appendChild(btn)
                    }
                }
            }
        })
    });
}
async function openRemainingCollectionCard(){
    getData('http://localhost:8080/collectInformation')
    .then((data) => {
        data.json().then(data =>{
            if (data.Collection.Collection.length > 0){
                if (start_collection == 0){
                    document.getElementById('empty_collections_message').remove();
                    document.getElementById('buy_collections_empty').remove();
                }
                else if(start_collection > 0){
                    document.getElementById('see_more_btn_collection').remove()
                }
                
                if (already_in_page_collection + 4 <= data.Collection.Collection.length){
                    start_collection = already_in_page_collection
                    end_collection = already_in_page_collection + 4
                }else{
                    end_collection = data.Collection.Collection.length
                }
                for (var i = start_collection; i< end_collection; i++){
                    var t = document.createElement('a');
                    t.classList.add('card_music','d-flex', 'justify-content-start');
                    t.appendChild(document.createElement('img'));
                    t.querySelector('img').classList.add('card_music_image','d-flex', 'mr-3', 'col-lg-3');
                    t.querySelector('img').id = 'card_image_' + i;
                    t.appendChild(document.createElement('div'));
                    t.querySelector('div').classList.add('media-body');
                    t.querySelector('div').appendChild(document.createElement('h5'));
                    t.querySelector('div').querySelector('h5').textContent = data.Collection.Collection[i].Song;
                    t.querySelector('div').querySelector('h5').id = "title";
                    t.querySelector('div').appendChild(document.createElement('p'));
                    t.querySelector('div').querySelectorAll('p')[0].textContent = data.Collection.Collection[i].Author;
                    t.querySelector('div').querySelectorAll('p')[0].id = "value";
                    t.querySelector('div').appendChild(document.createElement('p'));
                    t.querySelector('div').querySelectorAll('p')[1].textContent = data.Collection.Collection[i].Salon;
                    t.querySelector('div').querySelectorAll('p')[1].id = "salon";
                    document.getElementById('card_inner_collections').appendChild(t);
                    document.getElementById("card_image_"+i).src="/static/images/divide.jpg";
                    let s = data.Collection.Collection[i].Song
                    let a = data.Collection.Collection[i].Author
                    let sl = data.Collection.Collection[i].Salon
                    t.onclick = function(s,a,sl) {return function() { openCollectionCard(s,a,sl); }}(s,a,sl);
                }
                already_in_page_collection = end_collection
                start_collection = already_in_page_collection
                if (data.Collection.Collection.length > already_in_page_collection){
                    var btn = document.createElement('div');
                    btn.classList.add('d-flex', 'justify-content-center');
                    btn.appendChild(document.createElement('button'));
                    btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                    btn.querySelector('button').type = "submit";
                    btn.querySelector('button').textContent = "See more";
                    btn.querySelector('button').id = "see_more_btn_collection"
                    btn.onclick = function() {return function() { openRemainingCollectionCard(); }}();
                    document.getElementById('collections_inner').appendChild(btn)
                }
            }
        })
    })
}

async function openRemainingMusicCard(){
    getData('http://localhost:8080/collectInformation')
    .then((data) => {
        data.json().then(data =>{
            if (data.Music.Music.length > 0){
                if (start_music == 0){
                    document.getElementById('empty_tracks_message').remove();
                    document.getElementById('buy_tracks_empty').remove();
                    document.getElementById('see_more_btn_music').remove()
                }
                else if(start_music > 0){
                    document.getElementById('see_more_btn_music').remove()
                }
                
                if (data.Music.Music.length > 4){
                    start_music = already_in_page_music
                    end_music = already_in_page_music + 4
                }
                for (var i = start_music; i< end_music; i++){
                    var t = document.createElement('a');
                    t.classList.add('card_music','d-flex', 'justify-content-start');
                    t.appendChild(document.createElement('img'));
                    t.querySelector('img').classList.add('card_music_image','d-flex', 'mr-3', 'col-lg-3');
                    t.querySelector('img').id = 'card_image_music' + i;
                    t.appendChild(document.createElement('div'));
                    t.querySelector('div').classList.add('media-body');
                    t.querySelector('div').appendChild(document.createElement('h5'));
                    t.querySelector('div').querySelector('h5').textContent = data.Music.Music[i].Song;
                    t.querySelector('div').querySelector('h5').id = "song";
                    t.querySelector('div').appendChild(document.createElement('p'));
                    t.querySelector('div').querySelectorAll('p')[0].textContent = data.Music.Music[i].Author;
                    t.querySelector('div').querySelectorAll('p')[0].id = "author";
                    t.querySelector('div').appendChild(document.createElement('p'));
                    t.querySelector('div').querySelectorAll('p')[1].textContent = data.Music.Music[i].Salon;
                    t.querySelector('div').querySelectorAll('p')[1].id = "salon";
                    document.getElementById('card_inner_music').appendChild(t);
                    document.getElementById("card_image_music"+i).src="/static/images/divide.jpg";
                    let s = data.Music.Music[i].Song
                    let a = data.Music.Music[i].Author
                    let sl = data.Music.Music[i].Salon
                    t.onclick = function(s,a,sl) {return function() { openMusicCard(s,a,sl); }}(s,a,sl);
                }
                already_in_page_music = end_music
                start_music = already_in_page_music
                if (data.Music.Music.length > already_in_page_music){
                    btn = document.createElement('div');
                    btn.classList.add('d-flex', 'justify-content-center');
                    btn.appendChild(document.createElement('button'));
                    btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                    btn.querySelector('button').type = "submit";
                    btn.querySelector('button').textContent = "See more";
                    btn.querySelector('button').id = "see_more_btn_music"
                    btn.onclick = function() {return function() { openRemainingMusicCard(); }}();
                    document.getElementById('music_collections_inner').appendChild(btn)
                }
            }
        })
    })
}

var start = 0
async function generateBucketFromUserPage(){
    getData('http://localhost:8080/collectInformation')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (data.Bucket.Bucket.length > 0 && data.Bucket.Bucket.length > start){
                if (start == 0)
                {
                    document.getElementById('bucket_empty').remove();
                    document.getElementById('buy_btn_bucket').remove();
                }else if(data.Bucket.Bucket.length > start){
                    document.getElementById('buy_btn_bucket').remove();
                }
                for (var i = start; i< data.Bucket.Bucket.length; i++){
                    var t = document.createElement('div');
                    t.classList.add('bucket_inner_elem', 'd-flex', 'justify-content-between', 'align-items-center');
                    t.appendChild(document.createElement('a'));
                    t.querySelector('a').classList.add('bucket_elem','d-flex','flex-row');
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').querySelectorAll('p')[0].textContent = data.Bucket.Bucket[i].Song;
                    t.querySelector('a').querySelectorAll('p')[0].id = "song";
                    t.querySelector('a').querySelectorAll('p')[1].textContent = data.Bucket.Bucket[i].Author;
                    t.querySelector('a').querySelectorAll('p')[1].id = "author";
                    t.querySelector('a').querySelectorAll('p')[2].textContent = data.Bucket.Bucket[i].Salon;
                    t.querySelector('a').querySelectorAll('p')[2].id = "salon";
                    t.querySelector('a').querySelectorAll('p')[3].textContent = data.Bucket.Bucket[i].Price;
                    t.querySelector('a').querySelectorAll('p')[3].id = "price";
                    t.querySelector('a').querySelectorAll('p')[4].textContent = data.Bucket.Bucket[i].Count;
                    t.querySelector('a').querySelectorAll('p')[4].id = "count";
                    t.appendChild(document.createElement('p'));
                    t.appendChild(document.createElement('a'));
                    t.querySelectorAll('a')[1].appendChild(document.createElement('img'));
                    t.querySelectorAll('a')[1].querySelector('img').classList.add('bucket_close');
                    t.querySelectorAll('a')[1].querySelector('img').id = 'bucket_elem_close'
                    t.querySelectorAll('a')[1].querySelector('img').src = '/static/images/close.png';
                    document.getElementById('bucket_intro').appendChild(t);
                    let s = data.Bucket.Bucket[i].Song
                    let a = data.Bucket.Bucket[i].Author
                    let sl = data.Bucket.Bucket[i].Salon
                    t.onclick = function(s,a,sl) {return function() { openCollectionCard(s,a,sl); }}(s,a,sl);
                }
                start = data.Bucket.Bucket.length
                var btn = document.createElement('div');
                btn.classList.add('d-flex', 'justify-content-center');
                btn.appendChild(document.createElement('button'));
                btn.querySelector('button').classList.add('buy_btn_bucket', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                btn.querySelector('button').type = "submit";
                btn.querySelector('button').textContent = "Buy";
                btn.querySelector('button').onclick = function() {return function() { BuyAllBucket(); }}();
                document.getElementById('bucket_intro').appendChild(btn)
            }
        })
    })
}

var start_recom = 0
var already_in_page_recom = 0
var end_recom = 0

async function generateRecomendation(){
    getData('http://localhost:8080/generateRecomendation')
    .then((data) => {
        data.json().then(data =>{
            if (data.Collection!= null){
                if (data.Collection.length > 0){
                    if (start_recom == 0){
                        document.getElementById('cant_recomend').remove()
                    }
                    
                    if (already_in_page_recom + 2 <= data.Collection.length){
                        start_recom = already_in_page_recom
                        end_recom = already_in_page_recom + 2
                    }else{
                        end_recom = data.Collection.length
                    }
                    for (var i = start_recom; i< end_recom; i++){
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
                        t.querySelector('a').querySelector('div').querySelector('h5').textContent = data.Collection[i].Title
                        t.querySelector('a').querySelector('div').appendChild(document.createElement('p'));
                        t.querySelector('a').querySelector('div').querySelector('p').textContent = data.Collection[i].Owner
                        for (var j = 0; j< data.Collection[i].Collection.length; j++){
                            var main_div = t.appendChild(document.createElement('div'))
                            main_div.classList.add('extra_track', 'd-flex', 'align-items-center')
                            main_div.appendChild(document.createElement('div'))
                            main_div.querySelector('div').classList.add('d-flex', 'flex-row', 'justify-content-between')
                            main_div.querySelector('div').style.width = "100%"
                            main_div.querySelector('div').style.margin = " 0 10px 0 0"
                            main_div.querySelector('div').appendChild(document.createElement('p'))
                            main_div.querySelector('div').appendChild(document.createElement('p'))
                            main_div.querySelector('div').querySelectorAll('p')[0].textContent = data.Collection[i].Collection[j].Title
                            main_div.querySelector('div').querySelectorAll('p')[1].textContent = data.Collection[i].Collection[j].Artist
                            main_div.appendChild(document.createElement('button'))
                            main_div.querySelector('button').classList.add('w-50%', 'btn', 'btn-lg', 'btn-primary', 'buy_in_collection_card')
                            main_div.querySelector('button').type="submit"
                            main_div.querySelector('button').textContent="Buy"
                            let s = data.Collection[i].Collection[j].Title
                            let a = data.Collection[i].Collection[j].Artist
                            let sl = data.Collection[i].Owner
                            main_div.querySelector('button').onclick = function(s,a,sl) {return function() { addMusicCardToBucketFromCollection(s,a,sl); }}(s,a,sl);
                        }

                        var last = t.appendChild(document.createElement('div'))
                        last.classList.add('d-flex', 'justify-content-end', 'align-items-center')
                        last.appendChild(document.createElement('p'))
                        last.querySelector('p').style.margin = "0 10px"
                        last.querySelector('p').textContent = data.Collection[i].Price + "x" + data.Collection[i].Count
                        last.appendChild(document.createElement('button'))
                        last.querySelector('button').classList.add('w-50%','btn','btn-lg', 'btn-primary', 'buy_all_in_collection_card')
                        last.querySelector('button').type = "submit"
                        last.querySelector('button').textContent = "Buy all"
                        let title = data.Collection[i].Title
                        let s = data.Collection[i].Owner
                        last.querySelector('button').onclick = function(title,s) {return function() { addAllCollectionToBucket(title,s); }}(title,s);
                        document.getElementById('recom_collections_inner').appendChild(t)
                    }
                    already_in_page_recom = end_recom
                    start_recom = already_in_page_recom
                    if (data.Collection.length <= already_in_page_recom){
                        document.getElementById('see_more_recomendation').remove()
                    }
                }
            }
        })
    })
}