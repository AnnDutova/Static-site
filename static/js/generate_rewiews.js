async function generateRewiews(){
    getData('http://localhost:8080/getAllRewiews')
    .then((data) => {
        data.json().then(data =>{
            console.log(data);
            if (data[0].username != null){
                document.getElementById('buy_tracks_empty').remove();
                document.getElementById('empty_tracks_message').remove();
                for (var i = 0; i< data.length; i++){
                    var t = document.createElement('div')
                    t.classList.add('rewiew_old')
                    t.appendChild(document.createElement('header'))
                    t.querySelector('header').classList.add('d-flex', 'justify-content-between', 'flex-row')
                    t.querySelector('header').appendChild(document.createElement('div'))
                    t.querySelector('header').appendChild(document.createElement('div'))
                    t.querySelector('header').querySelectorAll('div')[0].classList.add('author', 'd-flex','align-items-center')
                    t.querySelector('header').querySelectorAll('div')[0].appendChild(document.createElement('img'))
                    t.querySelector('header').querySelectorAll('div')[0].querySelector('img').classList.add('image_in_rewiew')
                    t.querySelector('header').querySelectorAll('div')[0].querySelector('img').src = "/static/images/user_white.svg"
                    t.querySelector('header').querySelectorAll('div')[0].appendChild(document.createElement('p'))
                    t.querySelector('header').querySelectorAll('div')[0].querySelector('p').textContent = data[i].username;
                    t.querySelector('header').querySelectorAll('div')[1].classList.add('grade', 'd-flex', 'align-items-center')
                    t.querySelector('header').querySelectorAll('div')[1].appendChild(document.createElement('p'))
                    t.querySelector('header').querySelectorAll('div')[1].querySelector('p').textContent = " " + data[i].grade + "/5"
                    t.querySelector('header').querySelectorAll('div')[1].appendChild(document.createElement('div'))
                    t.querySelector('header').querySelectorAll('div')[1].querySelector('div').classList.add('rewiew_star')
                    t.querySelector('header').querySelectorAll('div')[1].querySelector('div').textContent = '★'
                    t.appendChild(document.createElement('textarea'))
                    t.querySelector('textarea').classList.add('old_rewiew_textblok')
                    t.querySelector('textarea').id = "userRewiew"
                    t.querySelector('textarea').type = "text" 
                    t.querySelector('textarea').name = "rewiew"
                    t.querySelector('textarea').readOnly
                    t.querySelector('textarea').textContent = data[i].text
                    document.getElementById('rewiew_inner').appendChild(t);
                }

            }

        })
    })
}