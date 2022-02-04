var openAllMusic = true
async function generateAllMusicForCollection(){
    getData('http://localhost:8080/collectMusicinSeller')
    .then((data) => {
        data.json().then(data =>{
            console.log(data);
            if (data.Music != null){
                if (openAllMusic){
                    var title = document.createElement('div')
                    title.classList.add('form-floating', 'm-2')
                    title.appendChild(document.createElement('input'))
                    title.appendChild(document.createElement('label'))
                    title.querySelector('input').type = "text"
                    title.querySelector('input').classList.add('form-control', 'form_style')
                    title.querySelector('input').id = "collection_title"
                    title.querySelector('input').placeholder =  "title"
                    title.querySelector('label').for = "floatingInput"
                    title.querySelector('label').style.opacity = "0.4"
                    title.querySelector('label').textContent = 'Composition title'
                    
                    var price = document.createElement('div')
                    price.classList.add('form-floating', 'm-2')
                    price.appendChild(document.createElement('input'))
                    price.appendChild(document.createElement('label'))
                    price.querySelector('input').type = "number"
                    price.querySelector('input').classList.add('form-control', 'form_style')
                    price.querySelector('input').id = "collection_price"
                    price.querySelector('input').placeholder =  "price"
                    price.querySelector('label').for = "floatingInput"
                    price.querySelector('label').style.opacity = "0.4"
                    price.querySelector('label').textContent = 'Price'

                    var sale = document.createElement('div')
                    sale.classList.add('form-floating', 'm-2')
                    sale.appendChild(document.createElement('input'))
                    sale.appendChild(document.createElement('label'))
                    sale.querySelector('input').type = "number"
                    sale.querySelector('input').classList.add('form-control', 'form_style')
                    sale.querySelector('input').id = "collection_sale"
                    sale.querySelector('input').placeholder =  "sale"
                    sale.querySelector('label').for = "floatingInput"
                    sale.querySelector('label').style.opacity = "0.4"
                    sale.querySelector('label').textContent = 'Sale'

                    document.getElementById('form').appendChild(title)
                    document.getElementById('form').appendChild(price)
                    document.getElementById('form').appendChild(sale)

                    for (var i = 0; i< data.Music.length; i++){
                        var t = document.createElement('div');
                        t.classList.add('d-flex', 'flex-row', 'align-items-center', 'custom-control', 'custom-checkbox');
                        t.appendChild(document.createElement('input'))
                        t.querySelector('input').type="checkbox"
                        t.querySelector('input').classList.add('custom-control-input')
                        t.querySelector('input').id = i
                        t.appendChild(document.createElement('a'))
                        t.querySelector('a').classList.add('card_music','d-flex', 'justify-content-start')
                        var inpput_a =  t.querySelector('a')
                        inpput_a.appendChild(document.createElement('img'));
                        inpput_a.querySelector('img').classList.add('card_music_image','d-flex', 'mr-3', 'col-lg-3');
                        inpput_a.querySelector('img').id = 'card_image_col_' + i;
                        inpput_a.appendChild(document.createElement('div'));
                        inpput_a.querySelector('div').classList.add('media-body');
                        inpput_a.querySelector('div').id = 'card_'+i
                        inpput_a.querySelector('div').appendChild(document.createElement('h5'));
                        inpput_a.querySelector('div').querySelector('h5').textContent = data.Music[i].Song;
                        inpput_a.querySelector('div').querySelector('h5').id = "song";
                        inpput_a.querySelector('div').appendChild(document.createElement('p'));
                        inpput_a.querySelector('div').querySelectorAll('p')[0].textContent = data.Music[i].Author;
                        inpput_a.querySelector('div').querySelectorAll('p')[0].id = "author";
                        inpput_a.querySelector('div').appendChild(document.createElement('p'));
                        inpput_a.querySelector('div').querySelectorAll('p')[1].textContent = data.Music[i].Salon;
                        inpput_a.querySelector('div').querySelectorAll('p')[1].id = "salon";
                        document.getElementById('collection_inner').appendChild(t);
                        document.getElementById("card_image_col_"+i).src="/static/images/divide.jpg";    
                    }
                    btn = document.createElement('div');
                    btn.classList.add('d-flex', 'justify-content-center');
                    btn.appendChild(document.createElement('button'));
                    btn.querySelector('button').classList.add('track_btn', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                    btn.querySelector('button').type = "submit";
                    btn.querySelector('button').textContent = "Create collection";
                    document.getElementById('collection_btn').appendChild(btn)
                    btn.onclick = function(number) {return function() { createCollectionSeller(number); }}(data.Music.length); 
                    document.getElementById('container_form').style.backgroundColor = "white"
                    openAllMusic = false
                }else {
                    document.getElementById('form').remove()
                    document.getElementById('collection_inner').remove()
                    document.getElementById('collection_btn').remove()
                    var form =  document.createElement('div');
                    form.id = "form"
                    var col_in = document.createElement('div');
                    col_in.classList.add('d-flex', 'flex-wrap')
                    col_in.id = "collection_inner"
                    var btn =  document.createElement('div');
                    btn.id = "collection_btn"
                    document.getElementById('container_form').appendChild(form)
                    document.getElementById('container_form').appendChild(col_in)
                    document.getElementById('container_form').appendChild(btn)
                    document.getElementById('container_form').style = null
                    openAllMusic = true
                }       
            }
        })
    });
}