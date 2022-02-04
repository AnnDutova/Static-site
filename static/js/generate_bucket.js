async function generateBucket(){
    getData('http://localhost:8080/collectBucket')
    .then((data) => {
        data.json().then(data =>{
            console.log(data);
            if (data.Bucket.length > 0){
                if (document.getElementById('bucket_empty') != null)
                {
                    document.getElementById('bucket_empty').remove();
                    document.getElementById('buy_btn_bucket').remove();
                }
                for (var i = 0; i< data.Bucket.length; i++){
                    var t = document.createElement('div');
                    t.id = "bucket_inner_elem_id"
                    t.classList.add('bucket_inner_elem', 'd-flex', 'justify-content-between', 'align-items-center');
                    t.appendChild(document.createElement('a'));
                    t.querySelector('a').classList.add('bucket_elem', 'd-flex');
                    t.querySelector('a').classList.add('flex-row');
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').appendChild(document.createElement('p'));
                    t.querySelector('a').querySelectorAll('p')[0].textContent = data.Bucket[i].Song;
                    t.querySelector('a').querySelectorAll('p')[0].id = "song";
                    t.querySelector('a').querySelectorAll('p')[1].textContent = data.Bucket[i].Author;
                    t.querySelector('a').querySelectorAll('p')[1].id = "author";
                    t.querySelector('a').querySelectorAll('p')[2].textContent = data.Bucket[i].Salon;
                    t.querySelector('a').querySelectorAll('p')[2].id = "salon";
                    t.querySelector('a').querySelectorAll('p')[3].textContent = data.Bucket[i].Price;
                    t.querySelector('a').querySelectorAll('p')[3].id = "price";
                    t.querySelector('a').querySelectorAll('p')[4].textContent = data.Bucket[i].Count;
                    t.querySelector('a').querySelectorAll('p')[4].id = "count";
                    t.appendChild(document.createElement('p'));
                    t.appendChild(document.createElement('a'));
                    t.querySelectorAll('a')[1].appendChild(document.createElement('img'));
                    t.querySelectorAll('a')[1].querySelector('img').classList.add('bucket_close');
                    t.querySelectorAll('a')[1].querySelector('img').id = 'bucket_elem_close'
                    t.querySelectorAll('a')[1].querySelector('img').src = '/static/images/close.png';
                    document.getElementById('bucket_intro').appendChild(t);
                    let s = data.Bucket[i].Song
                    let a = data.Bucket[i].Author
                    let sl = data.Bucket[i].Salon
                    t.querySelectorAll('a')[1].onclick = function(s,a,sl) {return function() { DeliteElementFromBucket(s,a,sl); }}(s,a,sl);
                }
                var btn = document.createElement('div');
                btn.classList.add('d-flex', 'justify-content-center');
                btn.appendChild(document.createElement('button'));
                btn.querySelector('button').classList.add('buy_btn_bucket', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                btn.querySelector('button').type = "submit";
                btn.querySelector('button').textContent = "Buy";
                btn.querySelector('button').id = "buy_btn_bucket"
                var bucket_mas = data.Bucket
                btn.querySelector('button').onclick = function(bucket_mas) {return function() { BuyAllBucket(bucket_mas); }}(bucket_mas);
                document.getElementById('bucket_intro').appendChild(btn)
            }else if (data.Bucket.length == 0 && document.getElementById('bucket_empty') == null){
                while (document.getElementById('bucket_inner_elem_id') != null){
                    document.getElementById('bucket_inner_elem_id').remove();
                }
                document.getElementById('buy_btn_bucket').remove();

                var empty_img = document.createElement('div');
                empty_img.classList.add('d-flex', 'justify-content-around', 'align-items-center')
                empty_img.id = "bucket_empty"
                empty_img.appendChild(document.createElement("img"))
                empty_img.querySelector('img').src = "/static/images/empty_box.png"
                empty_img.querySelector('img').style.height = "50px"
                empty_img.querySelector('img').style.width = "50px"
                empty_img.querySelector('img').style.margin = "10px 30px"
                empty_img.appendChild(document.createElement("p"))
                empty_img.querySelector('p').style.width = "160px"
                empty_img.querySelector('p').style.margin = "0 15px 0 0"
                empty_img.querySelector('p').textContent = "Your bucket is empty"

                var btn = document.createElement('div');
                btn.classList.add('d-flex', 'justify-content-center');
                btn.appendChild(document.createElement('a'));
                btn.querySelector('a').classList.add('buy_btn_bucket', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                btn.querySelector('a').id = "buy_btn_bucket"
                btn.querySelector('a').textContent = "Buy";
                btn.querySelector('a').href = "/market"

                document.getElementById('bucket_intro').appendChild(empty_img)
                document.getElementById('bucket_intro').appendChild(btn)
            }
        })
    })
}