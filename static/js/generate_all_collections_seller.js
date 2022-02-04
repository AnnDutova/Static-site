var openAllCollections = true
function generateAllCollectionsSeller(){
    getData('http://localhost:8080/collectCollectionsSeller')
    .then((data) => {
        data.json().then(data =>{
            console.log(data);
            if (data.Collection != null){
                console.log(open)
                if (openAllCollections){
                    for (var i = 0; i< data.Collection.length; i++){
                        var t = document.createElement('div');
                        t.classList.add('collection', 'collection_main_card')
                        t.appendChild(document.createElement('a'));
                        t.querySelector('a').classList.add('d-flex', 'justify-content-start', 'collection_header');
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
                            t.appendChild(document.createElement('div'))
                            t.querySelectorAll('div')[j+1].classList.add('extra_track', 'd-flex', 'justify-content-between', 'align-items-center')
                            t.querySelectorAll('div')[j+1].appendChild(document.createElement('p'))
                            t.querySelectorAll('div')[j+1].appendChild(document.createElement('p'))
                            t.querySelectorAll('div')[j+1].querySelectorAll('p')[0].textContent = data.Collection[i].Collection[j].Title
                            t.querySelectorAll('div')[j+1].querySelectorAll('p')[1].textContent = data.Collection[i].Collection[j].Artist
                            t.querySelectorAll('div')[j+1].appendChild(document.createElement('p'))
                            t.querySelectorAll('div')[j+1].querySelectorAll('p')[2].style.margin = "0 10px"
                            t.querySelectorAll('div')[j+1].querySelectorAll('p')[2].textContent = data.Collection[i].Collection[j].Price + "x" + data.Collection[i].Collection[j].Count +"x"+ data.Collection[i].Collection[j].Sale
                        }

                        var last = t.appendChild(document.createElement('div'))
                        last.classList.add('d-flex', 'justify-content-end', 'align-items-center')
                        last.appendChild(document.createElement('p'))
                        last.querySelector('p').style.margin = "0 10px"
                        last.querySelector('p').textContent = data.Collection[i].Price + "x" + data.Collection[i].Count
                        document.getElementById('collection_inner_main').appendChild(t);
                    }
                    openAllCollections = false
                }else{
                    document.getElementById('collection_inner_main').remove()
                    document.getElementById('seller_collectionn').appendChild(document.createElement('div'))
                    var new_inner = document.getElementById('seller_collectionn').querySelector('div')
                    new_inner.classList.add('d-flex', 'flex-wrap')
                    new_inner.id = "collection_inner_main"
                    openAllCollections = true 
                }
            }
        })
    });
}