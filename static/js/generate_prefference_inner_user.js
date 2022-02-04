var openPref = true
async function generatePreferenceInner(){
    getData('http://localhost:8080/getPreferences')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (data.length>0){
                if(openPref == true)
                {
                    for (var i = 0; i< data.length; i++){
                        var t = document.createElement('p')
                        t.classList.add('preference_type')
                        t.id = "preference_type_el"
                        t.textContent = data[i]
                        document.getElementById('preference_inner').appendChild(t)
                    }

                    var btn = document.createElement('div');
                    btn.classList.add('d-flex', 'justify-content-center');
                    btn.id = "new_preference"
                    btn.appendChild(document.createElement('button'));
                    btn.querySelector('button').classList.add('buy_btn_bucket', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                    btn.querySelector('button').type = "submit";
                    btn.querySelector('button').textContent = "Add new";
                    btn.querySelector('button').id = "add_pref_form"
                    btn.querySelector('button').onclick = function() {return function() { generateNewPreferenceForm(); }}();
                    document.getElementById('preference_inner').appendChild(btn)

                    openPref= false
                }else{
                    while (document.getElementById('preference_type_el') != null){
                        document.getElementById('preference_type_el').remove()
                    }
                    while (document.getElementById('new_preference') != null){
                        document.getElementById('new_preference').remove()
                    }
                    openPref = true
                }
            }else if(data.length==0 || data.error == 'sql: no rows in result set'){

                console.log("Need to add preferences")
                var t = document.createElement('p')
                t.classList.add('preference_type')
                t.id = "preference_type_el"
                t.textContent = "You dont have preferences :("
                document.getElementById('preference_inner').appendChild(t)

                var btn = document.createElement('div');
                btn.classList.add('d-flex', 'justify-content-center');
                btn.id = "new_preference"
                btn.appendChild(document.createElement('button'));
                btn.querySelector('button').classList.add('buy_btn_bucket', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                btn.querySelector('button').type = "submit";
                btn.querySelector('button').textContent = "Add new";
                btn.querySelector('button').id = "add_pref_form"
                btn.querySelector('button').onclick = function() {return function() { generateNewPreferenceForm(); }}();
                document.getElementById('preference_inner').appendChild(btn)

            }
        })
    })
}
async function generateNewPreferenceForm(){
    getData('http://localhost:8080/getArtistGenres')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (data != null){
                for (var i = 0; i< data.Genre.length; i++){
                    var t = document.createElement('div')
                    t.id = "preference_type_el"
                    t.classList.add('d-flex', 'align-items-center', 'justify-content-start')
                    t.appendChild(document.createElement('input'))
                    t.querySelector('input').type = "checkbox"
                    t.querySelector('input').value =  data.Genre[i]
                    t.querySelector('input').id = "checkbox_value_"+i
                    t.appendChild(document.createElement('p'))
                    t.querySelector('p').style.margin = "0 5px"
                    t.querySelector('p').textContent = data.Genre[i]
                    document.getElementById('preference_inner').appendChild(t);
                }
                var btn = document.createElement('div');
                    btn.classList.add('d-flex', 'justify-content-center');
                    btn.id = "new_preference"
                    btn.appendChild(document.createElement('button'));
                    btn.querySelector('button').classList.add('buy_btn_bucket', 'w-50%', 'btn', 'btn-lg', 'btn-primary');
                    btn.querySelector('button').type = "submit";
                    btn.querySelector('button').textContent = "Add new";
                    btn.querySelector('button').id = "add_new_pref"
                    var number = data.Genre.length
                    btn.querySelector('button').onclick = function(number) {return function() { AddNewPreference(number); }}(number);
                    document.getElementById('preference_inner').appendChild(btn)

            }
        })
    })
}